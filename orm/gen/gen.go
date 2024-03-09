package gen

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"sync"

	"github.com/fzf-labs/fdatabase/orm/gen/proto"
	"github.com/fzf-labs/fdatabase/orm/gen/repo"
	"github.com/fzf-labs/fdatabase/orm/utils"
	"github.com/fzf-labs/fdatabase/orm/utils/dbfunc"
	"github.com/fzf-labs/fdatabase/orm/utils/file"
	"github.com/iancoleman/strcase"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// //////////////////////////////////////
// NewGenerationDB SQL 生成 dao,model,repo
// //////////////////////////////////////
const (
	SQLNullTime = "sql.NullTime"
	TimeTime    = "time.Time"
)

type GenerationDB struct {
	db               *gorm.DB                                                      // 数据库
	outPutPath       string                                                        // 文件生成路径
	genRepo          bool                                                          // 是否生成repo文件
	dataMap          map[string]func(columnType gorm.ColumnType) (dataType string) // 自定义字段类型映射
	tables           []string                                                      // 指定表集合
	opts             []gen.ModelOpt                                                // 特殊处理逻辑函数
	dbNameOpt        func(*gorm.DB) string                                         // 指定数据库名
	generateModelOpt func(g *gen.Generator) map[string]any                         // 指定表对应的model
	fieldNullable    bool                                                          // 字段是否可空
}

func NewGenerationDB(db *gorm.DB, outPutPath string, opts ...OptionDB) *GenerationDB {
	g := &GenerationDB{
		db:         db,
		outPutPath: outPutPath,
		genRepo:    true,
		dataMap:    nil,
		tables:     nil,
		opts:       nil,
	}
	if len(opts) > 0 {
		for _, v := range opts {
			v(g)
		}
	}
	return g
}

type OptionDB func(gen *GenerationDB)

// WithOutRepo 选项函数-不生成repo
func WithOutRepo() OptionDB {
	return func(r *GenerationDB) {
		r.genRepo = false
	}
}

// WithTables 选项函数-自定义表
func WithTables(tables []string) OptionDB {
	return func(r *GenerationDB) {
		r.tables = tables
	}
}

// WithDataMap 选项函数-自定义关系映射
func WithDataMap(dataMap map[string]func(columnType gorm.ColumnType) (dataType string)) OptionDB {
	return func(r *GenerationDB) {
		r.dataMap = dataMap
	}
}

// WithDBOpts 选项函数-自定义特殊设置
func WithDBOpts(opts ...gen.ModelOpt) OptionDB {
	return func(r *GenerationDB) {
		r.opts = opts
	}
}

// WithDBNameOpts 选项函数-自定义数据库名
func WithDBNameOpts(fn func(*gorm.DB) string) OptionDB {
	return func(r *GenerationDB) {
		r.dbNameOpt = fn
	}
}

// WithGenerateModel 选项函数-自定义表的关联关系
func WithGenerateModel(fn func(g *gen.Generator) map[string]any) OptionDB {
	return func(r *GenerationDB) {
		r.generateModelOpt = fn
	}
}

// WithFieldNullable 选项函数-字段是否可空
func WithFieldNullable() OptionDB {
	return func(r *GenerationDB) {
		r.fieldNullable = true
	}
}

// Do 生成
func (g *GenerationDB) Do() {
	// 获取数据库名
	dbName := GetDBName(g.db, g.dbNameOpt)
	// 文件夹目录
	outPutPath := strings.Trim(g.outPutPath, "/")
	daoPath := fmt.Sprintf("%s/%s_dao", outPutPath, dbName)
	modelPath := fmt.Sprintf("%s/%s_model", outPutPath, dbName)
	repoPath := fmt.Sprintf("%s/%s_repo", outPutPath, dbName)
	// 初始化
	generator := gen.NewGenerator(gen.Config{
		OutPath:          daoPath,
		ModelPkgPath:     modelPath,
		FieldWithTypeTag: true, // gorm tag 中会增加type类型
		FieldNullable:    g.fieldNullable,
	})
	// 使用数据库
	generator.UseDB(g.db)
	// 指定数据库名
	if g.dbNameOpt != nil {
		generator.WithDbNameOpts(g.dbNameOpt)
	}
	// 自定义字段类型映射
	if g.dataMap != nil {
		generator.WithDataTypeMap(g.dataMap)
	}
	// json 小驼峰模型命名
	generator.WithJSONTagNameStrategy(JSONTagNameStrategy)
	// 特殊处理逻辑
	if len(g.opts) > 0 {
		generator.WithOpts(g.opts...)
	}
	// 获取所有表
	tables, err := g.db.Migrator().GetTables()
	if err != nil {
		return
	}
	// 指定表
	if len(g.tables) > 0 {
		tables = g.tables
	}
	// 查询分区表父级到子表的映射
	partitionTableToChildTables, err := dbfunc.GetPartitionTableToChildTables(g.db)
	if err != nil {
		return
	}
	partitionChildTables := make([]string, 0)
	for _, v := range partitionTableToChildTables {
		partitionChildTables = append(partitionChildTables, v...)
	}
	// 去掉tables中的partitionChildTables
	tables = utils.SliRemove(tables, partitionChildTables)
	models := make(map[string]any, len(tables))
	for _, tableName := range tables {
		generateModel := generator.GenerateModel(tableName)
		if _, ok := partitionTableToChildTables[tableName]; ok {
			generatePartitionChildModel := generator.GenerateModel(partitionTableToChildTables[tableName][0])
			generateModel.Fields = generatePartitionChildModel.Fields
			generatePartitionChildModel.Generated = false
		}
		models[tableName] = generateModel
	}
	if g.generateModelOpt != nil {
		customModels := g.generateModelOpt(generator)
		for k, v := range customModels {
			models[k] = v
		}
	}
	applyModels := make([]any, 0)
	for _, v := range models {
		applyModels = append(applyModels, v)
	}
	generator.ApplyBasic(applyModels...)
	// 生成model,dao
	generator.Execute()
	// 判断是否生成repo
	if !g.genRepo {
		return
	}
	// 生成repo的文件夹目录文件
	err = file.MkdirPath(repoPath)
	if err != nil {
		log.Println("repo MkdirPath err:", err)
		return
	}
	var wg sync.WaitGroup
	wg.Add(len(tables))
	for _, v := range tables {
		table := v
		// 表字段对应的类型
		columnNameToDataType := make(map[string]string)
		// 表字段对应的名称
		columnNameToName := make(map[string]string)
		// 表字段对应的dao字段类型
		columnNameToFieldType := make(map[string]string)
		queryStructMeta := generator.GenerateModel(table)
		for _, vv := range queryStructMeta.Fields {
			columnNameToDataType[vv.ColumnName] = strings.TrimLeft(vv.Type, "*")
			columnNameToName[vv.ColumnName] = vv.Name
			columnNameToFieldType[vv.ColumnName] = vv.GenType()
		}
		go func(db *gorm.DB, table string, columnNameToDataType, columnNameToName, columnNameToFieldType map[string]string) {
			defer wg.Done()
			// 数据表repo代码生成
			err2 := repo.GenerationTable(db, dbName, daoPath, modelPath, repoPath, table, columnNameToDataType, columnNameToName, columnNameToFieldType)
			if err2 != nil {
				log.Println("repo GenerationTable err:", err2)
				return
			}
		}(g.db, table, columnNameToDataType, columnNameToName, columnNameToFieldType)
	}
	wg.Wait()
}

// GetDBName 获取数据库名
func GetDBName(db *gorm.DB, fn func(*gorm.DB) string) string {
	tableName := db.Migrator().CurrentDatabase()
	if fn != nil {
		tableName = fn(db)
	}
	tablePrefix := ""
	if ns, ok := db.NamingStrategy.(schema.NamingStrategy); ok {
		tablePrefix = ns.TablePrefix
	}
	if !strings.HasPrefix(tableName, tablePrefix) {
		tableName = tablePrefix + tableName
	}
	return tableName
}

// JSONTagNameStrategy json tag 命名
func JSONTagNameStrategy(s string) string {
	// 下划线单词转为小写驼峰单词
	return strcase.ToLowerCamel(s)
}

// ModelOptionUnderline 前缀是下划线重命名
func ModelOptionUnderline(rename string) gen.ModelOpt {
	return gen.FieldModify(func(f gen.Field) gen.Field {
		if strings.HasPrefix(f.Name, "_") {
			f.Name = strings.Replace(f.Name, "_", rename, 1)
			f.Tag.Set(field.TagKeyJson, f.ColumnName)
		}
		return f
	})
}

// ModelOptionPgDefaultString Postgres默认字符串处理
func ModelOptionPgDefaultString() gen.ModelOpt {
	return gen.FieldGORMTagReg(".*?", func(tag field.GormTag) field.GormTag {
		regex := regexp.MustCompile(`default:'(.*?)'::character varying`)
		matches := regex.FindStringSubmatch(tag.Build())
		if len(matches) > 0 {
			tag.Set("default", matches[1])
		}
		return tag
	})
}

// ModelOptionRemoveGormTypeTag 移除gorm tag :type
func ModelOptionRemoveGormTypeTag() gen.ModelOpt {
	return gen.FieldGORMTagReg(".*?", func(tag field.GormTag) field.GormTag {
		tag.Remove("type")
		return tag
	})
}

// ModelOptionRemoveDefault 默认字符串移除(主键除外)
func ModelOptionRemoveDefault() gen.ModelOpt {
	return gen.FieldGORMTagReg(".*?", func(tag field.GormTag) field.GormTag {
		regex := regexp.MustCompile(`primaryKey`)
		matches := regex.FindStringSubmatch(tag.Build())
		if len(matches) == 0 {
			tag.Remove("default")
		}
		return tag
	})
}

// PGDataTypeMap 自定义字段类型映射
func PGDataTypeMap() map[string]func(columnType gorm.ColumnType) (dataType string) {
	return map[string]func(columnType gorm.ColumnType) (dataType string){
		"json":  func(_ gorm.ColumnType) string { return "datatypes.JSON" },
		"jsonb": func(_ gorm.ColumnType) string { return "datatypes.JSON" },
		"timestamptz": func(columnType gorm.ColumnType) string {
			if utils.StrSliFind([]string{"deleted_at", "deletedAt", "deleted_time", "deleted_time"}, columnType.Name()) {
				return "gorm.DeletedAt"
			}
			nullable, _ := columnType.Nullable()
			if nullable {
				return SQLNullTime
			}
			return TimeTime
		},
		"character varying[]": func(_ gorm.ColumnType) (dataType string) {
			return "pq.StringArray"
		},
		"smallint[]": func(_ gorm.ColumnType) (dataType string) {
			return "pq.Int32Array"
		},
		"integer[]": func(_ gorm.ColumnType) (dataType string) {
			return "pq.Int32Array"
		},
		"bigint[]": func(_ gorm.ColumnType) (dataType string) {
			return "pq.Int64Array"
		},
	}
}

// DBNameOpts 自定义数据库名函数
func DBNameOpts() func(*gorm.DB) string {
	return func(db *gorm.DB) string {
		tableName := db.Migrator().CurrentDatabase()
		tableName = strings.ReplaceAll(tableName, "-", "_")
		tableName = strings.ReplaceAll(tableName, " ", "")
		return tableName
	}
}

// ConnectDB 数据库连接
func ConnectDB(dbType, dsn string) *gorm.DB {
	var db *gorm.DB
	var err error
	switch dbType {
	case "mysql":
		db, err = gorm.Open(mysql.Open(dsn))
		if err != nil {
			panic(fmt.Errorf("connect mysql db fail: %s", err))
		}
	case "postgres":
		db, err = gorm.Open(postgres.Open(dsn))
		if err != nil {
			panic(fmt.Errorf("connect postgres db fail: %s", err))
		}
	default:
		panic(fmt.Errorf(" db type err"))
	}
	return db
}

////////////////////////////////////////
// NewGenerationPB SQL 生成 proto
////////////////////////////////////////

// NewGenerationPB SQL 生成 proto
func NewGenerationPB(db *gorm.DB, outPutPath, packageStr, goPackageStr string, opts ...OptionPB) *GenerationPb {
	g := &GenerationPb{
		db:           db,
		outPutPath:   outPutPath,
		packageStr:   packageStr,
		goPackageStr: goPackageStr,
	}
	if len(opts) > 0 {
		for _, v := range opts {
			v(g)
		}
	}
	return g
}

type GenerationPb struct {
	db           *gorm.DB       // 数据库
	outPutPath   string         // 文件生成地址
	opts         []gen.ModelOpt // 特殊处理逻辑函
	packageStr   string
	goPackageStr string
}

type OptionPB func(gen *GenerationPb)

// WithPBOpts 选项函数-自定义特殊设置
func WithPBOpts(opts ...gen.ModelOpt) OptionPB {
	return func(r *GenerationPb) {
		r.opts = opts
	}
}

func (g *GenerationPb) Do() {
	// 初始化
	generator := gen.NewGenerator(gen.Config{})
	// 使用数据库
	generator.UseDB(g.db)
	// json 小驼峰模型命名
	generator.WithJSONTagNameStrategy(JSONTagNameStrategy)
	// 特殊处理逻辑
	if len(g.opts) > 0 {
		generator.WithOpts(g.opts...)
	}
	// 获取所有表
	tables, err := g.db.Migrator().GetTables()
	if err != nil {
		return
	}
	// 查询分区所有子表
	partitionChildTables, err := dbfunc.GetPartitionChildTable(g.db)
	if err != nil {
		return
	}
	// 去掉tables中的partitionChildTables
	tables = utils.SliRemove(tables, partitionChildTables)
	var wg sync.WaitGroup
	wg.Add(len(tables))
	for _, v := range tables {
		table := v
		// 表字段对应的名称
		columnNameToName := make(map[string]string)
		queryStructMeta := generator.GenerateModel(table)
		for _, vv := range queryStructMeta.Fields {
			columnNameToName[vv.ColumnName] = vv.Name
		}
		go func(db *gorm.DB, outPutPath, packageStr, goPackageStr, table string, columnNameToName map[string]string) {
			defer wg.Done()
			// 数据表repo代码生成
			err2 := proto.GenerationPB(db, outPutPath, packageStr, goPackageStr, table, columnNameToName)
			if err2 != nil {
				log.Println("repo GenerationTable err:", err2)
				return
			}
		}(g.db, g.outPutPath, g.packageStr, g.goPackageStr, table, columnNameToName)
	}
	wg.Wait()
}

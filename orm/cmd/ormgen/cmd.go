package ormgen

import (
	"strings"

	"github.com/fzf-labs/fctl/utils"
	gormGen "gorm.io/gen"

	"github.com/spf13/cobra"
	"gitlab.yc345.tv/backend/utils/v2/orm/gen"
)

var CmdOrmGen = &cobra.Command{
	Use:   "ormgen",
	Short: "orm gen",
	Long:  "orm gen",
	Run:   Run,
}

var (
	dsn                     string // 数据库自定义连接
	targetTables            string // 数据库指定表
	outPutPath              string // 输出路径
	optionUnderline         string // 选项：下划线转驼峰 (默认'_'替换为'UL')
	optionPgDefaultString   bool   // 选项：移除gorm tag default:'(.*?)'::character varying  (默认是 true)
	optionRemoveDefault     bool   // 选项：移除默认值 （默认是 true）
	optionRemoveGormTypeTag bool   // 选项：移除gorm tag :type (默认是 false)
)

//nolint:gochecknoinits
func init() {
	CmdOrmGen.Flags().StringVarP(&dsn, "dsn", "d", "", "db dsn")
	CmdOrmGen.Flags().StringVarP(&targetTables, "tables", "t", "", "db tables")
	CmdOrmGen.Flags().StringVarP(&outPutPath, "outPutPath", "o", "./internal/data/gorm", "output path")
	CmdOrmGen.Flags().StringVarP(&optionUnderline, "optionUnderline", "", "UL", "option: underline '_' replace 'UL'")
	CmdOrmGen.Flags().BoolVarP(&optionPgDefaultString, "optionPgDefaultString", "", true, "option: pg default string ta removeying")
	CmdOrmGen.Flags().BoolVarP(&optionRemoveDefault, "optionRemoveDefault", "", true, "option: remove tag default")
	CmdOrmGen.Flags().BoolVarP(&optionRemoveGormTypeTag, "optionRemoveGormTypeTag", "", false, "option: remove gorm tag :type")
}

func Run(_ *cobra.Command, _ []string) {
	dbOpts := make([]gormGen.ModelOpt, 0)
	if optionUnderline != "" {
		dbOpts = append(dbOpts, gen.ModelOptionUnderline(optionUnderline))
	}
	if optionPgDefaultString {
		dbOpts = append(dbOpts, gen.ModelOptionPgDefaultString())
	}
	if optionRemoveDefault {
		dbOpts = append(dbOpts, gen.ModelOptionRemoveDefault())
	}
	if optionRemoveGormTypeTag {
		dbOpts = append(dbOpts, gen.ModelOptionRemoveGormTypeTag())
	}
	var tables []string
	if targetTables != "" {
		tables = strings.Split(targetTables, ",")
	}
	gen.NewGenerationDB(
		utils.NewDB(dsn),
		outPutPath,
		gen.WithDataMap(gen.DataTypeMap()),
		gen.WithTables(tables),
		gen.WithDBNameOpts(gen.DBNameOpts()),
		gen.WithDBOpts(dbOpts...),
	).Do()
}

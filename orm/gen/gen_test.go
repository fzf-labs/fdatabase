package gen

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.yc345.tv/backend/utils/v2/orm"
	"gorm.io/gen"
	"gorm.io/gen/field"
)

func TestGenerationPostgresWithOutRepo(t *testing.T) {
	// 初始化数据库
	db, err := orm.NewDBWithStruct(&orm.ORMConfig{
		User:     "postgres",
		Password: "7to12pg12",
		Host:     "10.8.8.110",
		Port:     5433,
		DBname:   "gorm_gen",
	})
	if err != nil {
		return
	}
	// 生成代码
	NewGenerationDB(
		db.Client,
		"./example/postgres/",
		WithOutRepo(),
		WithDBNameOpts(DBNameOpts()),
		WithTables([]string{"admin_demo", "admin_log_demo", "admin_role_demo"}),
		WithDataMap(DataTypeMap()), // 设置数据类型映射
		WithDBOpts(ModelOptionRemoveDefault(), ModelOptionPgDefaultString(), ModelOptionRemoveGormTypeTag(), ModelOptionUnderline("UL")), // 设置自定义选项
		WithFieldNullable(),
	).Do()
	assert.Equal(t, nil, err)
}

func TestGenerationPostgres(t *testing.T) {
	// 初始化数据库
	db, err := orm.NewDBWithStruct(&orm.ORMConfig{
		User:     "postgres",
		Password: "7to12pg12",
		Host:     "10.8.8.110",
		Port:     5433,
		DBname:   "gorm_gen",
	})
	if err != nil {
		return
	}
	// 生成代码
	NewGenerationDB(
		db.Client,
		"./example/postgres/",
		WithGenerateModel(func(g *gen.Generator) map[string]any { // 设置表关联关系(1对多,多对多...)
			adminLogDemo := g.GenerateModel("admin_log_demo")
			AdminRoleDemo := g.GenerateModel("admin_role_demo",
				gen.FieldRelate(field.Many2Many, "Admins", g.GenerateModel("admin_demo"),
					&field.RelateConfig{
						RelateSlicePointer: true,
						JSONTag:            JSONTagNameStrategy("Admins"),
						GORMTag:            field.GormTag{"joinForeignKey": []string{"role_id"}, "joinReferences": []string{"admin_id"}, "many2many": []string{"admin_to_role_demo"}},
					},
				),
			)
			adminDemo := g.GenerateModel("admin_demo",
				gen.FieldRelate(field.HasMany, "AdminLogDemos", adminLogDemo,
					&field.RelateConfig{
						RelateSlicePointer: true,
						JSONTag:            JSONTagNameStrategy("AdminLogDemos"),
						GORMTag:            field.GormTag{"foreignKey": []string{"admin_id"}},
					},
				),
				gen.FieldRelate(field.Many2Many, "AdminRoles", AdminRoleDemo,
					&field.RelateConfig{
						RelateSlicePointer: true,
						JSONTag:            JSONTagNameStrategy("AdminRoles"),
						GORMTag:            field.GormTag{"joinForeignKey": []string{"admin_id"}, "joinReferences": []string{"role_id"}, "many2many": []string{"admin_to_role_demo"}},
					},
				),
			)
			return map[string]any{
				"admin_demo":      adminDemo,
				"admin_log_demo":  adminLogDemo,
				"admin_role_demo": AdminRoleDemo,
			}
		}),
		WithDataMap(DataTypeMap()), // 设置数据类型映射
		WithDBOpts(ModelOptionRemoveDefault(), ModelOptionUnderline("UL")), // 设置自定义选项
	).Do()
	assert.Equal(t, nil, err)
}

func TestGenerationPostgresFieldNullable(t *testing.T) {
	// 初始化数据库
	db, err := orm.NewDBWithStruct(&orm.ORMConfig{
		User:     "postgres",
		Password: "7to12pg12",
		Host:     "10.8.8.110",
		Port:     5433,
		DBname:   "gorm_gen",
	})
	if err != nil {
		return
	}
	// 生成代码
	NewGenerationDB(
		db.Client,
		"./example/postgres/",
		WithGenerateModel(func(g *gen.Generator) map[string]any { // 设置表关联关系(1对多,多对多...)
			adminLogDemo := g.GenerateModel("admin_log_demo")
			AdminRoleDemo := g.GenerateModel("admin_role_demo",
				gen.FieldRelate(field.Many2Many, "Admins", g.GenerateModel("admin_demo"),
					&field.RelateConfig{
						RelateSlicePointer: true,
						JSONTag:            JSONTagNameStrategy("Admins"),
						GORMTag:            field.GormTag{"joinForeignKey": []string{"role_id"}, "joinReferences": []string{"admin_id"}, "many2many": []string{"admin_to_role_demo"}},
					},
				),
			)
			adminDemo := g.GenerateModel("admin_demo",
				gen.FieldRelate(field.HasMany, "AdminLogDemos", adminLogDemo,
					&field.RelateConfig{
						RelateSlicePointer: true,
						JSONTag:            JSONTagNameStrategy("AdminLogDemos"),
						GORMTag:            field.GormTag{"foreignKey": []string{"admin_id"}},
					},
				),
				gen.FieldRelate(field.Many2Many, "AdminRoles", AdminRoleDemo,
					&field.RelateConfig{
						RelateSlicePointer: true,
						JSONTag:            JSONTagNameStrategy("AdminRoles"),
						GORMTag:            field.GormTag{"joinForeignKey": []string{"admin_id"}, "joinReferences": []string{"role_id"}, "many2many": []string{"admin_to_role_demo"}},
					},
				),
			)
			return map[string]any{
				"admin_demo":      adminDemo,
				"admin_log_demo":  adminLogDemo,
				"admin_role_demo": AdminRoleDemo,
			}
		}),
		WithDataMap(DataTypeMap()), // 设置数据类型映射
		WithDBOpts(ModelOptionRemoveDefault(), ModelOptionUnderline("UL")), // 设置自定义选项
		WithFieldNullable(),
	).Do()
	assert.Equal(t, nil, err)
}

func TestNewGenerationPb(t *testing.T) {
	db, err := orm.NewDBWithStruct(&orm.ORMConfig{
		User:     "postgres",
		Password: "7to12pg12",
		Host:     "10.8.8.110",
		Port:     5433,
		DBname:   "gorm_gen",
	})
	if err != nil {
		return
	}
	NewGenerationPB(
		db.Client,
		"./example/postgres/pb",
		"api.gorm_gen.v1",
		"api/gorm_gen/v1;v1",
		WithPBOpts(ModelOptionRemoveDefault(), ModelOptionUnderline("ul_")),
	).Do()
	assert.Equal(t, nil, err)
}

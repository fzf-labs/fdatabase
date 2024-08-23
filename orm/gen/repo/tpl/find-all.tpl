// FindAll 查询所有数据
func ({{.firstTableChar}} *{{.upperTableName}}Repo) FindAll(ctx context.Context) ([]*{{.dbName}}_model.{{.upperTableName}}, error) {
	dao := {{.dbName}}_dao.Use({{.firstTableChar}}.db).{{.upperTableName}}
	result, err := dao.WithContext(ctx).Find()
	if err != nil {
		return nil, err
	}
	return result, nil
}
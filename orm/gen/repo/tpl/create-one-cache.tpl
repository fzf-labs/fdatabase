// CreateOneCache 创建一条数据
func ({{.firstTableChar}} *{{.upperTableName}}Repo) CreateOneCache(ctx context.Context, data *{{.dbName}}_model.{{.upperTableName}}) error {
	dao := {{.dbName}}_dao.Use({{.firstTableChar}}.db).{{.upperTableName}}
	err := dao.WithContext(ctx).Create(data)
	if err != nil {
		return err
	}
	err = {{.firstTableChar}}.DeleteUniqueIndexCache(ctx, []*{{.dbName}}_model.{{.upperTableName}}{data})
    if err != nil {
    	return err
    }
	return nil
}
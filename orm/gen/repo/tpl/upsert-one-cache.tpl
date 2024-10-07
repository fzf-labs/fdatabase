// UpsertOneCache Upsert一条数据, 并删除缓存
// Update all columns, except primary keys, to new value on conflict
func ({{.firstTableChar}} *{{.upperTableName}}Repo) UpsertOneCache(ctx context.Context, data *{{.dbName}}_model.{{.upperTableName}}) error {
	dao := {{.dbName}}_dao.Use({{.firstTableChar}}.db).{{.upperTableName}}
	err := dao.WithContext(ctx).Save(data)
	if err != nil {
		return err
	}
	err = {{.firstTableChar}}.DeleteIndexCache(ctx, []*{{.dbName}}_model.{{.upperTableName}}{data})
    if err != nil {
    	return err
    }
	return nil
}
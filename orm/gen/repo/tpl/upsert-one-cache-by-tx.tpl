// UpsertOneCacheByTx Upsert一条数据(事务), 并删除缓存
// Update all columns, except primary keys, to new value on conflict
func ({{.firstTableChar}} *{{.upperTableName}}Repo) UpsertOneCacheByTx(ctx context.Context, tx *{{.dbName}}_dao.Query, data *{{.dbName}}_model.{{.upperTableName}}) error {
	dao := tx.{{.upperTableName}}
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
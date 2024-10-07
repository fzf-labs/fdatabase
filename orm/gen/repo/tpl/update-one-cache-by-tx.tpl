// UpdateOneCacheByTx 更新一条数据(事务)，并删除缓存
// data 中主键字段必须有值，零值不会被更新
func ({{.firstTableChar}} *{{.upperTableName}}Repo) UpdateOneCacheByTx(ctx context.Context, tx *{{.dbName}}_dao.Query, data *{{.dbName}}_model.{{.upperTableName}}) error {
	dao := tx.{{.upperTableName}}
	_, err := dao.WithContext(ctx).Updates(data)
	if err != nil {
		return err
	}
    err = {{.firstTableChar}}.DeleteIndexCache(ctx, []*{{.dbName}}_model.{{.upperTableName}}{data})
    if err != nil {
        return err
    }
	return err
}
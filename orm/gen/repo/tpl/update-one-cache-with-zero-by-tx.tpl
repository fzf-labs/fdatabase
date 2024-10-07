// UpdateOneCacheWithZeroByTx 更新一条数据(事务),包含零值，并删除缓存
// data 中主键字段必须有值,并且会更新所有字段,包括零值
func ({{.firstTableChar}} *{{.upperTableName}}Repo) UpdateOneCacheWithZeroByTx(ctx context.Context, tx *{{.dbName}}_dao.Query, data *{{.dbName}}_model.{{.upperTableName}}) error {
	dao := tx.{{.upperTableName}}
	_, err := dao.WithContext(ctx).Select(dao.ALL.WithTable("")).Updates(data)
	if err != nil {
		return err
	}
    err = {{.firstTableChar}}.DeleteIndexCache(ctx, []*{{.dbName}}_model.{{.upperTableName}}{data})
    if err != nil {
        return err
    }
	return err
}
// UpdateOneCacheWithZeroByTx 更新一条数据(事务),包含零值，并删除缓存
UpdateOneCacheWithZeroByTx(ctx context.Context, tx *{{.dbName}}_dao.Query, data *{{.dbName}}_model.{{.upperTableName}}) error
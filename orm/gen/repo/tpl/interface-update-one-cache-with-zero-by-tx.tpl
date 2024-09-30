// UpdateOneCacheWithZeroByTx 更新一条数据,包含零值(事务)
UpdateOneCacheWithZeroByTx(ctx context.Context, tx *{{.dbName}}_dao.Query, data *{{.dbName}}_model.{{.upperTableName}}) error
// UpdateOneCacheByTx 更新一条数据(事务)
UpdateOneCacheByTx(ctx context.Context, tx *{{.dbName}}_dao.Query, data *{{.dbName}}_model.{{.upperTableName}}) error
// UpsertOneCacheByTx Upsert一条数据(事务)
UpsertOneCacheByTx(ctx context.Context, tx *{{.dbName}}_dao.Query, data *{{.dbName}}_model.{{.upperTableName}}) error
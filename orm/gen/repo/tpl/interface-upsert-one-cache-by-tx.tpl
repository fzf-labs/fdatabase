// UpsertOneCacheByTx Upsert一条数据(事务), 并删除缓存
// Update all columns, except primary keys, to new value on conflict
UpsertOneCacheByTx(ctx context.Context, tx *{{.dbName}}_dao.Query, data *{{.dbName}}_model.{{.upperTableName}}) error
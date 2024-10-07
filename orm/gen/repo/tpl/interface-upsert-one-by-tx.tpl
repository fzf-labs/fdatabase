// UpsertOneByTx Upsert一条数据(事务)
// Update all columns, except primary keys, to new value on conflict
UpsertOneByTx(ctx context.Context, tx *{{.dbName}}_dao.Query, data *{{.dbName}}_model.{{.upperTableName}}) error
// UpsertOneCache Upsert一条数据, 并删除缓存
// Update all columns, except primary keys, to new value on conflict
UpsertOneCache(ctx context.Context, data *{{.dbName}}_model.{{.upperTableName}}) error
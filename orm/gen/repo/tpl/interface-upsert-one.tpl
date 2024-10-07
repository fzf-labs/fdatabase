// UpsertOne Upsert一条数据
// Update all columns, except primary keys, to new value on conflict
UpsertOne(ctx context.Context, data *{{.dbName}}_model.{{.upperTableName}}) error
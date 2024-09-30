// UpsertOneCacheByFields Upsert一条数据，根据fields字段
UpsertOneCacheByFields(ctx context.Context, data *{{.dbName}}_model.{{.upperTableName}},fields []string) error
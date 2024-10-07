// UpdateOneCache 更新一条数据，并删除缓存
// data 中主键字段必须有值，零值不会被更新
UpdateOneCache(ctx context.Context, data *{{.dbName}}_model.{{.upperTableName}}) error
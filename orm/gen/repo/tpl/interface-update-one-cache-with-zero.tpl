// UpdateOneCacheWithZero 更新一条数据,包含零值，并删除缓存
UpdateOneCacheWithZero(ctx context.Context, data *{{.dbName}}_model.{{.upperTableName}}) error
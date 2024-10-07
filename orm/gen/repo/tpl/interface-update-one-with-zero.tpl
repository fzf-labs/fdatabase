// UpdateOneCacheWithZero 更新一条数据,包含零值，并删除缓存
// data 中主键字段必须有值,并且会更新所有字段,包括零值
UpdateOneWithZero(ctx context.Context, data *{{.dbName}}_model.{{.upperTableName}}) error
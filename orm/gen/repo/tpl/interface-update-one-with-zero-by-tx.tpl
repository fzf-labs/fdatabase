// UpdateOneCacheWithZeroByTx 更新一条数据(事务),包含零值，并删除缓存
// data 中主键字段必须有值,并且会更新所有字段,包括零值
UpdateOneWithZeroByTx(ctx context.Context, tx *{{.dbName}}_dao.Query, data *{{.dbName}}_model.{{.upperTableName}}) error
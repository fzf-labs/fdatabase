// UpdateOneCacheByTx 更新一条数据(事务)，并删除缓存
// data 中主键字段必须有值，零值不会被更新
UpdateOneCacheByTx(ctx context.Context, tx *{{.dbName}}_dao.Query, data *{{.dbName}}_model.{{.upperTableName}}) error
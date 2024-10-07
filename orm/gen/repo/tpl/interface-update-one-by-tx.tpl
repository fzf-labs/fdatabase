// UpdateOneByTx 更新一条数据(事务)
// data 中主键字段必须有值，零值不会被更新
UpdateOneByTx(ctx context.Context, tx *{{.dbName}}_dao.Query, data *{{.dbName}}_model.{{.upperTableName}}) error
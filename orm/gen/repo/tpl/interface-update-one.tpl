// UpdateOne 更新一条数据
// data 中主键字段必须有值，零值不会被更新
UpdateOne(ctx context.Context, data *{{.dbName}}_model.{{.upperTableName}}) error
// CreateBatchCacheByTx 批量创建数据(事务)
CreateBatchCacheByTx(ctx context.Context,tx *{{.dbName}}_dao.Query, data []*{{.dbName}}_model.{{.upperTableName}}, batchSize int) error
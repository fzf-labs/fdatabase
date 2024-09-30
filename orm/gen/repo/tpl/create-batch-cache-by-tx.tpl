// CreateBatchCacheByTx 批量创建数据(事务)
func ({{.firstTableChar}} *{{.upperTableName}}Repo) CreateBatchCacheByTx(ctx context.Context,tx *{{.dbName}}_dao.Query, data []*{{.dbName}}_model.{{.upperTableName}}, batchSize int) error {
	dao := tx.{{.upperTableName}}
	err := dao.WithContext(ctx).CreateInBatches(data,batchSize)
	if err != nil {
		return err
	}
	err = {{.firstTableChar}}.DeleteUniqueIndexCache(ctx, data)
	if err != nil {
		return err
	}
	return nil
}
// UpsertOneCacheByFieldsTx 根据fields字段Upsert一条数据(事务), 并删除缓存
func ({{.firstTableChar}} *{{.upperTableName}}Repo) UpsertOneCacheByFieldsTx(ctx context.Context,tx *{{.dbName}}_dao.Query, data *{{.dbName}}_model.{{.upperTableName}},fields []string) error {
	if len(fields) == 0 {
        return errors.New("UpsertOneByFieldsTx fields is empty")
    }
	columns := make([]clause.Column, 0)
	for _, v := range fields {
		columns = append(columns, clause.Column{Name: v})
	}
	dao := tx.{{.upperTableName}}
	err := dao.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   columns,
		UpdateAll: true,
	}).Create(data)
	if err != nil {
		return err
	}
	err = {{.firstTableChar}}.DeleteIndexCache(ctx, []*{{.dbName}}_model.{{.upperTableName}}{data})
	if err != nil {
		return err
	}
	return nil
}
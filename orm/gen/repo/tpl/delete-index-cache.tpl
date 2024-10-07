// DeleteUniqueIndexCache 删除索引存在的缓存
func ({{.firstTableChar}} *{{.upperTableName}}Repo) DeleteIndexCache(ctx context.Context, data []*{{.dbName}}_model.{{.upperTableName}}) error {
	keys := make([]string, 0)
	for _, v := range data {
		keys = append(
		    keys,
	    	{{.cacheDelKeys}}
		)
	}
	err := {{.firstTableChar}}.cache.DelBatch(ctx, keys)
	if err != nil {
		return err
	}
	return nil
}
// DeleteUniqueIndexCache 删除索引存在的缓存
func ({{.firstTableChar}} *{{.upperTableName}}Repo) DeleteIndexCache(ctx context.Context, data ...*{{.dbName}}_model.{{.upperTableName}}) error {
	KeyMap := make(map[string]struct{})
	for _, v := range data {
		if v != nil {
			{{.cacheDelKeys}}
		}
	}
	keys := make([]string, 0, len(KeyMap))
	for k := range KeyMap {
		keys = append(keys, k)
	}
	err := {{.firstTableChar}}.cache.DelBatch(ctx, keys)
	if err != nil {
		return err
	}
	return nil
}
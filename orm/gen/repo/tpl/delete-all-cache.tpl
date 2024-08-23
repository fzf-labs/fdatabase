// DeleteAllCache 删除所有数据缓存
func ({{.firstTableChar}} *{{.upperTableName}}Repo) DeleteAllCache(ctx context.Context) error {
	cacheKey := {{.firstTableChar}}.cache.Key(Cache{{.upperTableName}}All)
	err := {{.firstTableChar}}.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}
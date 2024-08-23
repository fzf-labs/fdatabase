// FindAllCache 查询所有数据并设置缓存
func ({{.firstTableChar}} *{{.upperTableName}}Repo) FindAllCache(ctx context.Context) ([]*{{.dbName}}_model.{{.upperTableName}}, error) {
	resp := make([]*{{.dbName}}_model.{{.upperTableName}}, 0)
	cacheKey := {{.firstTableChar}}.cache.Key(Cache{{.upperTableName}}All)
	cacheValue, err := {{.firstTableChar}}.cache.Fetch(ctx, cacheKey, func() (string, error) {
		dao := {{.dbName}}_dao.Use({{.firstTableChar}}.db).{{.upperTableName}}
		result, err := dao.WithContext(ctx).Find()
		if err != nil {
			return "", err
		}
        marshal, err := {{.firstTableChar}}.encoding.Marshal(result)
        if err != nil {
            return "", err
        }
		return string(marshal), nil
	})
	if err != nil {
		return nil, err
	}
	if cacheValue != "" {
		err = {{.firstTableChar}}.encoding.Unmarshal([]byte(cacheValue), resp)
		if err != nil {
			return nil, err
		}
	}
	return resp, nil
}
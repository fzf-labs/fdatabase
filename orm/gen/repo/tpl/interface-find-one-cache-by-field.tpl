// FindOneCacheBy{{.upperField}} 根据{{.lowerField}}查询一条数据，并设置缓存
FindOneCacheBy{{.upperField}}(ctx context.Context, {{.lowerField}} {{.dataType}}) (*{{.dbName}}_model.{{.upperTableName}}, error)
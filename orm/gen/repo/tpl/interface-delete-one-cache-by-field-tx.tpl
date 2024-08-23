// DeleteOneCacheBy{{.upperField}}Tx 根据{{.lowerField}}删除一条数据并清理缓存(事务)
DeleteOneCacheBy{{.upperField}}Tx(ctx context.Context,tx *{{.dbName}}_dao.Query, {{.lowerField}} {{.dataType}}) error
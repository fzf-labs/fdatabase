// UpdateBatchBy{{.upperSinglePrimaryKeyPlural}} 根据主键{{.upperSinglePrimaryKeyPlural}}批量更新
UpdateBatchBy{{.upperSinglePrimaryKeyPlural}}(ctx context.Context, {{.lowerSinglePrimaryKeyPlural}} []{{.dataTypeSinglePrimaryKey}}, data map[string]interface{}) error
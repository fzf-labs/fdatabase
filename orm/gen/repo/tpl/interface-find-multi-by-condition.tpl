// FindMultiByCondition 根据自定义条件查询数据
FindMultiByCondition(ctx context.Context, conditionReq *condition.Req) ([]*{{.dbName}}_model.{{.upperTableName}}, *condition.Reply, error)
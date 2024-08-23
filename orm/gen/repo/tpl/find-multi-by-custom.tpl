// Deprecated
// 请使用FindMultiByCondition替代
func ({{.firstTableChar}} *{{.upperTableName}}Repo) FindMultiByCustom(ctx context.Context, customReq *custom.Req) ([]*{{.dbName}}_model.{{.upperTableName}}, *custom.Reply, error) {
	result := make([]*{{.dbName}}_model.{{.upperTableName}}, 0)
	var total int64
	whereExpressions, orderExpressions, err := customReq.ConvertToGormExpression({{.dbName}}_model.{{.upperTableName}}{})
	if err != nil {
		return result, nil, err
	}
	err = {{.firstTableChar}}.db.WithContext(ctx).Model(&{{.dbName}}_model.{{.upperTableName}}{}).Select([]string{"*"}).Clauses(whereExpressions...).Count(&total).Error
	if err != nil {
		return result, nil, err
	}
	if total == 0 {
		return result, nil, nil
	}
	customReply,err := customReq.ConvertToPage(int(total))
	if err != nil {
		return result, nil, err
	}
	query := {{.firstTableChar}}.db.WithContext(ctx).Model(&{{.dbName}}_model.{{.upperTableName}}{}).Clauses(whereExpressions...).Clauses(orderExpressions...)
	if customReply.Offset != 0 {
		query = query.Offset(customReply.Offset)
	}
	if customReply.Limit != 0 {
		query = query.Limit(customReply.Limit)
	}
	err = query.Find(&result).Error
	if err != nil {
		return result, nil, err
	}
	return result, customReply, err
}

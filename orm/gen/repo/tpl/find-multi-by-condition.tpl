// FindMultiByCondition 自定义查询数据(通用)
func ({{.firstTableChar}} *{{.upperTableName}}Repo) FindMultiByCondition(ctx context.Context, conditionReq *condition.Req) ([]*{{.dbName}}_model.{{.upperTableName}}, *condition.Reply, error) {
	result := make([]*{{.dbName}}_model.{{.upperTableName}}, 0)
	var total int64
	whereExpressions, orderExpressions, err := conditionReq.ConvertToGormExpression({{.dbName}}_model.{{.upperTableName}}{})
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
	conditionReply,err := conditionReq.ConvertToPage(int32(total))
	if err != nil {
		return result, nil, err
	}
	query := {{.firstTableChar}}.db.WithContext(ctx).Model(&{{.dbName}}_model.{{.upperTableName}}{}).Clauses(whereExpressions...).Clauses(orderExpressions...)
	if conditionReply.Page != 0 && conditionReply.PageSize != 0 {
		query = query.Offset(int((conditionReply.Page - 1) * conditionReply.PageSize))
		query = query.Limit(int(conditionReply.PageSize))
	}
	err = query.Find(&result).Error
	if err != nil {
		return result, nil, err
	}
	return result, conditionReply, err
}

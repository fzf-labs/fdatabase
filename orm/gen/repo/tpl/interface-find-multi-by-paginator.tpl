// Deprecated
// 请使用FindMultiByCondition替代
FindMultiByPaginator(ctx context.Context, paginatorReq *orm.PaginatorReq) ([]*{{.dbName}}_model.{{.upperTableName}}, *orm.PaginatorReply, error)
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package gorm_gen_dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"gitlab.yc345.tv/backend/utils/v2/orm/gen/example/postgres/gorm_gen_model"
)

func newAdminLogDemo(db *gorm.DB, opts ...gen.DOOption) adminLogDemo {
	_adminLogDemo := adminLogDemo{}

	_adminLogDemo.adminLogDemoDo.UseDB(db, opts...)
	_adminLogDemo.adminLogDemoDo.UseModel(&gorm_gen_model.AdminLogDemo{})

	tableName := _adminLogDemo.adminLogDemoDo.TableName()
	_adminLogDemo.ALL = field.NewAsterisk(tableName)
	_adminLogDemo.ID = field.NewString(tableName, "id")
	_adminLogDemo.AdminID = field.NewString(tableName, "admin_id")
	_adminLogDemo.IP = field.NewString(tableName, "ip")
	_adminLogDemo.URI = field.NewString(tableName, "uri")
	_adminLogDemo.Useragent = field.NewString(tableName, "useragent")
	_adminLogDemo.Header = field.NewField(tableName, "header")
	_adminLogDemo.Req = field.NewField(tableName, "req")
	_adminLogDemo.Resp = field.NewField(tableName, "resp")
	_adminLogDemo.CreatedAt = field.NewTime(tableName, "created_at")

	_adminLogDemo.fillFieldMap()

	return _adminLogDemo
}

type adminLogDemo struct {
	adminLogDemoDo adminLogDemoDo

	ALL       field.Asterisk
	ID        field.String // 编号
	AdminID   field.String // 管理员ID
	IP        field.String // ip
	URI       field.String // 请求路径
	Useragent field.String // 浏览器标识
	Header    field.Field  // header
	Req       field.Field  // 请求数据
	Resp      field.Field  // 响应数据
	CreatedAt field.Time   // 创建时间

	fieldMap map[string]field.Expr
}

func (a adminLogDemo) Table(newTableName string) *adminLogDemo {
	a.adminLogDemoDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a adminLogDemo) As(alias string) *adminLogDemo {
	a.adminLogDemoDo.DO = *(a.adminLogDemoDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *adminLogDemo) updateTableName(table string) *adminLogDemo {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewString(table, "id")
	a.AdminID = field.NewString(table, "admin_id")
	a.IP = field.NewString(table, "ip")
	a.URI = field.NewString(table, "uri")
	a.Useragent = field.NewString(table, "useragent")
	a.Header = field.NewField(table, "header")
	a.Req = field.NewField(table, "req")
	a.Resp = field.NewField(table, "resp")
	a.CreatedAt = field.NewTime(table, "created_at")

	a.fillFieldMap()

	return a
}

func (a *adminLogDemo) WithContext(ctx context.Context) *adminLogDemoDo {
	return a.adminLogDemoDo.WithContext(ctx)
}

func (a adminLogDemo) TableName() string { return a.adminLogDemoDo.TableName() }

func (a adminLogDemo) Alias() string { return a.adminLogDemoDo.Alias() }

func (a adminLogDemo) Columns(cols ...field.Expr) gen.Columns {
	return a.adminLogDemoDo.Columns(cols...)
}

func (a *adminLogDemo) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *adminLogDemo) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 9)
	a.fieldMap["id"] = a.ID
	a.fieldMap["admin_id"] = a.AdminID
	a.fieldMap["ip"] = a.IP
	a.fieldMap["uri"] = a.URI
	a.fieldMap["useragent"] = a.Useragent
	a.fieldMap["header"] = a.Header
	a.fieldMap["req"] = a.Req
	a.fieldMap["resp"] = a.Resp
	a.fieldMap["created_at"] = a.CreatedAt
}

func (a adminLogDemo) clone(db *gorm.DB) adminLogDemo {
	a.adminLogDemoDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a adminLogDemo) replaceDB(db *gorm.DB) adminLogDemo {
	a.adminLogDemoDo.ReplaceDB(db)
	return a
}

type adminLogDemoDo struct{ gen.DO }

func (a adminLogDemoDo) Debug() *adminLogDemoDo {
	return a.withDO(a.DO.Debug())
}

func (a adminLogDemoDo) WithContext(ctx context.Context) *adminLogDemoDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a adminLogDemoDo) ReadDB() *adminLogDemoDo {
	return a.Clauses(dbresolver.Read)
}

func (a adminLogDemoDo) WriteDB() *adminLogDemoDo {
	return a.Clauses(dbresolver.Write)
}

func (a adminLogDemoDo) Session(config *gorm.Session) *adminLogDemoDo {
	return a.withDO(a.DO.Session(config))
}

func (a adminLogDemoDo) Clauses(conds ...clause.Expression) *adminLogDemoDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a adminLogDemoDo) Returning(value interface{}, columns ...string) *adminLogDemoDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a adminLogDemoDo) Not(conds ...gen.Condition) *adminLogDemoDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a adminLogDemoDo) Or(conds ...gen.Condition) *adminLogDemoDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a adminLogDemoDo) Select(conds ...field.Expr) *adminLogDemoDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a adminLogDemoDo) Where(conds ...gen.Condition) *adminLogDemoDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a adminLogDemoDo) Order(conds ...field.Expr) *adminLogDemoDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a adminLogDemoDo) Distinct(cols ...field.Expr) *adminLogDemoDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a adminLogDemoDo) Omit(cols ...field.Expr) *adminLogDemoDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a adminLogDemoDo) Join(table schema.Tabler, on ...field.Expr) *adminLogDemoDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a adminLogDemoDo) LeftJoin(table schema.Tabler, on ...field.Expr) *adminLogDemoDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a adminLogDemoDo) RightJoin(table schema.Tabler, on ...field.Expr) *adminLogDemoDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a adminLogDemoDo) Group(cols ...field.Expr) *adminLogDemoDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a adminLogDemoDo) Having(conds ...gen.Condition) *adminLogDemoDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a adminLogDemoDo) Limit(limit int) *adminLogDemoDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a adminLogDemoDo) Offset(offset int) *adminLogDemoDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a adminLogDemoDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *adminLogDemoDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a adminLogDemoDo) Unscoped() *adminLogDemoDo {
	return a.withDO(a.DO.Unscoped())
}

func (a adminLogDemoDo) Create(values ...*gorm_gen_model.AdminLogDemo) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a adminLogDemoDo) CreateInBatches(values []*gorm_gen_model.AdminLogDemo, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a adminLogDemoDo) Save(values ...*gorm_gen_model.AdminLogDemo) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a adminLogDemoDo) First() (*gorm_gen_model.AdminLogDemo, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*gorm_gen_model.AdminLogDemo), nil
	}
}

func (a adminLogDemoDo) Take() (*gorm_gen_model.AdminLogDemo, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*gorm_gen_model.AdminLogDemo), nil
	}
}

func (a adminLogDemoDo) Last() (*gorm_gen_model.AdminLogDemo, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*gorm_gen_model.AdminLogDemo), nil
	}
}

func (a adminLogDemoDo) Find() ([]*gorm_gen_model.AdminLogDemo, error) {
	result, err := a.DO.Find()
	return result.([]*gorm_gen_model.AdminLogDemo), err
}

func (a adminLogDemoDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*gorm_gen_model.AdminLogDemo, err error) {
	buf := make([]*gorm_gen_model.AdminLogDemo, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a adminLogDemoDo) FindInBatches(result *[]*gorm_gen_model.AdminLogDemo, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a adminLogDemoDo) Attrs(attrs ...field.AssignExpr) *adminLogDemoDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a adminLogDemoDo) Assign(attrs ...field.AssignExpr) *adminLogDemoDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a adminLogDemoDo) Joins(fields ...field.RelationField) *adminLogDemoDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a adminLogDemoDo) Preload(fields ...field.RelationField) *adminLogDemoDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a adminLogDemoDo) FirstOrInit() (*gorm_gen_model.AdminLogDemo, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*gorm_gen_model.AdminLogDemo), nil
	}
}

func (a adminLogDemoDo) FirstOrCreate() (*gorm_gen_model.AdminLogDemo, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*gorm_gen_model.AdminLogDemo), nil
	}
}

func (a adminLogDemoDo) FindByPage(offset int, limit int) (result []*gorm_gen_model.AdminLogDemo, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a adminLogDemoDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a adminLogDemoDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a adminLogDemoDo) Delete(models ...*gorm_gen_model.AdminLogDemo) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *adminLogDemoDo) withDO(do gen.Dao) *adminLogDemoDo {
	a.DO = *do.(*gen.DO)
	return a
}

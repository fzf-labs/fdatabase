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

func newAdminRoleDemo(db *gorm.DB, opts ...gen.DOOption) adminRoleDemo {
	_adminRoleDemo := adminRoleDemo{}

	_adminRoleDemo.adminRoleDemoDo.UseDB(db, opts...)
	_adminRoleDemo.adminRoleDemoDo.UseModel(&gorm_gen_model.AdminRoleDemo{})

	tableName := _adminRoleDemo.adminRoleDemoDo.TableName()
	_adminRoleDemo.ALL = field.NewAsterisk(tableName)
	_adminRoleDemo.ID = field.NewString(tableName, "id")
	_adminRoleDemo.Pid = field.NewString(tableName, "pid")
	_adminRoleDemo.Name = field.NewString(tableName, "name")
	_adminRoleDemo.Remark = field.NewString(tableName, "remark")
	_adminRoleDemo.Status = field.NewInt16(tableName, "status")
	_adminRoleDemo.Sort = field.NewInt64(tableName, "sort")
	_adminRoleDemo.CreatedAt = field.NewTime(tableName, "created_at")
	_adminRoleDemo.UpdatedAt = field.NewTime(tableName, "updated_at")
	_adminRoleDemo.DeletedAt = field.NewField(tableName, "deleted_at")
	_adminRoleDemo.Admins = adminRoleDemoManyToManyAdmins{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Admins", "gorm_gen_model.AdminDemo"),
	}

	_adminRoleDemo.fillFieldMap()

	return _adminRoleDemo
}

type adminRoleDemo struct {
	adminRoleDemoDo adminRoleDemoDo

	ALL       field.Asterisk
	ID        field.String // 编号
	Pid       field.String // 父级id
	Name      field.String // 名称
	Remark    field.String // 备注
	Status    field.Int16  // 0=禁用 1=开启
	Sort      field.Int64  // 排序值
	CreatedAt field.Time   // 创建时间
	UpdatedAt field.Time   // 更新时间
	DeletedAt field.Field  // 删除时间
	Admins    adminRoleDemoManyToManyAdmins

	fieldMap map[string]field.Expr
}

func (a adminRoleDemo) Table(newTableName string) *adminRoleDemo {
	a.adminRoleDemoDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a adminRoleDemo) As(alias string) *adminRoleDemo {
	a.adminRoleDemoDo.DO = *(a.adminRoleDemoDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *adminRoleDemo) updateTableName(table string) *adminRoleDemo {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewString(table, "id")
	a.Pid = field.NewString(table, "pid")
	a.Name = field.NewString(table, "name")
	a.Remark = field.NewString(table, "remark")
	a.Status = field.NewInt16(table, "status")
	a.Sort = field.NewInt64(table, "sort")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.UpdatedAt = field.NewTime(table, "updated_at")
	a.DeletedAt = field.NewField(table, "deleted_at")

	a.fillFieldMap()

	return a
}

func (a *adminRoleDemo) WithContext(ctx context.Context) *adminRoleDemoDo {
	return a.adminRoleDemoDo.WithContext(ctx)
}

func (a adminRoleDemo) TableName() string { return a.adminRoleDemoDo.TableName() }

func (a adminRoleDemo) Alias() string { return a.adminRoleDemoDo.Alias() }

func (a adminRoleDemo) Columns(cols ...field.Expr) gen.Columns {
	return a.adminRoleDemoDo.Columns(cols...)
}

func (a *adminRoleDemo) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *adminRoleDemo) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 10)
	a.fieldMap["id"] = a.ID
	a.fieldMap["pid"] = a.Pid
	a.fieldMap["name"] = a.Name
	a.fieldMap["remark"] = a.Remark
	a.fieldMap["status"] = a.Status
	a.fieldMap["sort"] = a.Sort
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["updated_at"] = a.UpdatedAt
	a.fieldMap["deleted_at"] = a.DeletedAt

}

func (a adminRoleDemo) clone(db *gorm.DB) adminRoleDemo {
	a.adminRoleDemoDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a adminRoleDemo) replaceDB(db *gorm.DB) adminRoleDemo {
	a.adminRoleDemoDo.ReplaceDB(db)
	return a
}

type adminRoleDemoManyToManyAdmins struct {
	db *gorm.DB

	field.RelationField
}

func (a adminRoleDemoManyToManyAdmins) Where(conds ...field.Expr) *adminRoleDemoManyToManyAdmins {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a adminRoleDemoManyToManyAdmins) WithContext(ctx context.Context) *adminRoleDemoManyToManyAdmins {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a adminRoleDemoManyToManyAdmins) Session(session *gorm.Session) *adminRoleDemoManyToManyAdmins {
	a.db = a.db.Session(session)
	return &a
}

func (a adminRoleDemoManyToManyAdmins) Model(m *gorm_gen_model.AdminRoleDemo) *adminRoleDemoManyToManyAdminsTx {
	return &adminRoleDemoManyToManyAdminsTx{a.db.Model(m).Association(a.Name())}
}

type adminRoleDemoManyToManyAdminsTx struct{ tx *gorm.Association }

func (a adminRoleDemoManyToManyAdminsTx) Find() (result []*gorm_gen_model.AdminDemo, err error) {
	return result, a.tx.Find(&result)
}

func (a adminRoleDemoManyToManyAdminsTx) Append(values ...*gorm_gen_model.AdminDemo) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a adminRoleDemoManyToManyAdminsTx) Replace(values ...*gorm_gen_model.AdminDemo) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a adminRoleDemoManyToManyAdminsTx) Delete(values ...*gorm_gen_model.AdminDemo) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a adminRoleDemoManyToManyAdminsTx) Clear() error {
	return a.tx.Clear()
}

func (a adminRoleDemoManyToManyAdminsTx) Count() int64 {
	return a.tx.Count()
}

type adminRoleDemoDo struct{ gen.DO }

func (a adminRoleDemoDo) Debug() *adminRoleDemoDo {
	return a.withDO(a.DO.Debug())
}

func (a adminRoleDemoDo) WithContext(ctx context.Context) *adminRoleDemoDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a adminRoleDemoDo) ReadDB() *adminRoleDemoDo {
	return a.Clauses(dbresolver.Read)
}

func (a adminRoleDemoDo) WriteDB() *adminRoleDemoDo {
	return a.Clauses(dbresolver.Write)
}

func (a adminRoleDemoDo) Session(config *gorm.Session) *adminRoleDemoDo {
	return a.withDO(a.DO.Session(config))
}

func (a adminRoleDemoDo) Clauses(conds ...clause.Expression) *adminRoleDemoDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a adminRoleDemoDo) Returning(value interface{}, columns ...string) *adminRoleDemoDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a adminRoleDemoDo) Not(conds ...gen.Condition) *adminRoleDemoDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a adminRoleDemoDo) Or(conds ...gen.Condition) *adminRoleDemoDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a adminRoleDemoDo) Select(conds ...field.Expr) *adminRoleDemoDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a adminRoleDemoDo) Where(conds ...gen.Condition) *adminRoleDemoDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a adminRoleDemoDo) Order(conds ...field.Expr) *adminRoleDemoDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a adminRoleDemoDo) Distinct(cols ...field.Expr) *adminRoleDemoDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a adminRoleDemoDo) Omit(cols ...field.Expr) *adminRoleDemoDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a adminRoleDemoDo) Join(table schema.Tabler, on ...field.Expr) *adminRoleDemoDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a adminRoleDemoDo) LeftJoin(table schema.Tabler, on ...field.Expr) *adminRoleDemoDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a adminRoleDemoDo) RightJoin(table schema.Tabler, on ...field.Expr) *adminRoleDemoDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a adminRoleDemoDo) Group(cols ...field.Expr) *adminRoleDemoDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a adminRoleDemoDo) Having(conds ...gen.Condition) *adminRoleDemoDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a adminRoleDemoDo) Limit(limit int) *adminRoleDemoDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a adminRoleDemoDo) Offset(offset int) *adminRoleDemoDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a adminRoleDemoDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *adminRoleDemoDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a adminRoleDemoDo) Unscoped() *adminRoleDemoDo {
	return a.withDO(a.DO.Unscoped())
}

func (a adminRoleDemoDo) Create(values ...*gorm_gen_model.AdminRoleDemo) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a adminRoleDemoDo) CreateInBatches(values []*gorm_gen_model.AdminRoleDemo, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a adminRoleDemoDo) Save(values ...*gorm_gen_model.AdminRoleDemo) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a adminRoleDemoDo) First() (*gorm_gen_model.AdminRoleDemo, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*gorm_gen_model.AdminRoleDemo), nil
	}
}

func (a adminRoleDemoDo) Take() (*gorm_gen_model.AdminRoleDemo, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*gorm_gen_model.AdminRoleDemo), nil
	}
}

func (a adminRoleDemoDo) Last() (*gorm_gen_model.AdminRoleDemo, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*gorm_gen_model.AdminRoleDemo), nil
	}
}

func (a adminRoleDemoDo) Find() ([]*gorm_gen_model.AdminRoleDemo, error) {
	result, err := a.DO.Find()
	return result.([]*gorm_gen_model.AdminRoleDemo), err
}

func (a adminRoleDemoDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*gorm_gen_model.AdminRoleDemo, err error) {
	buf := make([]*gorm_gen_model.AdminRoleDemo, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a adminRoleDemoDo) FindInBatches(result *[]*gorm_gen_model.AdminRoleDemo, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a adminRoleDemoDo) Attrs(attrs ...field.AssignExpr) *adminRoleDemoDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a adminRoleDemoDo) Assign(attrs ...field.AssignExpr) *adminRoleDemoDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a adminRoleDemoDo) Joins(fields ...field.RelationField) *adminRoleDemoDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a adminRoleDemoDo) Preload(fields ...field.RelationField) *adminRoleDemoDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a adminRoleDemoDo) FirstOrInit() (*gorm_gen_model.AdminRoleDemo, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*gorm_gen_model.AdminRoleDemo), nil
	}
}

func (a adminRoleDemoDo) FirstOrCreate() (*gorm_gen_model.AdminRoleDemo, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*gorm_gen_model.AdminRoleDemo), nil
	}
}

func (a adminRoleDemoDo) FindByPage(offset int, limit int) (result []*gorm_gen_model.AdminRoleDemo, count int64, err error) {
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

func (a adminRoleDemoDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a adminRoleDemoDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a adminRoleDemoDo) Delete(models ...*gorm_gen_model.AdminRoleDemo) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *adminRoleDemoDo) withDO(do gen.Dao) *adminRoleDemoDo {
	a.DO = *do.(*gen.DO)
	return a
}

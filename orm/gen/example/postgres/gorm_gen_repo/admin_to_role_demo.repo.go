// Code generated by gen/repo. DO NOT EDIT.
// Code generated by gen/repo. DO NOT EDIT.
// Code generated by gen/repo. DO NOT EDIT.

package gorm_gen_repo

import (
	"context"
	"errors"
	"reflect"
	"strings"

	"github.com/fzf-labs/fdatabase/orm/condition"
	"github.com/fzf-labs/fdatabase/orm/dbcache"
	"github.com/fzf-labs/fdatabase/orm/encoding"
	"github.com/fzf-labs/fdatabase/orm/gen/config"
	"github.com/fzf-labs/fdatabase/orm/gen/example/postgres/gorm_gen_dao"
	"github.com/fzf-labs/fdatabase/orm/gen/example/postgres/gorm_gen_model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var _ IAdminToRoleDemoRepo = (*AdminToRoleDemoRepo)(nil)

var (
	CacheAdminToRoleDemoByAdminIDRoleIDPrefix = "DBCache:gorm_gen:AdminToRoleDemoByAdminIDRoleID"
	CacheAdminToRoleDemoByRoleIDAdminIDPrefix = "DBCache:gorm_gen:AdminToRoleDemoByRoleIDAdminID"
	CacheAdminToRoleDemoByAdminIDPrefix       = "DBCache:gorm_gen:AdminToRoleDemoByAdminID"
	CacheAdminToRoleDemoByRoleIDPrefix        = "DBCache:gorm_gen:AdminToRoleDemoByRoleID"
)

type (
	IAdminToRoleDemoRepo interface {
		// DeepCopy 深拷贝
		DeepCopy(data *gorm_gen_model.AdminToRoleDemo) *gorm_gen_model.AdminToRoleDemo
		// CreateOne 创建一条数据
		CreateOne(ctx context.Context, data *gorm_gen_model.AdminToRoleDemo) error
		// CreateOneCache 创建一条数据, 并删除缓存
		CreateOneCache(ctx context.Context, data *gorm_gen_model.AdminToRoleDemo) error
		// CreateOneByTx 创建一条数据(事务)
		CreateOneByTx(ctx context.Context, tx *gorm_gen_dao.Query, data *gorm_gen_model.AdminToRoleDemo) error
		// CreateOneCacheByTx 创建一条数据(事务), 并删除缓存
		CreateOneCacheByTx(ctx context.Context, tx *gorm_gen_dao.Query, data *gorm_gen_model.AdminToRoleDemo) error
		// CreateBatch 批量创建数据
		CreateBatch(ctx context.Context, data []*gorm_gen_model.AdminToRoleDemo, batchSize int) error
		// CreateBatchCache 批量创建数据, 并删除缓存
		CreateBatchCache(ctx context.Context, data []*gorm_gen_model.AdminToRoleDemo, batchSize int) error
		// CreateBatchByTx 批量创建数据(事务)
		CreateBatchByTx(ctx context.Context, tx *gorm_gen_dao.Query, data []*gorm_gen_model.AdminToRoleDemo, batchSize int) error
		// CreateBatchCacheByTx 批量创建数据(事务), 并删除缓存
		CreateBatchCacheByTx(ctx context.Context, tx *gorm_gen_dao.Query, data []*gorm_gen_model.AdminToRoleDemo, batchSize int) error
		// UpsertOneByFields 根据fields字段Upsert一条数据
		UpsertOneByFields(ctx context.Context, data *gorm_gen_model.AdminToRoleDemo, fields []string) error
		// UpsertOneCacheByFields 根据fields字段Upsert一条数据, 并删除缓存
		UpsertOneCacheByFields(ctx context.Context, data *gorm_gen_model.AdminToRoleDemo, fields []string) error
		// UpsertOneByFieldsTx 根据fields字段Upsert一条数据(事务)
		UpsertOneByFieldsTx(ctx context.Context, tx *gorm_gen_dao.Query, data *gorm_gen_model.AdminToRoleDemo, fields []string) error
		// UpsertOneCacheByFieldsTx 根据fields字段Upsert一条数据(事务), 并删除缓存
		UpsertOneCacheByFieldsTx(ctx context.Context, tx *gorm_gen_dao.Query, data *gorm_gen_model.AdminToRoleDemo, fields []string) error
		// FindMultiByAdminIDRoleID 根据AdminIDRoleID查询多条数据，并设置缓存
		FindMultiByAdminIDRoleID(ctx context.Context, adminID string, roleID string) ([]*gorm_gen_model.AdminToRoleDemo, error)
		// FindMultiCacheByAdminIDRoleID 根据AdminIDRoleID查询多条数据，并设置缓存
		FindMultiCacheByAdminIDRoleID(ctx context.Context, adminID string, roleID string) ([]*gorm_gen_model.AdminToRoleDemo, error)
		// FindMultiByRoleIDAdminID 根据RoleIDAdminID查询多条数据，并设置缓存
		FindMultiByRoleIDAdminID(ctx context.Context, roleID string, adminID string) ([]*gorm_gen_model.AdminToRoleDemo, error)
		// FindMultiCacheByRoleIDAdminID 根据RoleIDAdminID查询多条数据，并设置缓存
		FindMultiCacheByRoleIDAdminID(ctx context.Context, roleID string, adminID string) ([]*gorm_gen_model.AdminToRoleDemo, error)
		// FindMultiByAdminID 根据adminID查询多条数据
		FindMultiByAdminID(ctx context.Context, adminID string) ([]*gorm_gen_model.AdminToRoleDemo, error)
		// FindMultiCacheByAdminID 根据adminID查询多条数据并设置缓存
		FindMultiCacheByAdminID(ctx context.Context, adminID string) ([]*gorm_gen_model.AdminToRoleDemo, error)
		// FindMultiByAdminIDS 根据adminIDS查询多条数据
		FindMultiByAdminIDS(ctx context.Context, adminIDS []string) ([]*gorm_gen_model.AdminToRoleDemo, error)
		// FindMultiCacheByAdminIDS 根据adminIDS查询多条数据，并设置缓存
		FindMultiCacheByAdminIDS(ctx context.Context, adminIDS []string) ([]*gorm_gen_model.AdminToRoleDemo, error)
		// FindMultiByRoleID 根据roleID查询多条数据
		FindMultiByRoleID(ctx context.Context, roleID string) ([]*gorm_gen_model.AdminToRoleDemo, error)
		// FindMultiCacheByRoleID 根据roleID查询多条数据并设置缓存
		FindMultiCacheByRoleID(ctx context.Context, roleID string) ([]*gorm_gen_model.AdminToRoleDemo, error)
		// FindMultiByRoleIDS 根据roleIDS查询多条数据
		FindMultiByRoleIDS(ctx context.Context, roleIDS []string) ([]*gorm_gen_model.AdminToRoleDemo, error)
		// FindMultiCacheByRoleIDS 根据roleIDS查询多条数据，并设置缓存
		FindMultiCacheByRoleIDS(ctx context.Context, roleIDS []string) ([]*gorm_gen_model.AdminToRoleDemo, error)
		// FindMultiByCondition 根据自定义条件查询数据
		FindMultiByCondition(ctx context.Context, conditionReq *condition.Req) ([]*gorm_gen_model.AdminToRoleDemo, *condition.Reply, error)
		// DeleteMultiByAdminIDRoleID 根据adminID删除多条数据
		DeleteMultiByAdminIDRoleID(ctx context.Context, adminID string, roleID string) error
		// DeleteMultiCacheByAdminIDRoleID 根据adminID删除多条数据，并删除缓存
		DeleteMultiCacheByAdminIDRoleID(ctx context.Context, adminID string, roleID string) error
		// DeleteMultiByAdminIDRoleIDTx 根据adminID删除多条数据(事务)
		DeleteMultiByAdminIDRoleIDTx(ctx context.Context, tx *gorm_gen_dao.Query, adminID string, roleID string) error
		// DeleteMultiCacheByAdminIDRoleIDTx 根据adminID删除多条数据，并删除缓存(事务)
		DeleteMultiCacheByAdminIDRoleIDTx(ctx context.Context, tx *gorm_gen_dao.Query, adminID string, roleID string) error
		// DeleteMultiByRoleIDAdminID 根据roleID删除多条数据
		DeleteMultiByRoleIDAdminID(ctx context.Context, roleID string, adminID string) error
		// DeleteMultiCacheByRoleIDAdminID 根据roleID删除多条数据，并删除缓存
		DeleteMultiCacheByRoleIDAdminID(ctx context.Context, roleID string, adminID string) error
		// DeleteMultiByRoleIDAdminIDTx 根据roleID删除多条数据(事务)
		DeleteMultiByRoleIDAdminIDTx(ctx context.Context, tx *gorm_gen_dao.Query, roleID string, adminID string) error
		// DeleteMultiCacheByRoleIDAdminIDTx 根据roleID删除多条数据，并删除缓存(事务)
		DeleteMultiCacheByRoleIDAdminIDTx(ctx context.Context, tx *gorm_gen_dao.Query, roleID string, adminID string) error
		// DeleteMultiByAdminID 根据AdminID删除多条数据
		DeleteMultiByAdminID(ctx context.Context, adminID string) error
		// DeleteMultiCacheByAdminID 根据adminID删除多条数据，并删除缓存
		DeleteMultiCacheByAdminID(ctx context.Context, adminID string) error
		// DeleteMultiByAdminIDTx 根据adminID删除多条数据
		DeleteMultiByAdminIDTx(ctx context.Context, tx *gorm_gen_dao.Query, adminID string) error
		// DeleteMultiCacheByAdminIDTx 根据adminID删除多条数据，并删除缓存
		DeleteMultiCacheByAdminIDTx(ctx context.Context, tx *gorm_gen_dao.Query, adminID string) error
		// DeleteMultiByAdminIDS 根据AdminIDS删除多条数据
		DeleteMultiByAdminIDS(ctx context.Context, adminIDS []string) error
		// DeleteMultiCacheByAdminIDS 根据AdminIDS删除多条数据，并删除缓存
		DeleteMultiCacheByAdminIDS(ctx context.Context, adminIDS []string) error
		// DeleteMultiByAdminIDSTx 根据AdminIDS删除多条数据(事务)
		DeleteMultiByAdminIDSTx(ctx context.Context, tx *gorm_gen_dao.Query, adminIDS []string) error
		// DeleteMultiCacheByAdminIDSTx 根据AdminIDS删除多条数据，并删除缓存(事务)
		DeleteMultiCacheByAdminIDSTx(ctx context.Context, tx *gorm_gen_dao.Query, adminIDS []string) error
		// DeleteMultiByRoleID 根据RoleID删除多条数据
		DeleteMultiByRoleID(ctx context.Context, roleID string) error
		// DeleteMultiCacheByRoleID 根据roleID删除多条数据，并删除缓存
		DeleteMultiCacheByRoleID(ctx context.Context, roleID string) error
		// DeleteMultiByRoleIDTx 根据roleID删除多条数据
		DeleteMultiByRoleIDTx(ctx context.Context, tx *gorm_gen_dao.Query, roleID string) error
		// DeleteMultiCacheByRoleIDTx 根据roleID删除多条数据，并删除缓存
		DeleteMultiCacheByRoleIDTx(ctx context.Context, tx *gorm_gen_dao.Query, roleID string) error
		// DeleteMultiByRoleIDS 根据RoleIDS删除多条数据
		DeleteMultiByRoleIDS(ctx context.Context, roleIDS []string) error
		// DeleteMultiCacheByRoleIDS 根据RoleIDS删除多条数据，并删除缓存
		DeleteMultiCacheByRoleIDS(ctx context.Context, roleIDS []string) error
		// DeleteMultiByRoleIDSTx 根据RoleIDS删除多条数据(事务)
		DeleteMultiByRoleIDSTx(ctx context.Context, tx *gorm_gen_dao.Query, roleIDS []string) error
		// DeleteMultiCacheByRoleIDSTx 根据RoleIDS删除多条数据，并删除缓存(事务)
		DeleteMultiCacheByRoleIDSTx(ctx context.Context, tx *gorm_gen_dao.Query, roleIDS []string) error
		// DeleteIndexCache 删除索引存在的缓存
		DeleteIndexCache(ctx context.Context, data ...*gorm_gen_model.AdminToRoleDemo) error
	}
	AdminToRoleDemoRepo struct {
		db       *gorm.DB
		cache    dbcache.IDBCache
		encoding encoding.API
	}
)

func NewAdminToRoleDemoRepo(cfg *config.Repo) *AdminToRoleDemoRepo {
	return &AdminToRoleDemoRepo{
		db:       cfg.DB,
		cache:    cfg.Cache,
		encoding: cfg.Encoding,
	}
}

// DeepCopy 深拷贝
func (a *AdminToRoleDemoRepo) DeepCopy(data *gorm_gen_model.AdminToRoleDemo) *gorm_gen_model.AdminToRoleDemo {
	newData := new(gorm_gen_model.AdminToRoleDemo)
	*newData = *data
	return newData
}

// CreateOne 创建一条数据
func (a *AdminToRoleDemoRepo) CreateOne(ctx context.Context, data *gorm_gen_model.AdminToRoleDemo) error {
	dao := gorm_gen_dao.Use(a.db).AdminToRoleDemo
	err := dao.WithContext(ctx).Create(data)
	if err != nil {
		return err
	}
	return nil
}

// CreateOneCache 创建一条数据, 并删除缓存
func (a *AdminToRoleDemoRepo) CreateOneCache(ctx context.Context, data *gorm_gen_model.AdminToRoleDemo) error {
	dao := gorm_gen_dao.Use(a.db).AdminToRoleDemo
	err := dao.WithContext(ctx).Create(data)
	if err != nil {
		return err
	}
	err = a.DeleteIndexCache(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

// CreateOneByTx 创建一条数据(事务)
func (a *AdminToRoleDemoRepo) CreateOneByTx(ctx context.Context, tx *gorm_gen_dao.Query, data *gorm_gen_model.AdminToRoleDemo) error {
	dao := tx.AdminToRoleDemo
	err := dao.WithContext(ctx).Create(data)
	if err != nil {
		return err
	}
	return nil
}

// CreateOneCacheByTx 创建一条数据(事务), 并删除缓存
func (a *AdminToRoleDemoRepo) CreateOneCacheByTx(ctx context.Context, tx *gorm_gen_dao.Query, data *gorm_gen_model.AdminToRoleDemo) error {
	dao := tx.AdminToRoleDemo
	err := dao.WithContext(ctx).Create(data)
	if err != nil {
		return err
	}
	err = a.DeleteIndexCache(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

// CreateBatch 批量创建数据
func (a *AdminToRoleDemoRepo) CreateBatch(ctx context.Context, data []*gorm_gen_model.AdminToRoleDemo, batchSize int) error {
	dao := gorm_gen_dao.Use(a.db).AdminToRoleDemo
	err := dao.WithContext(ctx).CreateInBatches(data, batchSize)
	if err != nil {
		return err
	}
	return nil
}

// CreateBatchCache 批量创建数据, 并删除缓存
func (a *AdminToRoleDemoRepo) CreateBatchCache(ctx context.Context, data []*gorm_gen_model.AdminToRoleDemo, batchSize int) error {
	dao := gorm_gen_dao.Use(a.db).AdminToRoleDemo
	err := dao.WithContext(ctx).CreateInBatches(data, batchSize)
	if err != nil {
		return err
	}
	err = a.DeleteIndexCache(ctx, data...)
	if err != nil {
		return err
	}
	return nil
}

// CreateBatchByTx 批量创建数据(事务)
func (a *AdminToRoleDemoRepo) CreateBatchByTx(ctx context.Context, tx *gorm_gen_dao.Query, data []*gorm_gen_model.AdminToRoleDemo, batchSize int) error {
	dao := tx.AdminToRoleDemo
	err := dao.WithContext(ctx).CreateInBatches(data, batchSize)
	if err != nil {
		return err
	}
	return nil
}

// CreateBatchCacheByTx 批量创建数据(事务), 并删除缓存
func (a *AdminToRoleDemoRepo) CreateBatchCacheByTx(ctx context.Context, tx *gorm_gen_dao.Query, data []*gorm_gen_model.AdminToRoleDemo, batchSize int) error {
	dao := tx.AdminToRoleDemo
	err := dao.WithContext(ctx).CreateInBatches(data, batchSize)
	if err != nil {
		return err
	}
	err = a.DeleteIndexCache(ctx, data...)
	if err != nil {
		return err
	}
	return nil
}

// UpsertOneByFields 根据fields字段Upsert一条数据
func (a *AdminToRoleDemoRepo) UpsertOneByFields(ctx context.Context, data *gorm_gen_model.AdminToRoleDemo, fields []string) error {
	if len(fields) == 0 {
		return errors.New("UpsertOneByFields fields is empty")
	}
	columns := make([]clause.Column, 0)
	for _, v := range fields {
		columns = append(columns, clause.Column{Name: v})
	}
	dao := gorm_gen_dao.Use(a.db).AdminToRoleDemo
	err := dao.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   columns,
		UpdateAll: true,
	}).Create(data)
	if err != nil {
		return err
	}
	return nil
}

// UpsertOneCacheByFields 根据fields字段Upsert一条数据, 并删除缓存
func (a *AdminToRoleDemoRepo) UpsertOneCacheByFields(ctx context.Context, data *gorm_gen_model.AdminToRoleDemo, fields []string) error {
	if len(fields) == 0 {
		return errors.New("UpsertOneByFields fields is empty")
	}
	fieldNameToValue := make(map[string]interface{})
	typ := reflect.TypeOf(data).Elem()
	val := reflect.ValueOf(data).Elem()
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		gormTag := field.Tag.Get("gorm")
		if gormTag != "" {
			gormTags := strings.Split(gormTag, ";")
			for _, v := range gormTags {
				if strings.Contains(v, "column") {
					columnName := strings.TrimPrefix(v, "column:")
					fieldValue := val.Field(i).Interface()
					fieldNameToValue[columnName] = fieldValue
					break
				}
			}
		}
	}
	whereExpressions := make([]clause.Expression, 0)
	columns := make([]clause.Column, 0)
	for _, v := range fields {
		whereExpressions = append(whereExpressions, clause.And(clause.Eq{Column: v, Value: fieldNameToValue[v]}))
		columns = append(columns, clause.Column{Name: v})
	}
	oldData := &gorm_gen_model.AdminToRoleDemo{}
	err := a.db.Model(&gorm_gen_model.AdminToRoleDemo{}).Clauses(whereExpressions...).First(oldData).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	dao := gorm_gen_dao.Use(a.db).AdminToRoleDemo
	err = dao.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   columns,
		UpdateAll: true,
	}).Create(data)
	if err != nil {
		return err
	}
	err = a.DeleteIndexCache(ctx, oldData, data)
	if err != nil {
		return err
	}
	return nil
}

// UpsertOneByFieldsTx 根据fields字段Upsert一条数据(事务)
func (a *AdminToRoleDemoRepo) UpsertOneByFieldsTx(ctx context.Context, tx *gorm_gen_dao.Query, data *gorm_gen_model.AdminToRoleDemo, fields []string) error {
	if len(fields) == 0 {
		return errors.New("UpsertOneByFieldsTx fields is empty")
	}
	columns := make([]clause.Column, 0)
	for _, v := range fields {
		columns = append(columns, clause.Column{Name: v})
	}
	dao := tx.AdminToRoleDemo
	err := dao.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   columns,
		UpdateAll: true,
	}).Create(data)
	if err != nil {
		return err
	}
	return nil
}

// UpsertOneCacheByFieldsTx 根据fields字段Upsert一条数据(事务), 并删除缓存
func (a *AdminToRoleDemoRepo) UpsertOneCacheByFieldsTx(ctx context.Context, tx *gorm_gen_dao.Query, data *gorm_gen_model.AdminToRoleDemo, fields []string) error {
	if len(fields) == 0 {
		return errors.New("UpsertOneByFieldsTx fields is empty")
	}
	fieldNameToValue := make(map[string]interface{})
	typ := reflect.TypeOf(data).Elem()
	val := reflect.ValueOf(data).Elem()
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		gormTag := field.Tag.Get("gorm")
		if gormTag != "" {
			gormTags := strings.Split(gormTag, ";")
			for _, v := range gormTags {
				if strings.Contains(v, "column") {
					columnName := strings.TrimPrefix(v, "column:")
					fieldValue := val.Field(i).Interface()
					fieldNameToValue[columnName] = fieldValue
					break
				}
			}
		}
	}
	whereExpressions := make([]clause.Expression, 0)
	columns := make([]clause.Column, 0)
	for _, v := range fields {
		whereExpressions = append(whereExpressions, clause.And(clause.Eq{Column: v, Value: fieldNameToValue[v]}))
		columns = append(columns, clause.Column{Name: v})
	}
	oldData := &gorm_gen_model.AdminToRoleDemo{}
	err := a.db.Model(&gorm_gen_model.AdminToRoleDemo{}).Clauses(whereExpressions...).First(oldData).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	dao := tx.AdminToRoleDemo
	err = dao.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   columns,
		UpdateAll: true,
	}).Create(data)
	if err != nil {
		return err
	}
	err = a.DeleteIndexCache(ctx, oldData, data)
	if err != nil {
		return err
	}
	return nil
}

// FindMultiByAdminIDRoleID 根据AdminIDRoleID查询多条数据
func (a *AdminToRoleDemoRepo) FindMultiByAdminIDRoleID(ctx context.Context, adminID string, roleID string) ([]*gorm_gen_model.AdminToRoleDemo, error) {
	dao := gorm_gen_dao.Use(a.db).AdminToRoleDemo
	result, err := dao.WithContext(ctx).Where(dao.AdminID.Eq(adminID), dao.RoleID.Eq(roleID)).Find()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FindMultiCacheByAdminIDRoleID 根据AdminIDRoleID查询多条数据，并设置缓存
func (a *AdminToRoleDemoRepo) FindMultiCacheByAdminIDRoleID(ctx context.Context, adminID string, roleID string) ([]*gorm_gen_model.AdminToRoleDemo, error) {
	resp := make([]*gorm_gen_model.AdminToRoleDemo, 0)
	cacheKey := a.cache.Key(CacheAdminToRoleDemoByAdminIDRoleIDPrefix, adminID, roleID)
	cacheValue, err := a.cache.Fetch(ctx, cacheKey, func() (string, error) {
		dao := gorm_gen_dao.Use(a.db).AdminToRoleDemo
		result, err := dao.WithContext(ctx).Where(dao.AdminID.Eq(adminID), dao.RoleID.Eq(roleID)).Find()
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return "", err
		}
		marshal, err := a.encoding.Marshal(result)
		if err != nil {
			return "", err
		}
		return string(marshal), nil
	})
	if err != nil {
		return nil, err
	}
	if cacheValue != "" {
		err = a.encoding.Unmarshal([]byte(cacheValue), &resp)
		if err != nil {
			return nil, err
		}
	}
	return resp, nil
}

// FindMultiByRoleIDAdminID 根据RoleIDAdminID查询多条数据
func (a *AdminToRoleDemoRepo) FindMultiByRoleIDAdminID(ctx context.Context, roleID string, adminID string) ([]*gorm_gen_model.AdminToRoleDemo, error) {
	dao := gorm_gen_dao.Use(a.db).AdminToRoleDemo
	result, err := dao.WithContext(ctx).Where(dao.RoleID.Eq(roleID), dao.AdminID.Eq(adminID)).Find()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FindMultiCacheByRoleIDAdminID 根据RoleIDAdminID查询多条数据，并设置缓存
func (a *AdminToRoleDemoRepo) FindMultiCacheByRoleIDAdminID(ctx context.Context, roleID string, adminID string) ([]*gorm_gen_model.AdminToRoleDemo, error) {
	resp := make([]*gorm_gen_model.AdminToRoleDemo, 0)
	cacheKey := a.cache.Key(CacheAdminToRoleDemoByRoleIDAdminIDPrefix, roleID, adminID)
	cacheValue, err := a.cache.Fetch(ctx, cacheKey, func() (string, error) {
		dao := gorm_gen_dao.Use(a.db).AdminToRoleDemo
		result, err := dao.WithContext(ctx).Where(dao.RoleID.Eq(roleID), dao.AdminID.Eq(adminID)).Find()
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return "", err
		}
		marshal, err := a.encoding.Marshal(result)
		if err != nil {
			return "", err
		}
		return string(marshal), nil
	})
	if err != nil {
		return nil, err
	}
	if cacheValue != "" {
		err = a.encoding.Unmarshal([]byte(cacheValue), &resp)
		if err != nil {
			return nil, err
		}
	}
	return resp, nil
}

// FindMultiByAdminID 根据adminID查询多条数据
func (a *AdminToRoleDemoRepo) FindMultiByAdminID(ctx context.Context, adminID string) ([]*gorm_gen_model.AdminToRoleDemo, error) {
	dao := gorm_gen_dao.Use(a.db).AdminToRoleDemo
	result, err := dao.WithContext(ctx).Where(dao.AdminID.Eq(adminID)).Find()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FindMultiCacheByAdminID 根据adminID查询多条数据，并设置缓存
func (a *AdminToRoleDemoRepo) FindMultiCacheByAdminID(ctx context.Context, adminID string) ([]*gorm_gen_model.AdminToRoleDemo, error) {
	resp := make([]*gorm_gen_model.AdminToRoleDemo, 0)
	cacheKey := a.cache.Key(CacheAdminToRoleDemoByAdminIDPrefix, adminID)
	cacheValue, err := a.cache.Fetch(ctx, cacheKey, func() (string, error) {
		dao := gorm_gen_dao.Use(a.db).AdminToRoleDemo
		result, err := dao.WithContext(ctx).Where(dao.AdminID.Eq(adminID)).Find()
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return "", err
		}
		marshal, err := a.encoding.Marshal(result)
		if err != nil {
			return "", err
		}
		return string(marshal), nil
	})
	if err != nil {
		return nil, err
	}
	if cacheValue != "" {
		err = a.encoding.Unmarshal([]byte(cacheValue), &resp)
		if err != nil {
			return nil, err
		}
	}
	return resp, nil
}

// FindMultiByAdminIDS 根据adminIDS查询多条数据
func (a *AdminToRoleDemoRepo) FindMultiByAdminIDS(ctx context.Context, adminIDS []string) ([]*gorm_gen_model.AdminToRoleDemo, error) {
	dao := gorm_gen_dao.Use(a.db).AdminToRoleDemo
	result, err := dao.WithContext(ctx).Where(dao.AdminID.In(adminIDS...)).Find()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FindMultiCacheByAdminIDS 根据adminIDS查询多条数据，并设置缓存
func (a *AdminToRoleDemoRepo) FindMultiCacheByAdminIDS(ctx context.Context, adminIDS []string) ([]*gorm_gen_model.AdminToRoleDemo, error) {
	resp := make([]*gorm_gen_model.AdminToRoleDemo, 0)
	cacheKeys := make([]string, 0)
	keyToParam := make(map[string]string)
	for _, v := range adminIDS {
		cacheKey := a.cache.Key(CacheAdminToRoleDemoByAdminIDPrefix, v)
		cacheKeys = append(cacheKeys, cacheKey)
		keyToParam[cacheKey] = v
	}
	cacheValue, err := a.cache.FetchBatch(ctx, cacheKeys, func(miss []string) (map[string]string, error) {
		dbValue := make(map[string]string)
		params := make([]string, 0)
		for _, v := range miss {
			dbValue[v] = ""
			params = append(params, keyToParam[v])
		}
		dao := gorm_gen_dao.Use(a.db).AdminToRoleDemo
		result, err := dao.WithContext(ctx).Where(dao.AdminID.In(params...)).Find()
		if err != nil {
			return nil, err
		}
		keyToValues := make(map[string][]*gorm_gen_model.AdminToRoleDemo)
		for _, v := range result {
			key := a.cache.Key(CacheAdminToRoleDemoByAdminIDPrefix, v.AdminID)
			if keyToValues[key] == nil {
				keyToValues[key] = make([]*gorm_gen_model.AdminToRoleDemo, 0)
			}
			keyToValues[key] = append(keyToValues[key], v)
		}
		for k := range dbValue {
			if keyToValues[k] != nil {
				marshal, err := a.encoding.Marshal(keyToValues[k])
				if err != nil {
					return nil, err
				}
				dbValue[k] = string(marshal)
			}
		}
		return dbValue, nil
	})
	if err != nil {
		return nil, err
	}
	for _, cacheKey := range cacheKeys {
		if cacheValue[cacheKey] != "" {
			tmp := make([]*gorm_gen_model.AdminToRoleDemo, 0)
			err := a.encoding.Unmarshal([]byte(cacheValue[cacheKey]), &tmp)
			if err != nil {
				return nil, err
			}
			resp = append(resp, tmp...)
		}
	}
	return resp, nil
}

// FindMultiByRoleID 根据roleID查询多条数据
func (a *AdminToRoleDemoRepo) FindMultiByRoleID(ctx context.Context, roleID string) ([]*gorm_gen_model.AdminToRoleDemo, error) {
	dao := gorm_gen_dao.Use(a.db).AdminToRoleDemo
	result, err := dao.WithContext(ctx).Where(dao.RoleID.Eq(roleID)).Find()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FindMultiCacheByRoleID 根据roleID查询多条数据，并设置缓存
func (a *AdminToRoleDemoRepo) FindMultiCacheByRoleID(ctx context.Context, roleID string) ([]*gorm_gen_model.AdminToRoleDemo, error) {
	resp := make([]*gorm_gen_model.AdminToRoleDemo, 0)
	cacheKey := a.cache.Key(CacheAdminToRoleDemoByRoleIDPrefix, roleID)
	cacheValue, err := a.cache.Fetch(ctx, cacheKey, func() (string, error) {
		dao := gorm_gen_dao.Use(a.db).AdminToRoleDemo
		result, err := dao.WithContext(ctx).Where(dao.RoleID.Eq(roleID)).Find()
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return "", err
		}
		marshal, err := a.encoding.Marshal(result)
		if err != nil {
			return "", err
		}
		return string(marshal), nil
	})
	if err != nil {
		return nil, err
	}
	if cacheValue != "" {
		err = a.encoding.Unmarshal([]byte(cacheValue), &resp)
		if err != nil {
			return nil, err
		}
	}
	return resp, nil
}

// FindMultiByRoleIDS 根据roleIDS查询多条数据
func (a *AdminToRoleDemoRepo) FindMultiByRoleIDS(ctx context.Context, roleIDS []string) ([]*gorm_gen_model.AdminToRoleDemo, error) {
	dao := gorm_gen_dao.Use(a.db).AdminToRoleDemo
	result, err := dao.WithContext(ctx).Where(dao.RoleID.In(roleIDS...)).Find()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FindMultiCacheByRoleIDS 根据roleIDS查询多条数据，并设置缓存
func (a *AdminToRoleDemoRepo) FindMultiCacheByRoleIDS(ctx context.Context, roleIDS []string) ([]*gorm_gen_model.AdminToRoleDemo, error) {
	resp := make([]*gorm_gen_model.AdminToRoleDemo, 0)
	cacheKeys := make([]string, 0)
	keyToParam := make(map[string]string)
	for _, v := range roleIDS {
		cacheKey := a.cache.Key(CacheAdminToRoleDemoByRoleIDPrefix, v)
		cacheKeys = append(cacheKeys, cacheKey)
		keyToParam[cacheKey] = v
	}
	cacheValue, err := a.cache.FetchBatch(ctx, cacheKeys, func(miss []string) (map[string]string, error) {
		dbValue := make(map[string]string)
		params := make([]string, 0)
		for _, v := range miss {
			dbValue[v] = ""
			params = append(params, keyToParam[v])
		}
		dao := gorm_gen_dao.Use(a.db).AdminToRoleDemo
		result, err := dao.WithContext(ctx).Where(dao.RoleID.In(params...)).Find()
		if err != nil {
			return nil, err
		}
		keyToValues := make(map[string][]*gorm_gen_model.AdminToRoleDemo)
		for _, v := range result {
			key := a.cache.Key(CacheAdminToRoleDemoByRoleIDPrefix, v.RoleID)
			if keyToValues[key] == nil {
				keyToValues[key] = make([]*gorm_gen_model.AdminToRoleDemo, 0)
			}
			keyToValues[key] = append(keyToValues[key], v)
		}
		for k := range dbValue {
			if keyToValues[k] != nil {
				marshal, err := a.encoding.Marshal(keyToValues[k])
				if err != nil {
					return nil, err
				}
				dbValue[k] = string(marshal)
			}
		}
		return dbValue, nil
	})
	if err != nil {
		return nil, err
	}
	for _, cacheKey := range cacheKeys {
		if cacheValue[cacheKey] != "" {
			tmp := make([]*gorm_gen_model.AdminToRoleDemo, 0)
			err := a.encoding.Unmarshal([]byte(cacheValue[cacheKey]), &tmp)
			if err != nil {
				return nil, err
			}
			resp = append(resp, tmp...)
		}
	}
	return resp, nil
}

// FindMultiByCondition 自定义查询数据(通用)
func (a *AdminToRoleDemoRepo) FindMultiByCondition(ctx context.Context, conditionReq *condition.Req) ([]*gorm_gen_model.AdminToRoleDemo, *condition.Reply, error) {
	result := make([]*gorm_gen_model.AdminToRoleDemo, 0)
	conditionReply := &condition.Reply{}
	var total int64
	whereExpressions, orderExpressions, err := conditionReq.ConvertToGormExpression(gorm_gen_model.AdminToRoleDemo{})
	if err != nil {
		return result, conditionReply, err
	}
	err = a.db.WithContext(ctx).Model(&gorm_gen_model.AdminToRoleDemo{}).Select([]string{"*"}).Clauses(whereExpressions...).Count(&total).Error
	if err != nil {
		return result, conditionReply, err
	}
	if total == 0 {
		return result, conditionReply, nil
	}
	conditionReply, err = conditionReq.ConvertToPage(int32(total))
	if err != nil {
		return result, conditionReply, err
	}
	query := a.db.WithContext(ctx).Model(&gorm_gen_model.AdminToRoleDemo{}).Clauses(whereExpressions...).Clauses(orderExpressions...)
	if conditionReply.Page != 0 && conditionReply.PageSize != 0 {
		query = query.Offset(int((conditionReply.Page - 1) * conditionReply.PageSize))
		query = query.Limit(int(conditionReply.PageSize))
	}
	err = query.Find(&result).Error
	if err != nil {
		return result, conditionReply, err
	}
	return result, conditionReply, err
}

// DeleteMultiByAdminIDRoleID 根据adminID删除多条数据
func (a *AdminToRoleDemoRepo) DeleteMultiByAdminIDRoleID(ctx context.Context, adminID string, roleID string) error {
	dao := gorm_gen_dao.Use(a.db).AdminToRoleDemo
	_, err := dao.WithContext(ctx).Where(dao.AdminID.Eq(adminID), dao.RoleID.Eq(roleID)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiCacheByAdminIDRoleID 根据adminID删除多条数据，并删除缓存
func (a *AdminToRoleDemoRepo) DeleteMultiCacheByAdminIDRoleID(ctx context.Context, adminID string, roleID string) error {
	dao := gorm_gen_dao.Use(a.db).AdminToRoleDemo
	result, err := dao.WithContext(ctx).Where(dao.AdminID.Eq(adminID), dao.RoleID.Eq(roleID)).Find()
	if err != nil {
		return err
	}
	if len(result) == 0 {
		return nil
	}
	_, err = dao.WithContext(ctx).Where(dao.AdminID.Eq(adminID), dao.RoleID.Eq(roleID)).Delete()
	if err != nil {
		return err
	}
	err = a.DeleteIndexCache(ctx, result...)
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiByAdminIDRoleIDTx 根据adminID删除多条数据
func (a *AdminToRoleDemoRepo) DeleteMultiByAdminIDRoleIDTx(ctx context.Context, tx *gorm_gen_dao.Query, adminID string, roleID string) error {
	dao := tx.AdminToRoleDemo
	_, err := dao.WithContext(ctx).Where(dao.AdminID.Eq(adminID), dao.RoleID.Eq(roleID)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiCacheByAdminIDRoleIDTx 根据adminID删除多条数据，并删除缓存
func (a *AdminToRoleDemoRepo) DeleteMultiCacheByAdminIDRoleIDTx(ctx context.Context, tx *gorm_gen_dao.Query, adminID string, roleID string) error {
	dao := tx.AdminToRoleDemo
	result, err := dao.WithContext(ctx).Where(dao.AdminID.Eq(adminID), dao.RoleID.Eq(roleID)).Find()
	if err != nil {
		return err
	}
	if len(result) == 0 {
		return nil
	}
	_, err = dao.WithContext(ctx).Where(dao.AdminID.Eq(adminID), dao.RoleID.Eq(roleID)).Delete()
	if err != nil {
		return err
	}
	err = a.DeleteIndexCache(ctx, result...)
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiByRoleIDAdminID 根据roleID删除多条数据
func (a *AdminToRoleDemoRepo) DeleteMultiByRoleIDAdminID(ctx context.Context, roleID string, adminID string) error {
	dao := gorm_gen_dao.Use(a.db).AdminToRoleDemo
	_, err := dao.WithContext(ctx).Where(dao.RoleID.Eq(roleID), dao.AdminID.Eq(adminID)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiCacheByRoleIDAdminID 根据roleID删除多条数据，并删除缓存
func (a *AdminToRoleDemoRepo) DeleteMultiCacheByRoleIDAdminID(ctx context.Context, roleID string, adminID string) error {
	dao := gorm_gen_dao.Use(a.db).AdminToRoleDemo
	result, err := dao.WithContext(ctx).Where(dao.RoleID.Eq(roleID), dao.AdminID.Eq(adminID)).Find()
	if err != nil {
		return err
	}
	if len(result) == 0 {
		return nil
	}
	_, err = dao.WithContext(ctx).Where(dao.RoleID.Eq(roleID), dao.AdminID.Eq(adminID)).Delete()
	if err != nil {
		return err
	}
	err = a.DeleteIndexCache(ctx, result...)
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiByRoleIDAdminIDTx 根据roleID删除多条数据
func (a *AdminToRoleDemoRepo) DeleteMultiByRoleIDAdminIDTx(ctx context.Context, tx *gorm_gen_dao.Query, roleID string, adminID string) error {
	dao := tx.AdminToRoleDemo
	_, err := dao.WithContext(ctx).Where(dao.RoleID.Eq(roleID), dao.AdminID.Eq(adminID)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiCacheByRoleIDAdminIDTx 根据roleID删除多条数据，并删除缓存
func (a *AdminToRoleDemoRepo) DeleteMultiCacheByRoleIDAdminIDTx(ctx context.Context, tx *gorm_gen_dao.Query, roleID string, adminID string) error {
	dao := tx.AdminToRoleDemo
	result, err := dao.WithContext(ctx).Where(dao.RoleID.Eq(roleID), dao.AdminID.Eq(adminID)).Find()
	if err != nil {
		return err
	}
	if len(result) == 0 {
		return nil
	}
	_, err = dao.WithContext(ctx).Where(dao.RoleID.Eq(roleID), dao.AdminID.Eq(adminID)).Delete()
	if err != nil {
		return err
	}
	err = a.DeleteIndexCache(ctx, result...)
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiByAdminID 根据AdminID删除多条数据
func (a *AdminToRoleDemoRepo) DeleteMultiByAdminID(ctx context.Context, adminID string) error {
	dao := gorm_gen_dao.Use(a.db).AdminToRoleDemo
	_, err := dao.WithContext(ctx).Where(dao.AdminID.Eq(adminID)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiCacheByAdminID 根据adminID删除多条数据，并删除缓存
func (a *AdminToRoleDemoRepo) DeleteMultiCacheByAdminID(ctx context.Context, adminID string) error {
	dao := gorm_gen_dao.Use(a.db).AdminToRoleDemo
	result, err := dao.WithContext(ctx).Where(dao.AdminID.Eq(adminID)).Find()
	if err != nil {
		return err
	}
	if len(result) == 0 {
		return nil
	}
	_, err = dao.WithContext(ctx).Where(dao.AdminID.Eq(adminID)).Delete()
	if err != nil {
		return err
	}
	err = a.DeleteIndexCache(ctx, result...)
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiByAdminIDTx 根据adminID删除多条数据
func (a *AdminToRoleDemoRepo) DeleteMultiByAdminIDTx(ctx context.Context, tx *gorm_gen_dao.Query, adminID string) error {
	dao := tx.AdminToRoleDemo
	_, err := dao.WithContext(ctx).Where(dao.AdminID.Eq(adminID)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiCacheByAdminIDTx 根据adminID删除多条数据，并删除缓存
func (a *AdminToRoleDemoRepo) DeleteMultiCacheByAdminIDTx(ctx context.Context, tx *gorm_gen_dao.Query, adminID string) error {
	dao := tx.AdminToRoleDemo
	result, err := dao.WithContext(ctx).Where(dao.AdminID.Eq(adminID)).Find()
	if err != nil {
		return err
	}
	if len(result) == 0 {
		return nil
	}
	_, err = dao.WithContext(ctx).Where(dao.AdminID.Eq(adminID)).Delete()
	if err != nil {
		return err
	}
	err = a.DeleteIndexCache(ctx, result...)
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiByAdminIDS 根据adminIDS删除多条数据
func (a *AdminToRoleDemoRepo) DeleteMultiByAdminIDS(ctx context.Context, adminIDS []string) error {
	dao := gorm_gen_dao.Use(a.db).AdminToRoleDemo
	_, err := dao.WithContext(ctx).Where(dao.AdminID.In(adminIDS...)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiCacheByAdminIDS 根据adminIDS删除多条数据，并删除缓存
func (a *AdminToRoleDemoRepo) DeleteMultiCacheByAdminIDS(ctx context.Context, adminIDS []string) error {
	dao := gorm_gen_dao.Use(a.db).AdminToRoleDemo
	result, err := dao.WithContext(ctx).Where(dao.AdminID.In(adminIDS...)).Find()
	if err != nil {
		return err
	}
	if len(result) == 0 {
		return nil
	}
	_, err = dao.WithContext(ctx).Where(dao.AdminID.In(adminIDS...)).Delete()
	if err != nil {
		return err
	}
	err = a.DeleteIndexCache(ctx, result...)
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiByAdminIDSTx 根据adminIDS删除多条数据
func (a *AdminToRoleDemoRepo) DeleteMultiByAdminIDSTx(ctx context.Context, tx *gorm_gen_dao.Query, adminIDS []string) error {
	dao := tx.AdminToRoleDemo
	_, err := dao.WithContext(ctx).Where(dao.AdminID.In(adminIDS...)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiCacheByAdminIDSTx 根据adminIDS删除多条数据，并删除缓存
func (a *AdminToRoleDemoRepo) DeleteMultiCacheByAdminIDSTx(ctx context.Context, tx *gorm_gen_dao.Query, adminIDS []string) error {
	dao := tx.AdminToRoleDemo
	result, err := dao.WithContext(ctx).Where(dao.AdminID.In(adminIDS...)).Find()
	if err != nil {
		return err
	}
	if len(result) == 0 {
		return nil
	}
	_, err = dao.WithContext(ctx).Where(dao.AdminID.In(adminIDS...)).Delete()
	if err != nil {
		return err
	}
	err = a.DeleteIndexCache(ctx, result...)
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiByRoleID 根据RoleID删除多条数据
func (a *AdminToRoleDemoRepo) DeleteMultiByRoleID(ctx context.Context, roleID string) error {
	dao := gorm_gen_dao.Use(a.db).AdminToRoleDemo
	_, err := dao.WithContext(ctx).Where(dao.RoleID.Eq(roleID)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiCacheByRoleID 根据roleID删除多条数据，并删除缓存
func (a *AdminToRoleDemoRepo) DeleteMultiCacheByRoleID(ctx context.Context, roleID string) error {
	dao := gorm_gen_dao.Use(a.db).AdminToRoleDemo
	result, err := dao.WithContext(ctx).Where(dao.RoleID.Eq(roleID)).Find()
	if err != nil {
		return err
	}
	if len(result) == 0 {
		return nil
	}
	_, err = dao.WithContext(ctx).Where(dao.RoleID.Eq(roleID)).Delete()
	if err != nil {
		return err
	}
	err = a.DeleteIndexCache(ctx, result...)
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiByRoleIDTx 根据roleID删除多条数据
func (a *AdminToRoleDemoRepo) DeleteMultiByRoleIDTx(ctx context.Context, tx *gorm_gen_dao.Query, roleID string) error {
	dao := tx.AdminToRoleDemo
	_, err := dao.WithContext(ctx).Where(dao.RoleID.Eq(roleID)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiCacheByRoleIDTx 根据roleID删除多条数据，并删除缓存
func (a *AdminToRoleDemoRepo) DeleteMultiCacheByRoleIDTx(ctx context.Context, tx *gorm_gen_dao.Query, roleID string) error {
	dao := tx.AdminToRoleDemo
	result, err := dao.WithContext(ctx).Where(dao.RoleID.Eq(roleID)).Find()
	if err != nil {
		return err
	}
	if len(result) == 0 {
		return nil
	}
	_, err = dao.WithContext(ctx).Where(dao.RoleID.Eq(roleID)).Delete()
	if err != nil {
		return err
	}
	err = a.DeleteIndexCache(ctx, result...)
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiByRoleIDS 根据roleIDS删除多条数据
func (a *AdminToRoleDemoRepo) DeleteMultiByRoleIDS(ctx context.Context, roleIDS []string) error {
	dao := gorm_gen_dao.Use(a.db).AdminToRoleDemo
	_, err := dao.WithContext(ctx).Where(dao.RoleID.In(roleIDS...)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiCacheByRoleIDS 根据roleIDS删除多条数据，并删除缓存
func (a *AdminToRoleDemoRepo) DeleteMultiCacheByRoleIDS(ctx context.Context, roleIDS []string) error {
	dao := gorm_gen_dao.Use(a.db).AdminToRoleDemo
	result, err := dao.WithContext(ctx).Where(dao.RoleID.In(roleIDS...)).Find()
	if err != nil {
		return err
	}
	if len(result) == 0 {
		return nil
	}
	_, err = dao.WithContext(ctx).Where(dao.RoleID.In(roleIDS...)).Delete()
	if err != nil {
		return err
	}
	err = a.DeleteIndexCache(ctx, result...)
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiByRoleIDSTx 根据roleIDS删除多条数据
func (a *AdminToRoleDemoRepo) DeleteMultiByRoleIDSTx(ctx context.Context, tx *gorm_gen_dao.Query, roleIDS []string) error {
	dao := tx.AdminToRoleDemo
	_, err := dao.WithContext(ctx).Where(dao.RoleID.In(roleIDS...)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiCacheByRoleIDSTx 根据roleIDS删除多条数据，并删除缓存
func (a *AdminToRoleDemoRepo) DeleteMultiCacheByRoleIDSTx(ctx context.Context, tx *gorm_gen_dao.Query, roleIDS []string) error {
	dao := tx.AdminToRoleDemo
	result, err := dao.WithContext(ctx).Where(dao.RoleID.In(roleIDS...)).Find()
	if err != nil {
		return err
	}
	if len(result) == 0 {
		return nil
	}
	_, err = dao.WithContext(ctx).Where(dao.RoleID.In(roleIDS...)).Delete()
	if err != nil {
		return err
	}
	err = a.DeleteIndexCache(ctx, result...)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUniqueIndexCache 删除索引存在的缓存
func (a *AdminToRoleDemoRepo) DeleteIndexCache(ctx context.Context, data ...*gorm_gen_model.AdminToRoleDemo) error {
	KeyMap := make(map[string]struct{})
	for _, v := range data {
		if v != nil {
			KeyMap[a.cache.Key(CacheAdminToRoleDemoByAdminIDRoleIDPrefix, v.AdminID, v.RoleID)] = struct{}{}
			KeyMap[a.cache.Key(CacheAdminToRoleDemoByRoleIDAdminIDPrefix, v.RoleID, v.AdminID)] = struct{}{}
			KeyMap[a.cache.Key(CacheAdminToRoleDemoByAdminIDPrefix, v.AdminID)] = struct{}{}
			KeyMap[a.cache.Key(CacheAdminToRoleDemoByRoleIDPrefix, v.RoleID)] = struct{}{}

		}
	}
	keys := make([]string, 0, len(KeyMap))
	for k := range KeyMap {
		keys = append(keys, k)
	}
	err := a.cache.DelBatch(ctx, keys)
	if err != nil {
		return err
	}
	return nil
}

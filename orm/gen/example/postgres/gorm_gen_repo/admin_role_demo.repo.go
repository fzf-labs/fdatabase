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

var _ IAdminRoleDemoRepo = (*AdminRoleDemoRepo)(nil)

var (
	CacheAdminRoleDemoByIDPrefix = "DBCache:gorm_gen:AdminRoleDemoByID"
)

type (
	IAdminRoleDemoRepo interface {
		// DeepCopy 深拷贝
		DeepCopy(data *gorm_gen_model.AdminRoleDemo) *gorm_gen_model.AdminRoleDemo
		// CreateOne 创建一条数据
		CreateOne(ctx context.Context, data *gorm_gen_model.AdminRoleDemo) error
		// CreateOneCache 创建一条数据, 并删除缓存
		CreateOneCache(ctx context.Context, data *gorm_gen_model.AdminRoleDemo) error
		// CreateOneByTx 创建一条数据(事务)
		CreateOneByTx(ctx context.Context, tx *gorm_gen_dao.Query, data *gorm_gen_model.AdminRoleDemo) error
		// CreateOneCacheByTx 创建一条数据(事务), 并删除缓存
		CreateOneCacheByTx(ctx context.Context, tx *gorm_gen_dao.Query, data *gorm_gen_model.AdminRoleDemo) error
		// CreateBatch 批量创建数据
		CreateBatch(ctx context.Context, data []*gorm_gen_model.AdminRoleDemo, batchSize int) error
		// CreateBatchCache 批量创建数据, 并删除缓存
		CreateBatchCache(ctx context.Context, data []*gorm_gen_model.AdminRoleDemo, batchSize int) error
		// CreateBatchByTx 批量创建数据(事务)
		CreateBatchByTx(ctx context.Context, tx *gorm_gen_dao.Query, data []*gorm_gen_model.AdminRoleDemo, batchSize int) error
		// CreateBatchCacheByTx 批量创建数据(事务), 并删除缓存
		CreateBatchCacheByTx(ctx context.Context, tx *gorm_gen_dao.Query, data []*gorm_gen_model.AdminRoleDemo, batchSize int) error
		// UpsertOne Upsert一条数据
		UpsertOne(ctx context.Context, data *gorm_gen_model.AdminRoleDemo) error
		// UpsertOneCache Upsert一条数据, 并删除缓存
		UpsertOneCache(ctx context.Context, data *gorm_gen_model.AdminRoleDemo) error
		// UpsertOneByTx Upsert一条数据(事务)
		UpsertOneByTx(ctx context.Context, tx *gorm_gen_dao.Query, data *gorm_gen_model.AdminRoleDemo) error
		// UpsertOneCacheByTx Upsert一条数据(事务), 并删除缓存
		UpsertOneCacheByTx(ctx context.Context, tx *gorm_gen_dao.Query, data *gorm_gen_model.AdminRoleDemo) error
		// UpsertOneByFields 根据fields字段Upsert一条数据
		UpsertOneByFields(ctx context.Context, data *gorm_gen_model.AdminRoleDemo, fields []string) error
		// UpsertOneCacheByFields 根据fields字段Upsert一条数据, 并删除缓存
		UpsertOneCacheByFields(ctx context.Context, data *gorm_gen_model.AdminRoleDemo, fields []string) error
		// UpsertOneByFieldsTx 根据fields字段Upsert一条数据(事务)
		UpsertOneByFieldsTx(ctx context.Context, tx *gorm_gen_dao.Query, data *gorm_gen_model.AdminRoleDemo, fields []string) error
		// UpsertOneCacheByFieldsTx 根据fields字段Upsert一条数据(事务), 并删除缓存
		UpsertOneCacheByFieldsTx(ctx context.Context, tx *gorm_gen_dao.Query, data *gorm_gen_model.AdminRoleDemo, fields []string) error
		// UpdateOne 更新一条数据
		UpdateOne(ctx context.Context, newData *gorm_gen_model.AdminRoleDemo) error
		// UpdateOneCache 更新一条数据，并删除缓存
		UpdateOneCache(ctx context.Context, newData *gorm_gen_model.AdminRoleDemo, oldData *gorm_gen_model.AdminRoleDemo) error
		// UpdateOneByTx 更新一条数据(事务)
		UpdateOneByTx(ctx context.Context, tx *gorm_gen_dao.Query, newData *gorm_gen_model.AdminRoleDemo) error
		// UpdateOneCacheByTx 更新一条数据(事务)，并删除缓存
		UpdateOneCacheByTx(ctx context.Context, tx *gorm_gen_dao.Query, newData *gorm_gen_model.AdminRoleDemo, oldData *gorm_gen_model.AdminRoleDemo) error
		// UpdateOneCacheWithZero 更新一条数据,包含零值，并删除缓存
		UpdateOneWithZero(ctx context.Context, newData *gorm_gen_model.AdminRoleDemo) error
		// UpdateOneCacheWithZero 更新一条数据,包含零值，并删除缓存
		UpdateOneCacheWithZero(ctx context.Context, newData *gorm_gen_model.AdminRoleDemo, oldData *gorm_gen_model.AdminRoleDemo) error
		// UpdateOneCacheWithZeroByTx 更新一条数据(事务),包含零值，并删除缓存
		UpdateOneWithZeroByTx(ctx context.Context, tx *gorm_gen_dao.Query, newData *gorm_gen_model.AdminRoleDemo) error
		// UpdateOneCacheWithZeroByTx 更新一条数据(事务),包含零值，并删除缓存
		UpdateOneCacheWithZeroByTx(ctx context.Context, tx *gorm_gen_dao.Query, newData *gorm_gen_model.AdminRoleDemo, oldData *gorm_gen_model.AdminRoleDemo) error
		// UpdateBatchByIDS 根据主键IDS批量更新
		UpdateBatchByIDS(ctx context.Context, IDS []string, data map[string]interface{}) error
		// UpdateBatchByIDSTx 根据主键IDS批量更新(事务)
		UpdateBatchByIDSTx(ctx context.Context, tx *gorm_gen_dao.Query, IDS []string, data map[string]interface{}) error
		// FindOneByID 根据ID查询一条数据
		FindOneByID(ctx context.Context, ID string) (*gorm_gen_model.AdminRoleDemo, error)
		// FindOneCacheByID 根据ID查询一条数据，并设置缓存
		FindOneCacheByID(ctx context.Context, ID string) (*gorm_gen_model.AdminRoleDemo, error)
		// FindMultiByIDS 根据IDS查询多条数据
		FindMultiByIDS(ctx context.Context, IDS []string) ([]*gorm_gen_model.AdminRoleDemo, error)
		// FindMultiCacheByIDS 根据IDS查询多条数据，并设置缓存
		FindMultiCacheByIDS(ctx context.Context, IDS []string) ([]*gorm_gen_model.AdminRoleDemo, error)
		// FindMultiByCondition 根据自定义条件查询数据
		FindMultiByCondition(ctx context.Context, conditionReq *condition.Req) ([]*gorm_gen_model.AdminRoleDemo, *condition.Reply, error)
		// DeleteOneByID 根据ID删除一条数据
		DeleteOneByID(ctx context.Context, ID string) error
		// DeleteOneCacheByID 根据ID删除一条数据，并删除缓存
		DeleteOneCacheByID(ctx context.Context, ID string) error
		// DeleteOneByIDTx 根据ID删除一条数据(事务)
		DeleteOneByIDTx(ctx context.Context, tx *gorm_gen_dao.Query, ID string) error
		// DeleteOneCacheByIDTx 根据ID删除一条数据，并删除缓存(事务)
		DeleteOneCacheByIDTx(ctx context.Context, tx *gorm_gen_dao.Query, ID string) error
		// DeleteMultiByIDS 根据IDS删除多条数据
		DeleteMultiByIDS(ctx context.Context, IDS []string) error
		// DeleteMultiCacheByIDS 根据IDS删除多条数据，并删除缓存
		DeleteMultiCacheByIDS(ctx context.Context, IDS []string) error
		// DeleteMultiByIDSTx 根据IDS删除多条数据(事务)
		DeleteMultiByIDSTx(ctx context.Context, tx *gorm_gen_dao.Query, IDS []string) error
		// DeleteMultiCacheByIDSTx 根据IDS删除多条数据，并删除缓存(事务)
		DeleteMultiCacheByIDSTx(ctx context.Context, tx *gorm_gen_dao.Query, IDS []string) error
		// DeleteIndexCache 删除索引存在的缓存
		DeleteIndexCache(ctx context.Context, data ...*gorm_gen_model.AdminRoleDemo) error
	}
	AdminRoleDemoRepo struct {
		db       *gorm.DB
		cache    dbcache.IDBCache
		encoding encoding.API
	}
)

func NewAdminRoleDemoRepo(cfg *config.Repo) *AdminRoleDemoRepo {
	return &AdminRoleDemoRepo{
		db:       cfg.DB,
		cache:    cfg.Cache,
		encoding: cfg.Encoding,
	}
}

// DeepCopy 深拷贝
func (a *AdminRoleDemoRepo) DeepCopy(data *gorm_gen_model.AdminRoleDemo) *gorm_gen_model.AdminRoleDemo {
	newData := new(gorm_gen_model.AdminRoleDemo)
	*newData = *data
	return newData
}

// CreateOne 创建一条数据
func (a *AdminRoleDemoRepo) CreateOne(ctx context.Context, data *gorm_gen_model.AdminRoleDemo) error {
	dao := gorm_gen_dao.Use(a.db).AdminRoleDemo
	err := dao.WithContext(ctx).Create(data)
	if err != nil {
		return err
	}
	return nil
}

// CreateOneCache 创建一条数据, 并删除缓存
func (a *AdminRoleDemoRepo) CreateOneCache(ctx context.Context, data *gorm_gen_model.AdminRoleDemo) error {
	dao := gorm_gen_dao.Use(a.db).AdminRoleDemo
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
func (a *AdminRoleDemoRepo) CreateOneByTx(ctx context.Context, tx *gorm_gen_dao.Query, data *gorm_gen_model.AdminRoleDemo) error {
	dao := tx.AdminRoleDemo
	err := dao.WithContext(ctx).Create(data)
	if err != nil {
		return err
	}
	return nil
}

// CreateOneCacheByTx 创建一条数据(事务), 并删除缓存
func (a *AdminRoleDemoRepo) CreateOneCacheByTx(ctx context.Context, tx *gorm_gen_dao.Query, data *gorm_gen_model.AdminRoleDemo) error {
	dao := tx.AdminRoleDemo
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
func (a *AdminRoleDemoRepo) CreateBatch(ctx context.Context, data []*gorm_gen_model.AdminRoleDemo, batchSize int) error {
	dao := gorm_gen_dao.Use(a.db).AdminRoleDemo
	err := dao.WithContext(ctx).CreateInBatches(data, batchSize)
	if err != nil {
		return err
	}
	return nil
}

// CreateBatchCache 批量创建数据, 并删除缓存
func (a *AdminRoleDemoRepo) CreateBatchCache(ctx context.Context, data []*gorm_gen_model.AdminRoleDemo, batchSize int) error {
	dao := gorm_gen_dao.Use(a.db).AdminRoleDemo
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
func (a *AdminRoleDemoRepo) CreateBatchByTx(ctx context.Context, tx *gorm_gen_dao.Query, data []*gorm_gen_model.AdminRoleDemo, batchSize int) error {
	dao := tx.AdminRoleDemo
	err := dao.WithContext(ctx).CreateInBatches(data, batchSize)
	if err != nil {
		return err
	}
	return nil
}

// CreateBatchCacheByTx 批量创建数据(事务), 并删除缓存
func (a *AdminRoleDemoRepo) CreateBatchCacheByTx(ctx context.Context, tx *gorm_gen_dao.Query, data []*gorm_gen_model.AdminRoleDemo, batchSize int) error {
	dao := tx.AdminRoleDemo
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

// UpsertOne Upsert一条数据
// Update all columns, except primary keys, to new value on conflict
func (a *AdminRoleDemoRepo) UpsertOne(ctx context.Context, data *gorm_gen_model.AdminRoleDemo) error {
	dao := gorm_gen_dao.Use(a.db).AdminRoleDemo
	err := dao.WithContext(ctx).Save(data)
	if err != nil {
		return err
	}
	return nil
}

// UpsertOneCache Upsert一条数据, 并删除缓存
// Update all columns, except primary keys, to new value on conflict
func (a *AdminRoleDemoRepo) UpsertOneCache(ctx context.Context, data *gorm_gen_model.AdminRoleDemo) error {
	dao := gorm_gen_dao.Use(a.db).AdminRoleDemo
	oldData, err := dao.WithContext(ctx).Where(dao.ID.Eq(data.ID)).First()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	err = dao.WithContext(ctx).Save(data)
	if err != nil {
		return err
	}
	err = a.DeleteIndexCache(ctx, oldData, data)
	if err != nil {
		return err
	}
	return nil
}

// UpsertOneByTx Upsert一条数据(事务)
// Update all columns, except primary keys, to new value on conflict
func (a *AdminRoleDemoRepo) UpsertOneByTx(ctx context.Context, tx *gorm_gen_dao.Query, data *gorm_gen_model.AdminRoleDemo) error {
	dao := tx.AdminRoleDemo
	err := dao.WithContext(ctx).Save(data)
	if err != nil {
		return err
	}
	return nil
}

// UpsertOneCacheByTx Upsert一条数据(事务), 并删除缓存
// Update all columns, except primary keys, to new value on conflict
func (a *AdminRoleDemoRepo) UpsertOneCacheByTx(ctx context.Context, tx *gorm_gen_dao.Query, data *gorm_gen_model.AdminRoleDemo) error {
	dao := tx.AdminRoleDemo
	oldData, err := dao.WithContext(ctx).Where(dao.ID.Eq(data.ID)).First()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	err = dao.WithContext(ctx).Save(data)
	if err != nil {
		return err
	}
	err = a.DeleteIndexCache(ctx, oldData, data)
	if err != nil {
		return err
	}
	return nil
}

// UpsertOneByFields 根据fields字段Upsert一条数据
func (a *AdminRoleDemoRepo) UpsertOneByFields(ctx context.Context, data *gorm_gen_model.AdminRoleDemo, fields []string) error {
	if len(fields) == 0 {
		return errors.New("UpsertOneByFields fields is empty")
	}
	columns := make([]clause.Column, 0)
	for _, v := range fields {
		columns = append(columns, clause.Column{Name: v})
	}
	dao := gorm_gen_dao.Use(a.db).AdminRoleDemo
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
func (a *AdminRoleDemoRepo) UpsertOneCacheByFields(ctx context.Context, data *gorm_gen_model.AdminRoleDemo, fields []string) error {
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
	oldData := &gorm_gen_model.AdminRoleDemo{}
	err := a.db.Model(&gorm_gen_model.AdminRoleDemo{}).Clauses(whereExpressions...).First(oldData).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	dao := gorm_gen_dao.Use(a.db).AdminRoleDemo
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
func (a *AdminRoleDemoRepo) UpsertOneByFieldsTx(ctx context.Context, tx *gorm_gen_dao.Query, data *gorm_gen_model.AdminRoleDemo, fields []string) error {
	if len(fields) == 0 {
		return errors.New("UpsertOneByFieldsTx fields is empty")
	}
	columns := make([]clause.Column, 0)
	for _, v := range fields {
		columns = append(columns, clause.Column{Name: v})
	}
	dao := tx.AdminRoleDemo
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
func (a *AdminRoleDemoRepo) UpsertOneCacheByFieldsTx(ctx context.Context, tx *gorm_gen_dao.Query, data *gorm_gen_model.AdminRoleDemo, fields []string) error {
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
	oldData := &gorm_gen_model.AdminRoleDemo{}
	err := a.db.Model(&gorm_gen_model.AdminRoleDemo{}).Clauses(whereExpressions...).First(oldData).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	dao := tx.AdminRoleDemo
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

// UpdateOne 更新一条数据
// data 中主键字段必须有值，零值不会被更新
func (a *AdminRoleDemoRepo) UpdateOne(ctx context.Context, newData *gorm_gen_model.AdminRoleDemo) error {
	dao := gorm_gen_dao.Use(a.db).AdminRoleDemo
	_, err := dao.WithContext(ctx).Updates(newData)
	if err != nil {
		return err
	}
	return nil
}

// UpdateOneCache 更新一条数据，并删除缓存
// data 中主键字段必须有值，零值不会被更新
// oldData 旧数据，删除缓存时使用
func (a *AdminRoleDemoRepo) UpdateOneCache(ctx context.Context, newData *gorm_gen_model.AdminRoleDemo, oldData *gorm_gen_model.AdminRoleDemo) error {
	dao := gorm_gen_dao.Use(a.db).AdminRoleDemo
	_, err := dao.WithContext(ctx).Updates(newData)
	if err != nil {
		return err
	}
	err = a.DeleteIndexCache(ctx, oldData, newData)
	if err != nil {
		return err
	}
	return nil
}

// UpdateOneByTx 更新一条数据(事务)
// data 中主键字段必须有值，零值不会被更新
func (a *AdminRoleDemoRepo) UpdateOneByTx(ctx context.Context, tx *gorm_gen_dao.Query, newData *gorm_gen_model.AdminRoleDemo) error {
	dao := tx.AdminRoleDemo
	_, err := dao.WithContext(ctx).Updates(newData)
	if err != nil {
		return err
	}
	return err
}

// UpdateOneCacheByTx 更新一条数据(事务)，并删除缓存
// data 中主键字段必须有值，零值不会被更新
// oldData 旧数据，删除缓存时使用
func (a *AdminRoleDemoRepo) UpdateOneCacheByTx(ctx context.Context, tx *gorm_gen_dao.Query, newData *gorm_gen_model.AdminRoleDemo, oldData *gorm_gen_model.AdminRoleDemo) error {
	dao := tx.AdminRoleDemo
	_, err := dao.WithContext(ctx).Updates(newData)
	if err != nil {
		return err
	}
	err = a.DeleteIndexCache(ctx, oldData, newData)
	if err != nil {
		return err
	}
	return err
}

// UpdateOneWithZero 更新一条数据,包含零值
// data 中主键字段必须有值,并且会更新所有字段,包括零值
func (a *AdminRoleDemoRepo) UpdateOneWithZero(ctx context.Context, newData *gorm_gen_model.AdminRoleDemo) error {
	dao := gorm_gen_dao.Use(a.db).AdminRoleDemo
	_, err := dao.WithContext(ctx).Select(dao.ALL.WithTable("")).Updates(newData)
	if err != nil {
		return err
	}
	return nil
}

// UpdateOneCacheWithZero 更新一条数据,包含零值，并删除缓存
// data 中主键字段必须有值,并且会更新所有字段,包括零值
// oldData 旧数据，删除缓存时使用
func (a *AdminRoleDemoRepo) UpdateOneCacheWithZero(ctx context.Context, newData *gorm_gen_model.AdminRoleDemo, oldData *gorm_gen_model.AdminRoleDemo) error {
	dao := gorm_gen_dao.Use(a.db).AdminRoleDemo
	_, err := dao.WithContext(ctx).Select(dao.ALL.WithTable("")).Updates(newData)
	if err != nil {
		return err
	}
	err = a.DeleteIndexCache(ctx, oldData, newData)
	if err != nil {
		return err
	}
	return nil
}

// UpdateOneWithZeroByTx 更新一条数据(事务),包含零值，
// data 中主键字段必须有值,并且会更新所有字段,包括零值
func (a *AdminRoleDemoRepo) UpdateOneWithZeroByTx(ctx context.Context, tx *gorm_gen_dao.Query, newData *gorm_gen_model.AdminRoleDemo) error {
	dao := tx.AdminRoleDemo
	_, err := dao.WithContext(ctx).Select(dao.ALL.WithTable("")).Updates(newData)
	if err != nil {
		return err
	}
	return err
}

// UpdateOneCacheWithZeroByTx 更新一条数据(事务),包含零值，并删除缓存
// data 中主键字段必须有值,并且会更新所有字段,包括零值
// oldData 旧数据，删除缓存时使用
func (a *AdminRoleDemoRepo) UpdateOneCacheWithZeroByTx(ctx context.Context, tx *gorm_gen_dao.Query, newData *gorm_gen_model.AdminRoleDemo, oldData *gorm_gen_model.AdminRoleDemo) error {
	dao := tx.AdminRoleDemo
	_, err := dao.WithContext(ctx).Select(dao.ALL.WithTable("")).Updates(newData)
	if err != nil {
		return err
	}
	err = a.DeleteIndexCache(ctx, oldData, newData)
	if err != nil {
		return err
	}
	return err
}

// UpdateBatchByIDS 根据主键IDS批量更新
// 零值会被更新
func (a *AdminRoleDemoRepo) UpdateBatchByIDS(ctx context.Context, IDS []string, data map[string]interface{}) error {
	dao := gorm_gen_dao.Use(a.db).AdminRoleDemo
	_, err := dao.WithContext(ctx).Where(dao.ID.In(IDS...)).Updates(data)
	if err != nil {
		return err
	}
	return nil
}

// UpdateBatchByIDSTx 根据主键IDS批量更新(事务)
// 零值会被更新
func (a *AdminRoleDemoRepo) UpdateBatchByIDSTx(ctx context.Context, tx *gorm_gen_dao.Query, IDS []string, data map[string]interface{}) error {
	dao := tx.AdminRoleDemo
	_, err := dao.WithContext(ctx).Where(dao.ID.In(IDS...)).Updates(data)
	if err != nil {
		return err
	}
	return nil
}

// FindOneByID 根据ID查询一条数据
func (a *AdminRoleDemoRepo) FindOneByID(ctx context.Context, ID string) (*gorm_gen_model.AdminRoleDemo, error) {
	dao := gorm_gen_dao.Use(a.db).AdminRoleDemo
	result, err := dao.WithContext(ctx).Where(dao.ID.Eq(ID)).First()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return result, nil
}

// FindOneCacheByID 根据ID查询一条数据，并设置缓存
func (a *AdminRoleDemoRepo) FindOneCacheByID(ctx context.Context, ID string) (*gorm_gen_model.AdminRoleDemo, error) {
	resp := new(gorm_gen_model.AdminRoleDemo)
	cacheKey := a.cache.Key(CacheAdminRoleDemoByIDPrefix, ID)
	cacheValue, err := a.cache.Fetch(ctx, cacheKey, func() (string, error) {
		dao := gorm_gen_dao.Use(a.db).AdminRoleDemo
		result, err := dao.WithContext(ctx).Where(dao.ID.Eq(ID)).First()
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
		err = a.encoding.Unmarshal([]byte(cacheValue), resp)
		if err != nil {
			return nil, err
		}
	}
	return resp, nil
}

// FindMultiByIDS 根据IDS查询多条数据
func (a *AdminRoleDemoRepo) FindMultiByIDS(ctx context.Context, IDS []string) ([]*gorm_gen_model.AdminRoleDemo, error) {
	dao := gorm_gen_dao.Use(a.db).AdminRoleDemo
	result, err := dao.WithContext(ctx).Where(dao.ID.In(IDS...)).Find()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FindMultiCacheByIDS 根据IDS查询多条数据，并设置缓存
func (a *AdminRoleDemoRepo) FindMultiCacheByIDS(ctx context.Context, IDS []string) ([]*gorm_gen_model.AdminRoleDemo, error) {
	resp := make([]*gorm_gen_model.AdminRoleDemo, 0)
	cacheKeys := make([]string, 0)
	keyToParam := make(map[string]string)
	for _, v := range IDS {
		cacheKey := a.cache.Key(CacheAdminRoleDemoByIDPrefix, v)
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
		dao := gorm_gen_dao.Use(a.db).AdminRoleDemo
		result, err := dao.WithContext(ctx).Where(dao.ID.In(params...)).Find()
		if err != nil {
			return nil, err
		}
		for _, v := range result {
			marshal, err := a.encoding.Marshal(v)
			if err != nil {
				return nil, err
			}
			dbValue[a.cache.Key(CacheAdminRoleDemoByIDPrefix, v.ID)] = string(marshal)
		}
		return dbValue, nil
	})
	if err != nil {
		return nil, err
	}
	for _, v := range IDS {
		cacheKey := a.cache.Key(CacheAdminRoleDemoByIDPrefix, v)
		if cacheValue[cacheKey] != "" {
			tmp := new(gorm_gen_model.AdminRoleDemo)
			err := a.encoding.Unmarshal([]byte(cacheValue[cacheKey]), tmp)
			if err != nil {
				return nil, err
			}
			resp = append(resp, tmp)
		}
	}
	return resp, nil
}

// FindMultiByCondition 自定义查询数据(通用)
func (a *AdminRoleDemoRepo) FindMultiByCondition(ctx context.Context, conditionReq *condition.Req) ([]*gorm_gen_model.AdminRoleDemo, *condition.Reply, error) {
	result := make([]*gorm_gen_model.AdminRoleDemo, 0)
	conditionReply := &condition.Reply{}
	var total int64
	whereExpressions, orderExpressions, err := conditionReq.ConvertToGormExpression(gorm_gen_model.AdminRoleDemo{})
	if err != nil {
		return result, conditionReply, err
	}
	err = a.db.WithContext(ctx).Model(&gorm_gen_model.AdminRoleDemo{}).Select([]string{"*"}).Clauses(whereExpressions...).Count(&total).Error
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
	query := a.db.WithContext(ctx).Model(&gorm_gen_model.AdminRoleDemo{}).Clauses(whereExpressions...).Clauses(orderExpressions...)
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

// DeleteOneByID 根据ID删除一条数据
func (a *AdminRoleDemoRepo) DeleteOneByID(ctx context.Context, ID string) error {
	dao := gorm_gen_dao.Use(a.db).AdminRoleDemo
	_, err := dao.WithContext(ctx).Where(dao.ID.Eq(ID)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneCacheByID 根据ID删除一条数据，并删除缓存
func (a *AdminRoleDemoRepo) DeleteOneCacheByID(ctx context.Context, ID string) error {
	dao := gorm_gen_dao.Use(a.db).AdminRoleDemo
	result, err := dao.WithContext(ctx).Where(dao.ID.Eq(ID)).First()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if result == nil {
		return nil
	}
	_, err = dao.WithContext(ctx).Where(dao.ID.Eq(ID)).Delete()
	if err != nil {
		return err
	}
	err = a.DeleteIndexCache(ctx, result)
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneByID 根据ID删除一条数据
func (a *AdminRoleDemoRepo) DeleteOneByIDTx(ctx context.Context, tx *gorm_gen_dao.Query, ID string) error {
	dao := tx.AdminRoleDemo
	_, err := dao.WithContext(ctx).Where(dao.ID.Eq(ID)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneCacheByIDTx 根据ID删除一条数据，并删除缓存
func (a *AdminRoleDemoRepo) DeleteOneCacheByIDTx(ctx context.Context, tx *gorm_gen_dao.Query, ID string) error {
	dao := tx.AdminRoleDemo
	result, err := dao.WithContext(ctx).Where(dao.ID.Eq(ID)).First()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if result == nil {
		return nil
	}
	_, err = dao.WithContext(ctx).Where(dao.ID.Eq(ID)).Delete()
	if err != nil {
		return err
	}
	err = a.DeleteIndexCache(ctx, result)
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiByIDS 根据IDS删除多条数据
func (a *AdminRoleDemoRepo) DeleteMultiByIDS(ctx context.Context, IDS []string) error {
	dao := gorm_gen_dao.Use(a.db).AdminRoleDemo
	_, err := dao.WithContext(ctx).Where(dao.ID.In(IDS...)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiCacheByIDS 根据IDS删除多条数据，并删除缓存
func (a *AdminRoleDemoRepo) DeleteMultiCacheByIDS(ctx context.Context, IDS []string) error {
	dao := gorm_gen_dao.Use(a.db).AdminRoleDemo
	result, err := dao.WithContext(ctx).Where(dao.ID.In(IDS...)).Find()
	if err != nil {
		return err
	}
	if len(result) == 0 {
		return nil
	}
	_, err = dao.WithContext(ctx).Where(dao.ID.In(IDS...)).Delete()
	if err != nil {
		return err
	}
	err = a.DeleteIndexCache(ctx, result...)
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiByIDSTx 根据IDS删除多条数据
func (a *AdminRoleDemoRepo) DeleteMultiByIDSTx(ctx context.Context, tx *gorm_gen_dao.Query, IDS []string) error {
	dao := tx.AdminRoleDemo
	_, err := dao.WithContext(ctx).Where(dao.ID.In(IDS...)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiCacheByIDSTx 根据IDS删除多条数据，并删除缓存
func (a *AdminRoleDemoRepo) DeleteMultiCacheByIDSTx(ctx context.Context, tx *gorm_gen_dao.Query, IDS []string) error {
	dao := tx.AdminRoleDemo
	result, err := dao.WithContext(ctx).Where(dao.ID.In(IDS...)).Find()
	if err != nil {
		return err
	}
	if len(result) == 0 {
		return nil
	}
	_, err = dao.WithContext(ctx).Where(dao.ID.In(IDS...)).Delete()
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
func (a *AdminRoleDemoRepo) DeleteIndexCache(ctx context.Context, data ...*gorm_gen_model.AdminRoleDemo) error {
	KeyMap := make(map[string]struct{})
	for _, v := range data {
		if v != nil {
			KeyMap[a.cache.Key(CacheAdminRoleDemoByIDPrefix, v.ID)] = struct{}{}

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

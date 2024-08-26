// Code generated by gen/repo. DO NOT EDIT.
// Code generated by gen/repo. DO NOT EDIT.
// Code generated by gen/repo. DO NOT EDIT.

package gorm_gen_repo

import (
	"context"
	"errors"

	"github.com/fzf-labs/fdatabase/orm/condition"
	"github.com/fzf-labs/fdatabase/orm/dbcache"
	"github.com/fzf-labs/fdatabase/orm/encoding"
	"github.com/fzf-labs/fdatabase/orm/gen/config"
	"github.com/fzf-labs/fdatabase/orm/gen/example/postgres/gorm_gen_dao"
	"github.com/fzf-labs/fdatabase/orm/gen/example/postgres/gorm_gen_model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var _ IAdminLogDemoRepo = (*AdminLogDemoRepo)(nil)

var (
	CacheAdminLogDemoAll        = "DBCache:gorm_gen:AdminLogDemoAll"
	CacheAdminLogDemoByIDPrefix = "DBCache:gorm_gen:AdminLogDemoByID"
)

type (
	IAdminLogDemoRepo interface {
		// CreateOne 创建一条数据
		CreateOne(ctx context.Context, data *gorm_gen_model.AdminLogDemo) error
		// CreateOneByTx 创建一条数据(事务)
		CreateOneByTx(ctx context.Context, tx *gorm_gen_dao.Query, data *gorm_gen_model.AdminLogDemo) error
		// CreateBatch 批量创建数据
		CreateBatch(ctx context.Context, data []*gorm_gen_model.AdminLogDemo, batchSize int) error
		// CreateBatchByTx 批量创建数据(事务)
		CreateBatchByTx(ctx context.Context, tx *gorm_gen_dao.Query, data []*gorm_gen_model.AdminLogDemo, batchSize int) error
		// UpsertOne Upsert一条数据
		UpsertOne(ctx context.Context, data *gorm_gen_model.AdminLogDemo) error
		// UpsertOneByTx Upsert一条数据(事务)
		UpsertOneByTx(ctx context.Context, tx *gorm_gen_dao.Query, data *gorm_gen_model.AdminLogDemo) error
		// UpsertOneByFields Upsert一条数据，根据fields字段
		UpsertOneByFields(ctx context.Context, data *gorm_gen_model.AdminLogDemo, fields []string) error
		// UpsertOneByFieldsTx Upsert一条数据，根据fields字段(事务)
		UpsertOneByFieldsTx(ctx context.Context, tx *gorm_gen_dao.Query, data *gorm_gen_model.AdminLogDemo, fields []string) error
		// UpdateOne 更新一条数据
		UpdateOne(ctx context.Context, data *gorm_gen_model.AdminLogDemo) error
		// UpdateOneByTx 更新一条数据(事务)
		UpdateOneByTx(ctx context.Context, tx *gorm_gen_dao.Query, data *gorm_gen_model.AdminLogDemo) error
		// UpdateOneWithZero 更新一条数据,包含零值
		UpdateOneWithZero(ctx context.Context, data *gorm_gen_model.AdminLogDemo) error
		// UpdateOneWithZeroByTx 更新一条数据,包含零值(事务)
		UpdateOneWithZeroByTx(ctx context.Context, tx *gorm_gen_dao.Query, data *gorm_gen_model.AdminLogDemo) error
		// FindOneCacheByID 根据ID查询一条数据并设置缓存
		FindOneCacheByID(ctx context.Context, ID string) (*gorm_gen_model.AdminLogDemo, error)
		// FindOneByID 根据ID查询一条数据
		FindOneByID(ctx context.Context, ID string) (*gorm_gen_model.AdminLogDemo, error)
		// FindMultiCacheByIDS 根据IDS查询多条数据并设置缓存
		FindMultiCacheByIDS(ctx context.Context, IDS []string) ([]*gorm_gen_model.AdminLogDemo, error)
		// FindMultiByIDS 根据IDS查询多条数据
		FindMultiByIDS(ctx context.Context, IDS []string) ([]*gorm_gen_model.AdminLogDemo, error)
		// FindAll 查询所有数据
		FindAll(ctx context.Context) ([]*gorm_gen_model.AdminLogDemo, error)
		// FindAllCache 查询所有数据并设置缓存
		FindAllCache(ctx context.Context) ([]*gorm_gen_model.AdminLogDemo, error)
		// FindMultiByCondition 根据自定义条件查询数据
		FindMultiByCondition(ctx context.Context, conditionReq *condition.Req) ([]*gorm_gen_model.AdminLogDemo, *condition.Reply, error)
		// DeleteOneCacheByID 根据ID删除一条数据并清理缓存
		DeleteOneCacheByID(ctx context.Context, ID string) error
		// DeleteOneCacheByIDTx 根据ID删除一条数据并清理缓存(事务)
		DeleteOneCacheByIDTx(ctx context.Context, tx *gorm_gen_dao.Query, ID string) error
		// DeleteOneByID 根据ID删除一条数据
		DeleteOneByID(ctx context.Context, ID string) error
		// DeleteOneByIDTx 根据ID删除一条数据(事务)
		DeleteOneByIDTx(ctx context.Context, tx *gorm_gen_dao.Query, ID string) error
		// DeleteMultiCacheByIDS 根据IDS删除多条数据并清理缓存
		DeleteMultiCacheByIDS(ctx context.Context, IDS []string) error
		// DeleteMultiCacheByIDSTx 根据IDS删除多条数据并清理缓存(事务)
		DeleteMultiCacheByIDSTx(ctx context.Context, tx *gorm_gen_dao.Query, IDS []string) error
		// DeleteMultiByIDS 根据IDS删除多条数据
		DeleteMultiByIDS(ctx context.Context, IDS []string) error
		// DeleteMultiByIDSTx 根据IDS删除多条数据(事务)
		DeleteMultiByIDSTx(ctx context.Context, tx *gorm_gen_dao.Query, IDS []string) error
		// DeleteAllCache 删除所有数据缓存
		DeleteAllCache(ctx context.Context) error
		// DeleteUniqueIndexCache 删除唯一索引存在的缓存
		DeleteUniqueIndexCache(ctx context.Context, data []*gorm_gen_model.AdminLogDemo) error
	}
	AdminLogDemoRepo struct {
		db       *gorm.DB
		cache    dbcache.IDBCache
		encoding encoding.API
	}
)

func NewAdminLogDemoRepo(cfg *config.Repo) *AdminLogDemoRepo {
	return &AdminLogDemoRepo{
		db:       cfg.DB,
		cache:    cfg.Cache,
		encoding: cfg.Encoding,
	}
}

// CreateOne 创建一条数据
func (a *AdminLogDemoRepo) CreateOne(ctx context.Context, data *gorm_gen_model.AdminLogDemo) error {
	dao := gorm_gen_dao.Use(a.db).AdminLogDemo
	err := dao.WithContext(ctx).Create(data)
	if err != nil {
		return err
	}
	err = a.DeleteUniqueIndexCache(ctx, []*gorm_gen_model.AdminLogDemo{data})
	if err != nil {
		return err
	}
	return nil
}

// CreateOneByTx 创建一条数据(事务)
func (a *AdminLogDemoRepo) CreateOneByTx(ctx context.Context, tx *gorm_gen_dao.Query, data *gorm_gen_model.AdminLogDemo) error {
	dao := tx.AdminLogDemo
	err := dao.WithContext(ctx).Create(data)
	if err != nil {
		return err
	}
	err = a.DeleteUniqueIndexCache(ctx, []*gorm_gen_model.AdminLogDemo{data})
	if err != nil {
		return err
	}
	return nil
}

// CreateBatch 批量创建数据
func (a *AdminLogDemoRepo) CreateBatch(ctx context.Context, data []*gorm_gen_model.AdminLogDemo, batchSize int) error {
	dao := gorm_gen_dao.Use(a.db).AdminLogDemo
	err := dao.WithContext(ctx).CreateInBatches(data, batchSize)
	if err != nil {
		return err
	}
	err = a.DeleteUniqueIndexCache(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

// CreateBatchByTx 批量创建数据(事务)
func (a *AdminLogDemoRepo) CreateBatchByTx(ctx context.Context, tx *gorm_gen_dao.Query, data []*gorm_gen_model.AdminLogDemo, batchSize int) error {
	dao := tx.AdminLogDemo
	err := dao.WithContext(ctx).CreateInBatches(data, batchSize)
	if err != nil {
		return err
	}
	err = a.DeleteUniqueIndexCache(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

// UpsertOne Upsert一条数据
func (a *AdminLogDemoRepo) UpsertOne(ctx context.Context, data *gorm_gen_model.AdminLogDemo) error {
	dao := gorm_gen_dao.Use(a.db).AdminLogDemo
	err := dao.WithContext(ctx).Save(data)
	if err != nil {
		return err
	}
	err = a.DeleteUniqueIndexCache(ctx, []*gorm_gen_model.AdminLogDemo{data})
	if err != nil {
		return err
	}
	return nil
}

// UpsertOneByTx Upsert一条数据(事务)
func (a *AdminLogDemoRepo) UpsertOneByTx(ctx context.Context, tx *gorm_gen_dao.Query, data *gorm_gen_model.AdminLogDemo) error {
	dao := tx.AdminLogDemo
	err := dao.WithContext(ctx).Save(data)
	if err != nil {
		return err
	}
	err = a.DeleteUniqueIndexCache(ctx, []*gorm_gen_model.AdminLogDemo{data})
	if err != nil {
		return err
	}
	return nil
}

// UpsertOneByFields Upsert一条数据，根据fields字段
func (a *AdminLogDemoRepo) UpsertOneByFields(ctx context.Context, data *gorm_gen_model.AdminLogDemo, fields []string) error {
	if len(fields) == 0 {
		return errors.New("UpsertOneByFields fields is empty")
	}
	columns := make([]clause.Column, 0)
	for _, v := range fields {
		columns = append(columns, clause.Column{Name: v})
	}
	dao := gorm_gen_dao.Use(a.db).AdminLogDemo
	err := dao.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   columns,
		UpdateAll: true,
	}).Create(data)
	if err != nil {
		return err
	}
	err = a.DeleteUniqueIndexCache(ctx, []*gorm_gen_model.AdminLogDemo{data})
	if err != nil {
		return err
	}
	return nil
}

// UpsertOneByFieldsTx Upsert一条数据，根据fields字段(事务)
func (a *AdminLogDemoRepo) UpsertOneByFieldsTx(ctx context.Context, tx *gorm_gen_dao.Query, data *gorm_gen_model.AdminLogDemo, fields []string) error {
	if len(fields) == 0 {
		return errors.New("UpsertOneByFieldsTx fields is empty")
	}
	columns := make([]clause.Column, 0)
	for _, v := range fields {
		columns = append(columns, clause.Column{Name: v})
	}
	dao := tx.AdminLogDemo
	err := dao.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   columns,
		UpdateAll: true,
	}).Create(data)
	if err != nil {
		return err
	}
	err = a.DeleteUniqueIndexCache(ctx, []*gorm_gen_model.AdminLogDemo{data})
	if err != nil {
		return err
	}
	return nil
}

// UpdateOne 更新一条数据
func (a *AdminLogDemoRepo) UpdateOne(ctx context.Context, data *gorm_gen_model.AdminLogDemo) error {
	dao := gorm_gen_dao.Use(a.db).AdminLogDemo
	_, err := dao.WithContext(ctx).Where(dao.ID.Eq(data.ID)).Updates(data)
	if err != nil {
		return err
	}
	err = a.DeleteUniqueIndexCache(ctx, []*gorm_gen_model.AdminLogDemo{data})
	if err != nil {
		return err
	}
	return nil
}

// UpdateOneByTx 更新一条数据(事务)
func (a *AdminLogDemoRepo) UpdateOneByTx(ctx context.Context, tx *gorm_gen_dao.Query, data *gorm_gen_model.AdminLogDemo) error {
	dao := tx.AdminLogDemo
	_, err := dao.WithContext(ctx).Where(dao.ID.Eq(data.ID)).Updates(data)
	if err != nil {
		return err
	}
	err = a.DeleteUniqueIndexCache(ctx, []*gorm_gen_model.AdminLogDemo{data})
	if err != nil {
		return err
	}
	return err
}

// UpdateOneWithZero 更新一条数据,包含零值
func (a *AdminLogDemoRepo) UpdateOneWithZero(ctx context.Context, data *gorm_gen_model.AdminLogDemo) error {
	dao := gorm_gen_dao.Use(a.db).AdminLogDemo
	_, err := dao.WithContext(ctx).Where(dao.ID.Eq(data.ID)).Select(dao.ALL.WithTable("")).Updates(data)
	if err != nil {
		return err
	}
	err = a.DeleteUniqueIndexCache(ctx, []*gorm_gen_model.AdminLogDemo{data})
	if err != nil {
		return err
	}
	return nil
}

// UpdateOneWithZeroByTx 更新一条数据(事务),包含零值
func (a *AdminLogDemoRepo) UpdateOneWithZeroByTx(ctx context.Context, tx *gorm_gen_dao.Query, data *gorm_gen_model.AdminLogDemo) error {
	dao := tx.AdminLogDemo
	_, err := dao.WithContext(ctx).Where(dao.ID.Eq(data.ID)).Select(dao.ALL.WithTable("")).Updates(data)
	if err != nil {
		return err
	}
	err = a.DeleteUniqueIndexCache(ctx, []*gorm_gen_model.AdminLogDemo{data})
	if err != nil {
		return err
	}
	return err
}

// DeleteOneCacheByID 根据ID删除一条数据并清理缓存
func (a *AdminLogDemoRepo) DeleteOneCacheByID(ctx context.Context, ID string) error {
	dao := gorm_gen_dao.Use(a.db).AdminLogDemo
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
	err = a.DeleteUniqueIndexCache(ctx, []*gorm_gen_model.AdminLogDemo{result})
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneCacheByIDTx 根据ID删除一条数据并清理缓存
func (a *AdminLogDemoRepo) DeleteOneCacheByIDTx(ctx context.Context, tx *gorm_gen_dao.Query, ID string) error {
	dao := tx.AdminLogDemo
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
	err = a.DeleteUniqueIndexCache(ctx, []*gorm_gen_model.AdminLogDemo{result})
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneByID 根据ID删除一条数据
func (a *AdminLogDemoRepo) DeleteOneByID(ctx context.Context, ID string) error {
	dao := gorm_gen_dao.Use(a.db).AdminLogDemo
	_, err := dao.WithContext(ctx).Where(dao.ID.Eq(ID)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneByID 根据ID删除一条数据
func (a *AdminLogDemoRepo) DeleteOneByIDTx(ctx context.Context, tx *gorm_gen_dao.Query, ID string) error {
	dao := tx.AdminLogDemo
	_, err := dao.WithContext(ctx).Where(dao.ID.Eq(ID)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiCacheByIDS 根据IDS删除多条数据并清理缓存
func (a *AdminLogDemoRepo) DeleteMultiCacheByIDS(ctx context.Context, IDS []string) error {
	dao := gorm_gen_dao.Use(a.db).AdminLogDemo
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
	err = a.DeleteUniqueIndexCache(ctx, result)
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiCacheByIDSTx 根据IDS删除多条数据并清理缓存
func (a *AdminLogDemoRepo) DeleteMultiCacheByIDSTx(ctx context.Context, tx *gorm_gen_dao.Query, IDS []string) error {
	dao := tx.AdminLogDemo
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
	err = a.DeleteUniqueIndexCache(ctx, result)
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiByIDS 根据IDS删除多条数据
func (a *AdminLogDemoRepo) DeleteMultiByIDS(ctx context.Context, IDS []string) error {
	dao := gorm_gen_dao.Use(a.db).AdminLogDemo
	_, err := dao.WithContext(ctx).Where(dao.ID.In(IDS...)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiByIDSTx 根据IDS删除多条数据
func (a *AdminLogDemoRepo) DeleteMultiByIDSTx(ctx context.Context, tx *gorm_gen_dao.Query, IDS []string) error {
	dao := tx.AdminLogDemo
	_, err := dao.WithContext(ctx).Where(dao.ID.In(IDS...)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteAllCache 删除所有数据缓存
func (a *AdminLogDemoRepo) DeleteAllCache(ctx context.Context) error {
	cacheKey := a.cache.Key(CacheAdminLogDemoAll)
	err := a.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUniqueIndexCache 删除唯一索引存在的缓存
func (a *AdminLogDemoRepo) DeleteUniqueIndexCache(ctx context.Context, data []*gorm_gen_model.AdminLogDemo) error {
	keys := make([]string, 0)
	for _, v := range data {
		keys = append(keys, a.cache.Key(CacheAdminLogDemoByIDPrefix, v.ID))

	}
	err := a.cache.DelBatch(ctx, keys)
	if err != nil {
		return err
	}
	return nil
}

// FindOneCacheByID 根据ID查询一条数据并设置缓存
func (a *AdminLogDemoRepo) FindOneCacheByID(ctx context.Context, ID string) (*gorm_gen_model.AdminLogDemo, error) {
	resp := new(gorm_gen_model.AdminLogDemo)
	cacheKey := a.cache.Key(CacheAdminLogDemoByIDPrefix, ID)
	cacheValue, err := a.cache.Fetch(ctx, cacheKey, func() (string, error) {
		dao := gorm_gen_dao.Use(a.db).AdminLogDemo
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

// FindOneByID 根据ID查询一条数据
func (a *AdminLogDemoRepo) FindOneByID(ctx context.Context, ID string) (*gorm_gen_model.AdminLogDemo, error) {
	dao := gorm_gen_dao.Use(a.db).AdminLogDemo
	result, err := dao.WithContext(ctx).Where(dao.ID.Eq(ID)).First()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return result, nil
}

// FindMultiCacheByIDS 根据IDS查询多条数据并设置缓存
func (a *AdminLogDemoRepo) FindMultiCacheByIDS(ctx context.Context, IDS []string) ([]*gorm_gen_model.AdminLogDemo, error) {
	resp := make([]*gorm_gen_model.AdminLogDemo, 0)
	cacheKeys := make([]string, 0)
	keyToParam := make(map[string]string)
	for _, v := range IDS {
		cacheKey := a.cache.Key(CacheAdminLogDemoByIDPrefix, v)
		cacheKeys = append(cacheKeys, cacheKey)
		keyToParam[cacheKey] = v
	}
	cacheValue, err := a.cache.FetchBatch(ctx, cacheKeys, func(miss []string) (map[string]string, error) {
		parameters := make([]string, 0)
		for _, v := range miss {
			parameters = append(parameters, keyToParam[v])
		}
		dao := gorm_gen_dao.Use(a.db).AdminLogDemo
		result, err := dao.WithContext(ctx).Where(dao.ID.In(parameters...)).Find()
		if err != nil {
			return nil, err
		}
		value := make(map[string]string)
		for _, v := range miss {
			value[v] = ""
		}
		for _, v := range result {
			marshal, err := a.encoding.Marshal(v)
			if err != nil {
				return nil, err
			}
			value[a.cache.Key(CacheAdminLogDemoByIDPrefix, v.ID)] = string(marshal)
		}
		return value, nil
	})
	if err != nil {
		return nil, err
	}
	for _, v := range IDS {
		cacheKey := a.cache.Key(CacheAdminLogDemoByIDPrefix, v)
		if cacheValue[cacheKey] != "" {
			tmp := new(gorm_gen_model.AdminLogDemo)
			err := a.encoding.Unmarshal([]byte(cacheValue[cacheKey]), tmp)
			if err != nil {
				return nil, err
			}
			resp = append(resp, tmp)
		}
	}
	return resp, nil
}

// FindMultiByIDS 根据IDS查询多条数据
func (a *AdminLogDemoRepo) FindMultiByIDS(ctx context.Context, IDS []string) ([]*gorm_gen_model.AdminLogDemo, error) {
	dao := gorm_gen_dao.Use(a.db).AdminLogDemo
	result, err := dao.WithContext(ctx).Where(dao.ID.In(IDS...)).Find()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FindAll 查询所有数据
func (a *AdminLogDemoRepo) FindAll(ctx context.Context) ([]*gorm_gen_model.AdminLogDemo, error) {
	dao := gorm_gen_dao.Use(a.db).AdminLogDemo
	result, err := dao.WithContext(ctx).Find()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FindAllCache 查询所有数据并设置缓存
func (a *AdminLogDemoRepo) FindAllCache(ctx context.Context) ([]*gorm_gen_model.AdminLogDemo, error) {
	resp := make([]*gorm_gen_model.AdminLogDemo, 0)
	cacheKey := a.cache.Key(CacheAdminLogDemoAll)
	cacheValue, err := a.cache.Fetch(ctx, cacheKey, func() (string, error) {
		dao := gorm_gen_dao.Use(a.db).AdminLogDemo
		result, err := dao.WithContext(ctx).Find()
		if err != nil {
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

// FindMultiByCondition 自定义查询数据(通用)
func (a *AdminLogDemoRepo) FindMultiByCondition(ctx context.Context, conditionReq *condition.Req) ([]*gorm_gen_model.AdminLogDemo, *condition.Reply, error) {
	result := make([]*gorm_gen_model.AdminLogDemo, 0)
	var total int64
	whereExpressions, orderExpressions, err := conditionReq.ConvertToGormExpression(gorm_gen_model.AdminLogDemo{})
	if err != nil {
		return result, nil, err
	}
	err = a.db.WithContext(ctx).Model(&gorm_gen_model.AdminLogDemo{}).Select([]string{"*"}).Clauses(whereExpressions...).Count(&total).Error
	if err != nil {
		return result, nil, err
	}
	if total == 0 {
		return result, nil, nil
	}
	conditionReply, err := conditionReq.ConvertToPage(int32(total))
	if err != nil {
		return result, nil, err
	}
	query := a.db.WithContext(ctx).Model(&gorm_gen_model.AdminLogDemo{}).Clauses(whereExpressions...).Clauses(orderExpressions...)
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

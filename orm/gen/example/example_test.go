package postgres

import (
	"context"
	"fmt"
	"github.com/fzf-labs/fdatabase/orm"
	"github.com/fzf-labs/fdatabase/orm/condition"
	"github.com/fzf-labs/fdatabase/orm/dbcache/goredisdbcache"
	"github.com/fzf-labs/fdatabase/orm/encoding"
	"github.com/fzf-labs/fdatabase/orm/gen/config"
	"github.com/fzf-labs/fdatabase/orm/gen/example/postgres/gorm_gen_dao"
	"github.com/fzf-labs/fdatabase/orm/gen/example/postgres/gorm_gen_model"
	"github.com/fzf-labs/fdatabase/orm/gen/example/postgres/gorm_gen_repo"
	"github.com/go-redis/redismock/v9"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

var (
	db          *gorm.DB
	redisClient *redis.Client
	err         error
)

func init() {
	db, err = orm.NewGormPostgresClient(&orm.GormPostgresClientConfig{
		DataSourceName:  "host=0.0.0.0 port=5432 user=postgres password=123456 dbname=gorm_gen sslmode=disable TimeZone=Asia/Shanghai",
		MaxIdleConn:     0,
		MaxOpenConn:     0,
		ConnMaxLifeTime: 0,
		ShowLog:         true,
		Tracing:         false,
	})
	if err != nil {
		panic(err)
	}
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "0.0.0.0:6379",
		Password: "123456",
	})
}

// Test_FindOneCacheByID 根据ID查询单条数据
func Test_FindOneCacheByID(t *testing.T) {
	dbCache := goredisdbcache.NewGoRedisDBCache(redisClient)
	ctx := context.Background()
	cfg := config.NewRepoConfig(db, dbCache, encoding.NewMsgPack())
	repo := gorm_gen_repo.NewUserDemoRepo(cfg)
	result, err := repo.FindOneByID(ctx, "182a65a0-ee20-4fe0-a0e8-ba30edcf402b")
	if err != nil {
		return
	}
	fmt.Println(result)
	assert.Equal(t, nil, err)
}

func Test_FindMultiCacheByUsernames(t *testing.T) {
	dbCache := goredisdbcache.NewGoRedisDBCache(redisClient)
	ctx := context.Background()
	cfg := config.NewRepoConfig(db, dbCache, encoding.NewMsgPack())
	repo := gorm_gen_repo.NewUserDemoRepo(cfg)
	result, err := repo.FindMultiCacheByUsernames(ctx, []string{"a", "b", "c", "d", "e", "f", "g"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
	assert.Equal(t, nil, err)
}

func Test_UpdateOneCache(t *testing.T) {
	dbCache := goredisdbcache.NewGoRedisDBCache(redisClient)
	ctx := context.Background()
	cfg := config.NewRepoConfig(db, dbCache, encoding.NewMsgPack())
	repo := gorm_gen_repo.NewUserDemoRepo(cfg)
	data, err := repo.FindOneByID(ctx, "182a65a0-ee20-4fe0-a0e8-ba30edcf402b")
	if err != nil {
		return
	}
	data.Remark = "123"
	err = repo.UpdateOneCache(ctx, data)
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal(t, nil, err)
}

func Test_FindMultiCacheByTenantIDS(t *testing.T) {
	dbCache := goredisdbcache.NewGoRedisDBCache(redisClient)
	ctx := context.Background()
	cfg := config.NewRepoConfig(db, dbCache, encoding.NewMsgPack())
	repo := gorm_gen_repo.NewUserDemoRepo(cfg)
	result, err := repo.FindMultiCacheByTenantIDS(ctx, []int64{1, 2})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
	assert.Equal(t, nil, err)
}

// Test_FindMultiByCustom 自定义查询
func Test_FindMultiByCondition(t *testing.T) {
	client, _ := redismock.NewClientMock()
	dbCache := goredisdbcache.NewGoRedisDBCache(client)
	ctx := context.Background()
	cfg := config.NewRepoConfig(db, dbCache, encoding.NewMsgPack())
	repo := gorm_gen_repo.NewAdminDemoRepo(cfg)
	result, p, err := repo.FindMultiByCondition(ctx, &condition.Req{
		Page:     1,
		PageSize: 10,
		Order: []*condition.OrderParam{
			{
				Field: "created_at",
				Order: condition.DESC,
			},
		},
		Query: []*condition.QueryParam{
			{
				Field: "username",
				Value: "admin",
				Exp:   condition.EQ,
				Logic: condition.AND,
			},
			{
				Field: "username",
				Value: []interface{}{"admin", "admin2"},
				Exp:   condition.IN,
				Logic: "",
			},
			{
				Field: "username",
				Value: "123",
				Exp:   condition.LIKE,
				Logic: "",
			},
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", result)
	fmt.Printf("%+v\n", p)
	assert.Equal(t, nil, err)
}

// Test_Tx 使用事务
func Test_Tx(t *testing.T) {
	db, err := orm.NewGormPostgresClient(&orm.GormPostgresClientConfig{
		DataSourceName:  "host=0.0.0.0 port=5432 user=postgres password=123456 dbname=gorm_gen sslmode=disable TimeZone=Asia/Shanghai",
		MaxIdleConn:     0,
		MaxOpenConn:     0,
		ConnMaxLifeTime: 0,
		ShowLog:         false,
		Tracing:         false,
	})
	if err != nil {
		return
	}
	client, _ := redismock.NewClientMock()
	dbCache := goredisdbcache.NewGoRedisDBCache(client)
	ctx := context.Background()
	cfg := config.NewRepoConfig(db, dbCache, encoding.NewMsgPack())
	adminDemoRepo := gorm_gen_repo.NewAdminDemoRepo(cfg)
	adminLogDemoRepo := gorm_gen_repo.NewAdminLogDemoRepo(cfg)
	err = gorm_gen_dao.Use(db).Transaction(func(tx *gorm_gen_dao.Query) error {
		err2 := adminDemoRepo.UpsertOneByTx(ctx, tx, &gorm_gen_model.AdminDemo{
			ID:       "c8ddd930-339a-408b-8acb-fac22f5b43aa",
			Username: "admin",
			Nickname: "admin",
			Gender:   0,
			RoleIds:  nil,
			Salt:     "123",
			Status:   1,
		})
		if err2 != nil {
			return err2
		}
		err2 = adminLogDemoRepo.CreateOneByTx(ctx, tx, &gorm_gen_model.AdminLogDemo{
			AdminID:   "c8ddd930-339a-408b-8acb-fac22f5b43aa",
			IP:        "0.0.0.0",
			URI:       "www.baidu.com",
			Useragent: "apifox",
			Header:    nil,
			Req:       nil,
			Resp:      nil,
		})
		if err2 != nil {
			return err2
		}
		//return errors.New("rollback")
		return nil
	})
	if err != nil {
		return
	}
	assert.Equal(t, nil, err)
}

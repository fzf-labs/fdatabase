package postgres

import (
	"context"
	"fmt"
	"gitlab.yc345.tv/backend/utils/v2/orm/gen/config"
	"gitlab.yc345.tv/backend/utils/v2/orm/gen/encoding"
	"testing"
	"time"

	"github.com/go-redis/redismock/v8"
	"github.com/stretchr/testify/assert"
	"gitlab.yc345.tv/backend/utils/v2/orm"
	"gitlab.yc345.tv/backend/utils/v2/orm/gen/cache/goredisdbcache"
	"gitlab.yc345.tv/backend/utils/v2/orm/gen/condition"
	"gitlab.yc345.tv/backend/utils/v2/orm/gen/example/postgres/gorm_gen_dao"
	"gitlab.yc345.tv/backend/utils/v2/orm/gen/example/postgres/gorm_gen_model"
	"gitlab.yc345.tv/backend/utils/v2/orm/gen/example/postgres/gorm_gen_repo"
)

// Test_FindOneCacheByID 根据ID查询单条数据
func Test_FindOneCacheByID(t *testing.T) {
	db, err := orm.NewDBWithStruct(&orm.ORMConfig{
		User:            "postgres",
		Password:        "7to12pg12",
		Host:            "10.8.8.110",
		Port:            5433,
		DBname:          "gorm_gen",
		MaxIdleConns:    10,
		MaxOpenConns:    10,
		ConnMaxLifeTime: 3600 * time.Second,
		LogMode:         orm.LogModeInfo,
	})
	if err != nil {
		t.Error("createDBClientError")
		return
	}
	client, _ := redismock.NewClientMock()
	dbCache := goredisdbcache.NewGoRedisDBCache(client)
	ctx := context.Background()
	cfg := config.NewRepoConfig(db.Client, dbCache, encoding.NewMsgPack())
	repo := gorm_gen_repo.NewUserDemoRepo(cfg)
	result, err := repo.FindOneByID(ctx, "0822971b-cbf1-4a44-a2bc-e62cb5a1dd5e")
	if err != nil {
		return
	}
	fmt.Println(result)
	assert.Equal(t, nil, err)
}

// Test_FindMultiByCustom 自定义查询
func Test_FindMultiByCustom(t *testing.T) {
	db, err := orm.NewDBWithStruct(&orm.ORMConfig{
		User:            "postgres",
		Password:        "7to12pg12",
		Host:            "10.8.8.110",
		Port:            5433,
		DBname:          "gorm_gen",
		MaxIdleConns:    10,
		MaxOpenConns:    10,
		ConnMaxLifeTime: 3600 * time.Second,
		LogMode:         orm.LogModeInfo,
	})
	if err != nil {
		t.Error("createDBClientError")
		return
	}
	client, _ := redismock.NewClientMock()
	dbCache := goredisdbcache.NewGoRedisDBCache(client)
	ctx := context.Background()
	cfg := config.NewRepoConfig(db.Client, dbCache, encoding.NewMsgPack())
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
	db, err := orm.NewDBWithStruct(&orm.ORMConfig{
		User:            "postgres",
		Password:        "7to12pg12",
		Host:            "10.8.8.110",
		Port:            5433,
		DBname:          "gorm_gen",
		MaxIdleConns:    10,
		MaxOpenConns:    10,
		ConnMaxLifeTime: 3600 * time.Second,
		LogMode:         orm.LogModeInfo,
	})
	if err != nil {
		t.Error("createDBClientError")
		return
	}
	client, _ := redismock.NewClientMock()
	dbCache := goredisdbcache.NewGoRedisDBCache(client)
	ctx := context.Background()
	cfg := config.NewRepoConfig(db.Client, dbCache, encoding.NewMsgPack())
	adminDemoRepo := gorm_gen_repo.NewAdminDemoRepo(cfg)
	adminLogDemoRepo := gorm_gen_repo.NewAdminLogDemoRepo(cfg)
	err = gorm_gen_dao.Use(db.Client).Transaction(func(tx *gorm_gen_dao.Query) error {
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

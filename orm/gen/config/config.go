package config

import (
	"gitlab.yc345.tv/backend/utils/v2/orm/gen/cache"
	"gitlab.yc345.tv/backend/utils/v2/orm/gen/encoding"
	"gorm.io/gorm"
)

func NewRepoConfig(db *gorm.DB, cache cache.IDBCache, encode encoding.API) *Repo {
	return &Repo{
		DB:       db,
		Cache:    cache,
		Encoding: encode,
	}
}

type Repo struct {
	DB       *gorm.DB
	Cache    cache.IDBCache
	Encoding encoding.API
}

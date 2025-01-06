package gen

import (
	"testing"

	"github.com/fzf-labs/fdatabase/orm"

	"github.com/stretchr/testify/assert"
)

func TestNewGenerationPb(t *testing.T) {
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
	NewGenerationPB(
		db,
		"./example/postgres/pb",
		"api.gorm_gen.v1",
		"api/gorm_gen/v1;v1",
		WithPBOpts(ModelOptionRemoveDefault(), ModelOptionUnderline("ul_")),
	).Do()
	assert.Equal(t, nil, err)
}

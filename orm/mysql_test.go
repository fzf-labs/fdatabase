package orm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDumpMySQL(t *testing.T) {
	config := GormMysqlClientConfig{
		DataSourceName:  "root:123456@tcp(127.0.0.1:3306)/user?charset=utf8mb4&loc=Asia%2FShanghai&parseTime=true",
		MaxIdleConn:     0,
		MaxOpenConn:     0,
		ConnMaxLifeTime: 0,
		ShowLog:         false,
		Tracing:         false,
	}
	gorm, err := NewGormMysqlClient(&config)
	if err != nil {
		return
	}
	DumpMySQL(gorm, "./")
	assert.Equal(t, nil, err)
}

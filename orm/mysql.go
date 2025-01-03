package orm

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/fzf-labs/fdatabase/orm/utils"
	"github.com/fzf-labs/fdatabase/orm/utils/file"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/opentelemetry/tracing"
)

// GormMysqlClientConfig 配置
type GormMysqlClientConfig struct {
	DataSourceName  string        `json:"DataSourceName"`
	MaxIdleConn     int           `json:"MaxIdleConn"`
	MaxOpenConn     int           `json:"MaxOpenConn"`
	ConnMaxLifeTime time.Duration `json:"ConnMaxLifeTime"`
	ShowLog         bool          `json:"ShowLog"`
	Tracing         bool          `json:"Tracing"`
}

// NewGormMysqlClient 初始化gorm mysql 客户端
func NewGormMysqlClient(cfg *GormMysqlClientConfig) (*gorm.DB, error) {
	sqlDB, err := sql.Open("mysql", cfg.DataSourceName)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("open mysql failed! err: %+v", err))
	}
	// set for db connection
	// 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConn)
	// 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConn)
	// 设置连接可以重复使用的最长时间.
	sqlDB.SetConnMaxLifetime(cfg.ConnMaxLifeTime)
	gormConfig := gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	}
	if cfg.ShowLog {
		gormConfig.Logger = logger.Default.LogMode(logger.Info)
	}
	db, err := gorm.Open(mysql.New(mysql.Config{Conn: sqlDB}), &gormConfig)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("mysql database connection failed!  err: %+v", err))
	}
	db.Set("gorm:table_options", "CHARSET=utf8mb4")
	if cfg.Tracing {
		if err := db.Use(tracing.NewPlugin()); err != nil {
			return nil, errors.New(fmt.Sprintf("db use tracing failed!  err: %+v", err))
		}
	}
	return db, nil
}

// DumpMySQL 导出创建语句
func DumpMySQL(db *gorm.DB, outPath string) {
	tables, err := db.Migrator().GetTables()
	if err != nil {
		return
	}
	outPath = filepath.Join(strings.Trim(outPath, "/"), db.Migrator().CurrentDatabase())
	err = os.MkdirAll(outPath, os.ModePerm)
	if err != nil {
		log.Println("DumpMySQL create path err:", err)
		return
	}
	for _, v := range tables {
		result := make(map[string]any)
		err := db.Raw(fmt.Sprintf("SHOW CREATE TABLE `%s`.`%s`", db.Migrator().CurrentDatabase(), v)).Scan(result).Error
		if err != nil {
			log.Println("DumpMySQL sql err:", err)
			return
		}
		tableContent := utils.ConvToString(result["Create Table"])
		if tableContent != "" {
			err := file.WriteContentCover(filepath.Join(outPath, fmt.Sprintf("%s.sql", v)), tableContent)
			if err != nil {
				log.Println("DumpMySQL file write err:", err)
				return
			}
		}
	}
}

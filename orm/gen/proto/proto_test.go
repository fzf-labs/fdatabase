package proto

import (
	"testing"

	"github.com/fzf-labs/fdatabase/orm"

	"gorm.io/gorm"
)

func newDB() *gorm.DB {
	db, err := orm.NewGormPostgresClient(&orm.GormPostgresClientConfig{
		DataSourceName:  "host=0.0.0.0 port=5432 user=postgres password=123456 dbname=gorm_gen sslmode=disable TimeZone=Asia/Shanghai",
		MaxIdleConn:     0,
		MaxOpenConn:     0,
		ConnMaxLifeTime: 0,
		ShowLog:         false,
		Tracing:         false,
	})
	if err != nil {
		panic(err)
	}
	return db
}

func TestGenerationPB(t *testing.T) {
	db := newDB()
	type args struct {
		db               *gorm.DB
		outPutPath       string
		packageStr       string
		goPackageStr     string
		table            string
		columnNameToName map[string]string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				db:           db,
				outPutPath:   "../example/postgres/pb",
				packageStr:   "api.gorm_gen.v1",
				goPackageStr: "api/gorm_gen/v1;v1",
				table:        "admin_log_demo",
				columnNameToName: map[string]string{
					"id":         "ID",
					"admin_id":   "adminID",
					"ip":         "IP",
					"uri":        "URI",
					"useragent":  "Useragent",
					"header":     "Header",
					"req":        "Req",
					"resp":       "Resp",
					"created_at": "CreatedAt",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GenerationPB(tt.args.db, tt.args.outPutPath, tt.args.packageStr, tt.args.goPackageStr, tt.args.table, tt.args.columnNameToName); (err != nil) != tt.wantErr {
				t.Errorf("GenerationPB() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

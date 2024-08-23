package proto

import (
	"testing"

	"gitlab.yc345.tv/backend/utils/v2/orm"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	newDBWithStruct, err := orm.NewDBWithStruct(&orm.ORMConfig{
		User:     "postgres",
		Password: "7to12pg12",
		Host:     "10.8.8.110",
		Port:     5433,
		DBname:   "gorm_gen",
	})
	if err != nil {
		return
	}
	db = newDBWithStruct.Client
}
func TestGenerationPB(t *testing.T) {
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

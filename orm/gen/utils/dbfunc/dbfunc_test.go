package dbfunc

import (
	"reflect"
	"testing"

	"gitlab.yc345.tv/backend/utils/v2/orm"
	"gorm.io/gorm"
)

var db *gorm.DB
var dbErr *gorm.DB

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
	newDBWithStructErr, err := orm.NewDBWithStruct(&orm.ORMConfig{
		User:     "postgres",
		Password: "7to12pg12",
		Host:     "10.8.8.110",
		Port:     5433,
		DBname:   "gorm_gen",
	})
	if err != nil {
		return
	}
	dbErr = newDBWithStructErr.Client
	newDBWithStructErr.Close()
}

func TestSortIndexColumns(t *testing.T) {
	type args struct {
		db    *gorm.DB
		table string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string][]string
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				db:    db,
				table: "admin_demo",
			},
			want:    map[string][]string{"admin_demo_pkey": {"id"}, "admin_demo_username_idx": {"username"}},
			wantErr: false,
		},
		{
			name: "test-err",
			args: args{
				db:    dbErr,
				table: "admin_demo",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SortIndexColumns(tt.args.db, tt.args.table)
			if (err != nil) != tt.wantErr {
				t.Errorf("SortIndexColumns() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortIndexColumns() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPartitionTableName(t *testing.T) {
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name     string
		args     args
		wantResp []string
		wantErr  bool
	}{
		{
			name: "test1",
			args: args{
				db: db,
			},
			wantResp: []string{"partition_table"},
			wantErr:  false,
		},
		{
			name: "test-err",
			args: args{
				db: dbErr,
			},
			wantResp: nil,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := GetPartitionTableName(tt.args.db)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPartitionTableName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("GetPartitionTableName() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestGetPartitionChildTableForTable(t *testing.T) {
	type args struct {
		db        *gorm.DB
		tableName string
	}
	tests := []struct {
		name     string
		args     args
		wantResp []string
		wantErr  bool
	}{
		{
			name: "test1",
			args: args{
				db:        db,
				tableName: "partition_table",
			},
			wantResp: []string{"partition_table_2022", "partition_table_2023"},
			wantErr:  false,
		},
		{
			name: "test-err",
			args: args{
				db:        dbErr,
				tableName: "partition_table",
			},
			wantResp: nil,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := GetPartitionChildTableForTable(tt.args.db, tt.args.tableName)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPartitionChildTableForTable() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("GetPartitionChildTableForTable() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestGetPartitionChildTable(t *testing.T) {
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name     string
		args     args
		wantResp []string
		wantErr  bool
	}{
		{
			name: "test1",
			args: args{
				db: db,
			},
			wantResp: []string{"partition_table_2022", "partition_table_2023"},
			wantErr:  false,
		},
		{
			name: "test-err",
			args: args{
				db: dbErr,
			},
			wantResp: nil,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := GetPartitionChildTable(tt.args.db)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPartitionChildTable() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("GetPartitionChildTable() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

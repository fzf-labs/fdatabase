package sqldump

import (
	"reflect"
	"testing"
)

func TestNewSQLDump(t *testing.T) {
	type args struct {
		db           string
		dsn          string
		outPutPath   string
		targetTables string
		fileCover    bool
	}
	tests := []struct {
		name string
		args args
		want *SQLDump
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSQLDump(tt.args.db, tt.args.dsn, tt.args.outPutPath, tt.args.targetTables, tt.args.fileCover); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSQLDump() = %v, want %v", got, tt.want)
			}
		})
	}
}

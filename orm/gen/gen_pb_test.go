package gen

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewGenerationPB(t *testing.T) {
	type args struct {
		db           *gorm.DB
		outPutPath   string
		packageStr   string
		goPackageStr string
		opts         []OptionPB
	}
	tests := []struct {
		name string
		args args
		want *GenerationPb
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewGenerationPB(tt.args.db, tt.args.outPutPath, tt.args.packageStr, tt.args.goPackageStr, tt.args.opts...), "NewGenerationPB(%v, %v, %v, %v, %v)", tt.args.db, tt.args.outPutPath, tt.args.packageStr, tt.args.goPackageStr, tt.args.opts...)
		})
	}
}

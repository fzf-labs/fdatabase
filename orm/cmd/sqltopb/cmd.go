package sqltopb

import (
	"github.com/fzf-labs/fdatabase/orm/gen"
	"github.com/fzf-labs/fdatabase/orm/utils/dbfunc"
	"github.com/spf13/cobra"
)

var CmdSQLToPb = &cobra.Command{
	Use:   "sqltopb",
	Short: "sql generate proto file",
	Long:  "sql generate proto file",
	Run:   Run,
}

var (
	db          string // 数据库类型 mysql postgres
	dsn         string // 数据库连接
	pbPackage   string // proto 包名
	pbGoPackage string // proto go包名
	outPutPath  string // 输出路径
)

//nolint:gochecknoinits
func init() {
	CmdSQLToPb.Flags().StringVarP(&dsn, "db", "", "", "db：mysql postgres")
	CmdSQLToPb.Flags().StringVarP(&dsn, "dsn", "", "", "dsn")
	CmdSQLToPb.Flags().StringVarP(&pbPackage, "pbPackage", "", "kratos_demo.v1", "pbPackage")
	CmdSQLToPb.Flags().StringVarP(&pbGoPackage, "pbGoPackage", "", "gitlab.yc345.tv/backend/kratos-demo/api/kratos_demo/v1;v1", "pbGoPackage")
	CmdSQLToPb.Flags().StringVarP(&outPutPath, "outPutPath", "", "./api/kratos_demo/v1", "outPutPath")
}

func Run(_ *cobra.Command, _ []string) {
	gen.NewGenerationPB(dbfunc.NewSimpleDB(db, dsn), outPutPath, pbPackage, pbGoPackage, gen.WithPBOpts(gen.ModelOptionRemoveDefault(), gen.ModelOptionUnderline("UL"))).Do()
}

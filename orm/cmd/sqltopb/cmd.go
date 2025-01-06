package sqltopb

import (
	"github.com/fzf-labs/fctl/utils"
	"github.com/spf13/cobra"
	"gitlab.yc345.tv/backend/utils/v2/orm/gen"
)

var CmdSQLToPb = &cobra.Command{
	Use:   "sqltopb",
	Short: "sql to pb",
	Long:  "sql to pb.",
	Run:   Run,
}

var (
	dsn         string // 数据库连接
	pbPackage   string // proto 包名
	pbGoPackage string // proto go包名
	outPutPath  string // 输出路径
)

//nolint:gochecknoinits
func init() {
	CmdSQLToPb.Flags().StringVarP(&dsn, "dsn", "d", "", "dsn")
	CmdSQLToPb.Flags().StringVarP(&pbPackage, "pbPackage", "p", "kratos_demo.v1", "pbPackage")
	CmdSQLToPb.Flags().StringVarP(&pbGoPackage, "pbGoPackage", "g", "gitlab.yc345.tv/backend/kratos-demo/api/kratos_demo/v1;v1", "pbGoPackage")
	CmdSQLToPb.Flags().StringVarP(&outPutPath, "outPutPath", "o", "./api/kratos_demo/v1", "outPutPath")
}

func Run(_ *cobra.Command, _ []string) {
	gen.NewGenerationPB(utils.NewDB(dsn), outPutPath, pbPackage, pbGoPackage, gen.WithPBOpts(gen.ModelOptionRemoveDefault(), gen.ModelOptionUnderline("UL"))).Do()
}

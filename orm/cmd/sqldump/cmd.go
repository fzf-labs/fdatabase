package sqldump

import (
	"github.com/spf13/cobra"
)

var CmdSQLDump = &cobra.Command{
	Use:   "sqldump",
	Short: "sql dump",
	Long:  "sql dump.",
	Run:   Run,
}

var (
	dsn           string // 数据库连接
	outPutPath    string // 输出路径
	targetTables  string // 指定表
	fileOverwrite bool   // 是否覆盖
)

//nolint:gochecknoinits
func init() {
	CmdSQLDump.Flags().StringVarP(&dsn, "dsn", "d", "", "dsn")
	CmdSQLDump.Flags().StringVarP(&outPutPath, "outPutPath", "o", "./doc/sql", "outPutPath")
	CmdSQLDump.Flags().StringVarP(&targetTables, "tables", "t", "", "tables")
	CmdSQLDump.Flags().BoolVarP(&fileOverwrite, "fileOverwrite", "f", false, "file overwrite")
}

func Run(_ *cobra.Command, _ []string) {
	NewSQLDump(dsn, outPutPath, targetTables, fileOverwrite).Run()
}

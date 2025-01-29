package main

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/fzf-labs/fdatabase/cmd/fdatabase/ormgen"
	"github.com/fzf-labs/fdatabase/cmd/fdatabase/sqldump"
	"github.com/fzf-labs/fdatabase/cmd/fdatabase/sqltopb"
)

var rootCmd = &cobra.Command{
	Use:     "fdatabase",
	Short:   "fdatabase: an db toolkit",
	Long:    `fdatabase: an db toolkit`,
	Version: "v0.0.1",
}

func init() {
	rootCmd.AddCommand(ormgen.CmdOrmGen)
	rootCmd.AddCommand(sqldump.CmdSQLDump)
	rootCmd.AddCommand(sqltopb.CmdSQLToPb)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

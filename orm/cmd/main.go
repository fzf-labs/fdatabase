package main

import (
	"log"

	"github.com/fzf-labs/fdatabase/orm/cmd/ormgen"
	"github.com/fzf-labs/fdatabase/orm/cmd/sqldump"
	"github.com/fzf-labs/fdatabase/orm/cmd/sqltopb"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "fdatabase",
	Short:   "fdatabase: an db toolkit",
	Long:    `fdatabase: An toolkit.`,
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

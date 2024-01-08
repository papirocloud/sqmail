package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sqmail",
	Short: "SQmaiL allows you to query your IMAP mailbox using SQL",
}

var format string

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&format, "format", "f", "table", "Output format (table, csv, json, html, markdown)")
}

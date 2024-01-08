package cmd

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/papirocloud/sqmail/sql"
	"github.com/spf13/cobra"
)

var getFieldsFormat string

// getFieldsCmd represents the getFields command
var getFieldsCmd = &cobra.Command{
	Use:   "getFields",
	Short: "Get the list of fields",
	Run: func(cmd *cobra.Command, args []string) {
		fields := sql.ListFields()

		t := table.NewWriter()
		t.SetOutputMirror(cmd.OutOrStdout())

		t.AppendHeader(table.Row{"Name", "Aliases", "Operators", "Selectable", "Searchable"})

		for _, field := range fields {
			t.AppendRow(table.Row{field.Name, field.Aliases, field.AllowedOperators, field.Selectable, field.Searchable})
		}

		switch format {
		case "table":
			t.Render()
		case "markdown":
			t.RenderMarkdown()
		}
	},
}

func init() {
	rootCmd.AddCommand(getFieldsCmd)
}

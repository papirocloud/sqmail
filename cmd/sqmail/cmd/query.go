package cmd

import (
	"io"
	"os"

	"github.com/papirocloud/sqmail/email"
	sqmailImap "github.com/papirocloud/sqmail/imap"
	"github.com/papirocloud/sqmail/sql"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

var (
	host     string
	port     int64
	tls      bool
	username string
	password string
	output   string
	query    string
	silent   bool
)

// queryCmd represents the query command
var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "Query your IMAP mailbox using SQL",
	Run: func(cmd *cobra.Command, args []string) {
		if silent {
			zerolog.SetGlobalLevel(zerolog.Disabled)
		}

		c, err := sqmailImap.Connect(host, port, tls)
		if err != nil {
			panic(err)
		}
		defer func() { _ = c.Close() }()

		if err := sqmailImap.Login(c, username, password); err != nil {
			panic(err)
		}

		fields := sql.GetFieldsFromQuery(query)

		msgCh := make(chan *email.Message)
		mapsCh := make(chan map[string]interface{})
		outputCh := make(chan struct{})

		var w io.Writer
		if output == "" {
			w = os.Stdout
		} else {
			f, err := os.Create(output)
			if err != nil {
				panic(err)
			}
			defer func() { _ = f.Close() }()
			w = f
		}

		go writeOutput(w, format, fields, mapsCh, outputCh)

		go func() {
			if err := sql.Query(c, query, msgCh); err != nil {
				panic(err)
			}
		}()

		for msg := range msgCh {
			handleMessage(fields, mapsCh, msg)
		}

		close(mapsCh)
		<-outputCh
	},
}

func init() {
	rootCmd.AddCommand(queryCmd)

	queryCmd.Flags().StringVarP(&host, "host", "H", "", "IMAP server hostname")
	_ = queryCmd.MarkFlagRequired("host")

	queryCmd.Flags().Int64VarP(&port, "port", "p", 993, "IMAP server port")

	queryCmd.Flags().BoolVarP(&tls, "tls", "t", true, "Use TLS")

	queryCmd.Flags().StringVarP(&username, "username", "u", "", "IMAP username")
	_ = queryCmd.MarkFlagRequired("username")

	queryCmd.Flags().StringVarP(&password, "password", "P", "", "IMAP password")
	_ = queryCmd.MarkFlagRequired("password")

	queryCmd.Flags().StringVarP(&output, "output", "o", "", "Output file (default: stdout)")

	queryCmd.Flags().StringVarP(&query, "query", "q", "", "SQL query")
	_ = queryCmd.MarkFlagRequired("query")

	queryCmd.Flags().BoolVarP(&silent, "silent", "s", false, "Silent mode (no logging)")
}

func handleMessage(fields []*sql.Field, mapsCh chan<- map[string]interface{}, msg *email.Message) {
	mfields := sql.GetFieldsFromMessage(msg, fields)
	mapsCh <- mfields
}

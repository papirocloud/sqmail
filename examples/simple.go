//go:build ignore

package main

import (
	"fmt"

	sqmailImap "github.com/papirocloud/sqmail/imap"
	"github.com/papirocloud/sqmail/sql"
)

var host = "imap.fastmail.com"
var port = 993
var tls = true
var user = "user@domain.com"
var password = "12d1das98df8as9f7d"
var query = "SELECT from_, subject FROM INBOX WHERE date > '2024-01-01' AND subject LIKE 'FooBar'"

func main() {
	c, err := sqmailImap.Connect(host, port, tls)
	if err != nil {
		panic(err)
	}

	if err := sqmailImap.Login(c, user, password); err != nil {
		panic(err)
	}

	fields := sql.GetFieldsFromQuery(query)

	msgs, err := sql.Query(c, query)
	if err != nil {
		panic(err)
	}

	for _, msg := range msgs {
		mfields := sql.GetFieldsFromMessage(msg, fields)
		fmt.Println(mfields)
	}
}

# SQmaiL

SQmaiL allows you to query your IMAP email server using SQL.

## Usage

```
Usage:
  sqmail query [flags]

Flags:
  -f, --format string     Output format (table, csv, json, html, markdown) (default "table")
  -h, --help              help for query
  -H, --host string       IMAP server hostname
  -o, --output string     Output file (default: stdout)
  -P, --password string   IMAP password
  -p, --port int          IMAP server port (default 993)
  -q, --query string      SQL query
  -s, --silent            Silent mode (no logging)
  -t, --tls               Use TLS (default true)
  -u, --username string   IMAP username
```

### Examples

#### Get all emails, in any folder, containing "GitHub" in the subject:
```bash
sqmail query -f table -H 'imap.fastmail.com' -u 'user@domain.com' -P "pAsSwOrD1!" -q "SELECT uid,from_,to_,subject FROM emails WHERE subject LIKE 'GitHub' AND mailbox = 'ANYWHERE'"
```

```
+-------+-----------------------+---------------------------------------+---------------------------------------------------------------------------------------------------------------+
| UID   | FROM                  | TO                                    | SUBJECT                                                                                                       |
+-------+-----------------------+---------------------------------------+---------------------------------------------------------------------------------------------------------------+
| 15941 | [support@github.com]  | [user@email.com]                     | [GitHub] Two-factor authentication enabled                                                                     |
```

#### Get up to 10 emails from the "INBOX" folder in JSON format
```bash
sqmail query -f json -s -H 'imap.fastmail.com' -u 'user@domain.com' -P "pAsSwOrD1!" -q "SELECT uid,from_,to_,subject FROM emails WHERE mailbox = 'INBOX' LIMIT 10"
```

#### Get all emails from the "INBOX" folder with an arbitrary header named "X-My-Header" containing "foo" in the value
```bash
sqmail query -f json -s -H 'imap.fastmail.com' -u 'user@domain.com' -P "pAsSwOrD1!" -q "SELECT uid,from_,to_,subject FROM emails WHERE mailbox = 'INBOX' AND headers = ('X-My-Header', 'foo')"
```

#### Get all emails from the "Special Emails" folder containing "GitHub" in the subject, and save the results to a CSV file
```bash
sqmail query -f csv -s -o emails.csv -H 'imap.fastmail.com' -u 'user@domain.com' -P "pAsSwOrD1!" -q "SELECT uid,from_,to_,subject FROM emails WHERE mailbox = 'Special Emails' subject LIKE 'GitHub'"
```

## Limitations

### Equals vs Like
Due to the way IMAP search works, both the `=` and `LIKE` operators are identical.

They are case-insensitive and perform a substring match (aka "contains").


### Output Formats
When using `csv` or `json` as the output format, messages are streamed as they are received from the server.

When using `table`, `html`, or `markdown` as the output format, messages are buffered in memory and then pretty-printed.
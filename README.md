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

#### Getting a list of all fields available for querying

```bash
sqmail getFields
```

```markdown
| Name                             | Aliases               | Operators     | Selectable | Searchable |
|----------------------------------|-----------------------|---------------|------------|------------|
| raw                              | []                    | []            | true       | false      |
| from                             | [from_ fromAddresses] | [= LIKE]      | true       | true       |
| bcc                              | [bcc_ bccAddresses]   | [= LIKE]      | true       | true       |
| has_events                       | []                    | [=]           | true       | true       |
| events                           | []                    | []            | true       | false      |
| cc                               | [cc_ ccAddresses]     | [= LIKE]      | true       | true       |
| attachments                      | []                    | []            | true       | false      |
| to                               | [to_ toAddresses]     | [= LIKE]      | true       | true       |
| date                             | []                    | [= < > <= >=] | true       | true       |
| html                             | []                    | [LIKE]        | true       | true       |
| subject                          | []                    | [= LIKE]      | true       | true       |
| embedded                         | []                    | []            | true       | false      |
| has_attachment_with_content_type | []                    | [=]           | false      | true       |
| flags                            | []                    | [= !=]        | true       | true       |
| has_attachments                  | []                    | [=]           | true       | true       |
| seqnum                           | []                    | [=]           | false      | true       |
| serverdate                       | []                    | [= < > <= >=] | true       | true       |
| raw_invites                      | []                    | []            | true       | false      |
| mailbox                          | []                    | [=]           | false      | true       |
| uid                              | []                    | [=]           | true       | true       |
| size                             | []                    | [= < > <= >=] | true       | true       |
| headers                          | []                    | [= LIKE]      | true       | true       |
| has_embeds                       | []                    | [=]           | true       | true       |
| text                             | []                    | [LIKE]        | true       | true       |
```

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

_Note: For some reason, some IMAP servers hang when you try to query for an arbitrary header._

```bash
sqmail query -f json -s -H 'imap.fastmail.com' -u 'user@domain.com' -P "pAsSwOrD1!" -q "SELECT uid,from_,to_,subject FROM emails WHERE mailbox = 'INBOX' AND headers = ('X-My-Header', 'foo')"
```

#### Get all emails from the "Special Emails" folder containing "GitHub" in the subject, and save the results to a CSV file

```bash
sqmail query -f csv -s -o emails.csv -H 'imap.fastmail.com' -u 'user@domain.com' -P "pAsSwOrD1!" -q "SELECT uid,from_,to_,subject FROM emails WHERE mailbox = 'Special Emails' AND subject LIKE 'GitHub'"
```

## Limitations

### Equals vs Like

Due to the way IMAP search works, both the `=` and `LIKE` operators are identical.

They are case-insensitive and perform a substring match (aka "contains").

### Output Formats

When using `csv` or `json` as the output format, messages are streamed as they are received from the
server.

When using `table`, `html`, or `markdown` as the output format, messages are buffered in memory and
then pretty-printed.

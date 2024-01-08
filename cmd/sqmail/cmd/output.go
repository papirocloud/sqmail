package cmd

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/papirocloud/sqmail/sql"
)

const (
	jsonFormat     = "json"
	csvFormat      = "csv"
	tableFormat    = "table"
	htmlFormat     = "html"
	markdownFormat = "markdown"
)

func renderTable(w io.Writer, format string, fields []*sql.Field, mapsCh <-chan map[string]interface{}) {
	var maps []map[string]interface{}

	for m := range mapsCh {
		maps = append(maps, m)
	}

	t := table.NewWriter()
	t.SetOutputMirror(w)

	var header table.Row

	for _, f := range fields {
		header = append(header, f.Name)
	}
	t.AppendHeader(header)

	for _, m := range maps {
		var row table.Row
		for _, h := range header {
			row = append(row, fmt.Sprintf("%v", m[h.(string)]))
		}
		t.AppendRow(row)
	}

	switch format {
	case tableFormat:
		t.Render()
	case htmlFormat:
		t.RenderHTML()
	case markdownFormat:
		t.RenderMarkdown()
	}
}

func writeCsv(w io.Writer, fields []*sql.Field, mapsCh <-chan map[string]interface{}) {
	csvWriter := csv.NewWriter(w)
	defer csvWriter.Flush()

	var headers []string

	for _, f := range fields {
		headers = append(headers, f.Name)
	}
	if err := csvWriter.Write(headers); err != nil {
		panic(err)
	}

	for m := range mapsCh {
		var row []string
		for _, h := range headers {
			row = append(row, fmt.Sprintf("%v", m[h]))
		}
		if err := csvWriter.Write(row); err != nil {
			panic(err)
		}
		csvWriter.Flush()
	}

	if err := csvWriter.Error(); err != nil {
		panic(err)
	}
}

func writeJson(w io.Writer, mapsCh <-chan map[string]interface{}) {
	encoder := json.NewEncoder(w)

	for m := range mapsCh {
		if err := encoder.Encode(m); err != nil {
			panic(err)
		}
	}
}

func writeOutput(w io.Writer, format string, fields []*sql.Field, mapsCh <-chan map[string]interface{}, outputCh chan<- struct{}) {
	switch format {
	case tableFormat, htmlFormat, markdownFormat:
		renderTable(w, format, fields, mapsCh)
	case csvFormat:
		writeCsv(w, fields, mapsCh)
	case jsonFormat:
		writeJson(w, mapsCh)
	}

	outputCh <- struct{}{}
}

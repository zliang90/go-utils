package table

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
)

type TableData struct {
	Header []string
	Data   [][]string
	Footer []string
}

func (t *TableData) GetCols() int {
	return len(t.Header)
}

func (t *TableData) PrintTable(displayFooter bool) error {

	// write to stdout
	table := tablewriter.NewWriter(os.Stdout)

	// setting table style
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetAutoMergeCells(false)

	// table header
	table.SetHeader(t.Header)

	// fill row data
	table.AppendBulk(t.Data)

	// footer
	if displayFooter {
		footer := make([]string, t.GetCols())
		footer[t.GetCols()-1] = fmt.Sprintf(" COUNT: %d ", len(t.Data))
		t.Footer = footer

		// table footer
		table.SetFooter(t.Footer)
	}

	table.Render()
	return nil
}

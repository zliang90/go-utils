package table

import "testing"

func TestTableData_PrintTable(t *testing.T) {
	table := TableData{
		Header: []string{"id", "name", "sex", "age"},
		Data: [][]string{
			{"1", "tom", "male", "28"},
			{"2", "jack", "male", "30"},
			{"3", "mary", "female", "23"},
		},
	}

	table.PrintTable(true)
}

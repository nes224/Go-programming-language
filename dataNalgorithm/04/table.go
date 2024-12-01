package main

import "fmt"

type Table struct {
	Rows        []Row
	Name        string
	ColumnNames []string
}

type Row struct {
	Columns []Column
	Id      int
}

type Column struct {
	Id    int
	Value string
}

func printTable(table Table) {
	var rows []Row = table.Rows
	fmt.Println(table.Name)
	for _, rows := range rows {
		var columns []Column = rows.Columns
		for i, column := range columns {
			fmt.Println(table.ColumnNames[i], column.Id, column.Value)
		}
	}
}

func main() {
	var table Table = Table{}
	table.Name = "Customer"
	table.ColumnNames = []string{"Id", "Name", "SSN"}
	var rows []Row = make([]Row, 2)
	rows[0] = Row{}
	var columns1 []Column = make([]Column, 3)
	columns1[0] = Column{1, "323"}
	columns1[1] = Column{1, "John Smith"}
	columns1[2] = Column{1, "3453223"}
	rows[0].Columns = columns1
	rows[1] = Row{}
	var columns2 []Column = make([]Column, 3)
	columns2[0] = Column{2, "223"}
	columns2[1] = Column{2, "Curran Smith"}
	columns2[1] = Column{2, "3223211"}
	rows[1].Columns = columns2
	table.Rows = rows
	fmt.Println(table)
	printTable(table)
}

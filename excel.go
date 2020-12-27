package main

import (
	"log"

	"github.com/tealeg/xlsx/v3"
)

// ReadExcel reads an excel file and parses it into grid cells
func ReadExcel(path string) (err error) {
	wb, err := xlsx.OpenFile(path)
	if err != nil {
		return
	}

	sh := wb.Sheets[0]

	log.Println(sh.MaxRow)
	log.Println(sh.MaxCol)

	grid.rowTotal = sh.MaxRow
	grid.colTotal = sh.MaxCol

	sh.ForEachRow(rowVisitor)

	return
}

func rowVisitor(row *xlsx.Row) (err error) {
	return row.ForEachCell(cellVisitor)
}

func cellVisitor(cell *xlsx.Cell) (err error) {
	col, row := cell.GetCoordinates()

	value, err := cell.FormattedValue()
	if err != nil {
		log.Println(err.Error())
		return
	}

	grid.SetCell(row, col, value)

	return
}

package main

import (
	"encoding/csv"
	"os"
)

var inFile *os.File

// ReadCSV reads a csv file and parses it into grid cells
func ReadCSV() (err error) {
	if len(os.Args) < 2 {
		return
	}

	path := os.Args[1]

	inFile, err = os.Open(path)
	if err != nil {
		return
	}

	reader := csv.NewReader(inFile)
	records, err := reader.ReadAll()
	if err != nil {
		return
	}

	grid.rowTotal = len(records)
	grid.colTotal = len(records[0])

	for row, record := range records {
		for col, value := range record {
			grid.SetCell(row, col, value)
		}
	}

	return
}
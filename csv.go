package main

import (
	"encoding/csv"
	"os"
)

// ReadCSV reads a csv file and parses it into grid cells
func ReadCSV() (err error) {
	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		return
	}

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return
	}

	for row, record := range records {
		for col, value := range record {
			grid.SetCell(row, col, value)
		}
	}

	return
}

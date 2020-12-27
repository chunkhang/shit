package file

import (
	"encoding/csv"
	"os"
)

// CSVReader is a reader for csv file
type CSVReader struct{}

func (r *CSVReader) Read(path string) (file *File, err error) {
	f, err := os.Open(path)
	if err != nil {
		return
	}

	c := csv.NewReader(f)
	records, err := c.ReadAll()
	if err != nil {
		return
	}

	grid := &Grid{
		cells: map[int]map[int]*Cell{},
	}
	grid.RowTotal = len(records)
	grid.ColTotal = len(records[0])
	for row, record := range records {
		for col, value := range record {
			grid.SetCell(row, col, value)
		}
	}

	sheet := &Sheet{
		Grid: grid,
	}

	file = &File{
		Name:   path,
		Sheets: []*Sheet{sheet},
	}

	return
}

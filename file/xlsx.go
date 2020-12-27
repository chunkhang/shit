package file

import (
	"log"

	"github.com/tealeg/xlsx/v3"
)

// XLSXReader is a reader for xlsx file
type XLSXReader struct{}

func (r *XLSXReader) Read(path string) (file *File, err error) {
	wb, err := xlsx.OpenFile(path)
	if err != nil {
		return
	}

	sheets := []*Sheet{}
	for _, sh := range wb.Sheets {
		grid := &Grid{
			cells: map[int]map[int]*Cell{},
		}
		grid.RowTotal = sh.MaxRow
		grid.ColTotal = sh.MaxCol

		for row := 0; row < grid.RowTotal; row++ {
			for col := 0; col < grid.ColTotal; col++ {
				c, err := sh.Cell(row, col)
				if err != nil {
					log.Println(err)
					continue
				}
				value, err := c.FormattedValue()
				if err != nil {
					log.Println(err)
					continue
				}
				grid.SetCell(row, col, value)
			}
		}

		sheet := &Sheet{
			Name: sh.Name,
			Grid: grid,
		}
		sheets = append(sheets, sheet)
	}

	file = &File{
		Name:   path,
		Sheets: sheets,
	}

	return
}

package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	f "github.com/chunkhang/shit/file"
)

var (
	version string
	mode    = "dev"
)

var (
	file  *f.File
	sheet *f.Sheet
	grid  *f.Grid
)

func main() {
	// Usage information
	if len(os.Args) < 2 {
		fmt.Println("usage: shit [filename]")
		os.Exit(0)
	}

	arg := os.Args[1]

	// Handle flags
	if strings.HasPrefix(arg, "-") {
		switch arg {
		case "-v", "--version":
			fmt.Printf(version)
			if mode == "dev" {
				fmt.Printf(" (%s)", mode)
			}
			fmt.Println()
			os.Exit(0)
		default:
			fmt.Printf("invalid option \"%s\"\n", arg)
			os.Exit(1)
		}
	}

	err := StartLog()
	checkErr(err)
	defer StopLog()

	// List of file readers
	readers := []f.Reader{
		&f.CSVReader{},
		&f.XLSXReader{},
	}

	// Try to read with all file readers
	// Consider failed if no file reader works
	for _, reader := range readers {
		file, err = reader.Read(arg)
		if err == nil {
			break
		}
	}
	if err != nil {
		err = errors.New("File cannot be read")
	}
	checkErr(err)

	sheet = file.Sheets[0]
	grid = sheet.Grid

	// Set first cell as cursor
	cell := grid.GetCell(0, 0)
	cursor = &Cursor{cell}

	err = StartScreen()
	checkErr(err)
	defer StopScreen()

	StartEngine()
}

func checkErr(err error) {
	if err != nil {
		fmt.Printf("%v\n", err.Error())
		os.Exit(1)
	}
}

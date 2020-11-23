package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	version string
	mode    = "dev"
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
			fmt.Printf("invalid option \"%s\"", arg)
			os.Exit(1)
		}
	}

	err := StartLog()
	checkErr(err)
	defer StopLog()

	log.Println("start")

	err = ReadCSV(arg)
	checkErr(err)

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

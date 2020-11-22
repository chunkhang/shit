package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	err := StartLog()
	checkErr(err)
	defer StopLog()

	err = ReadCSV()
	checkErr(err)

	err = StartScreen()
	checkErr(err)
	defer StopScreen()

	StartEngine()
}

func checkErr(err error) {
	if err != nil {
		log.Printf("%v", err.Error())
		fmt.Printf("%v\n", err.Error())
		os.Exit(1)
	}
}

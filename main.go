package main

import (
	"log"
)

func main() {
	err := StartLog()
	checkErr(err)
	defer StopLog()

	err = StartScreen()
	checkErr(err)
	defer StopScreen()

	StartEngine()
}

func checkErr(err error) {
	if err != nil {
		log.Fatalf("%v", err.Error())
	}
}

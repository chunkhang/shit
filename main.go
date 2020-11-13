package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("xr.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("%+v", err)
	}
	log.SetOutput(file)
	defer file.Close()

	StartScreen()

	StartEngine()

	StopScreen()
}

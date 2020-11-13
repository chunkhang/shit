package main

import (
	"log"
	"os"
)

const fileName = "xr.log"

var file *os.File

// StartLog opens the log file
func StartLog() (err error) {
	file, err = os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return
	}

	log.SetOutput(file)

	return
}

// StopLog closes the log file
func StopLog() (err error) {
	return file.Close()
}
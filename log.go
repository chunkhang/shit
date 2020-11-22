package main

import (
	"log"
	"os"
)

const fileName = "xr.log"

var logFile *os.File

// StartLog opens the log file
func StartLog() (err error) {
	logFile, err = os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return
	}

	log.SetOutput(logFile)

	return
}

// StopLog closes the log file
func StopLog() (err error) {
	return logFile.Close()
}

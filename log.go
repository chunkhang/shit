package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var logFile *os.File

// StartLog opens the log file
func StartLog() (err error) {
	// Only log in development mode
	if mode != "dev" {
		log.SetOutput(ioutil.Discard)
		return
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return
	}

	path := filepath.Join(home, ".shit.log")
	logFile, err = os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return
	}

	log.SetOutput(logFile)

	return
}

// StopLog closes the log file
func StopLog() (err error) {
	if logFile == nil {
		return nil
	}
	return logFile.Close()
}

package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

const (
	InfoMode = iota
	WarningMode
	ErrorMode
)

const (
	blueFormat   = "\033[1;34m[%d] %s:%d %s\n\033[0m"
	yellowFormat = "\033[1;33m[%d] %s:%d %s\n\033[0m"
	redFormat    = "\033[1;31m[%d] %s:%d %s\n\033[0m"
)

func init() {

	// Remove timestamp from logs
	log.SetFlags(0)
}

func CheckError(err error, mode int) {

	if err != nil {
		// Get file and line of the error
		_, file, line, _ := runtime.Caller(1)
		file = filepath.Base(file)

		switch mode {
		case InfoMode:
			fmt.Printf(blueFormat, time.Now().Unix(), file, line, err)
		case WarningMode:
			fmt.Printf(yellowFormat, time.Now().Unix(), file, line, err)
		case ErrorMode:
			fmt.Printf(redFormat, time.Now().Unix(), file, line, err)
			os.Exit(1)
		}
	}
}

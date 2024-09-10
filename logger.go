package main

import (
	"fmt"
	"log"
	"os"
)

// LogLevel constants
const (
	DebugLog	   string = "DEBUG"
	InformationLog string = "INFO"
	WarningLog	   string = "WARN"
	ErrorLog	   string = "ERROR"
)

var (
	logLevel			string
	StdoutInfoLogger	*log.Logger
	StdoutWarningLogger *log.Logger
	StdoutErrorLogger   *log.Logger
	StdoutDebugLogger   *log.Logger
	logLevelPriority	map[string]int
)

func init() {
	// Initialize loggers
	StdoutInfoLogger = log.New(os.Stdout, "INFO ", log.Ldate|log.Ltime)
	StdoutWarningLogger = log.New(os.Stdout, "WARNING ", log.Ldate|log.Ltime)
	StdoutErrorLogger = log.New(os.Stdout, "ERROR ", log.Ldate|log.Ltime)
	StdoutDebugLogger = log.New(os.Stdout, "DEBUG ", log.Ldate|log.Ltime)
	
	// Initialize log level priorities
	logLevelPriority = map[string]int{
		DebugLog:		1,
		InformationLog: 2,
		WarningLog:		3,
		ErrorLog:		4,
	}
}

func DDNSLogger(logType, host, domain, message string) {
	// Ensure logType maps to a valid priority
	if logLevelPriority[logType] >= logLevelPriority[logLevel] {
		logMessage := fmt.Sprintf("%s.%s %s", host, domain, message)
		switch logType {
		case DebugLog:
			StdoutDebugLogger.Println(logMessage)
		case InformationLog:
			StdoutInfoLogger.Println(logMessage)
		case WarningLog:
			StdoutWarningLogger.Println(logMessage)
		case ErrorLog:
			StdoutErrorLogger.Println(logMessage)
		default:
			fmt.Println(logMessage)
		}
	}
}
package main

import (
	"fmt"
	"log"
	"os"
)

const (
	ErrorLog       string = "ERROR"
	InformationLog string = "INFO"
	WarningLog     string = "WARN"
)

func DDNSLogger(logType, hosts, domain, message string) {

	var (
		StdoutInfoLogger    *log.Logger
		StdoutWarningLogger *log.Logger
		StdoutErrorLogger   *log.Logger
	)

	StdoutInfoLogger = log.New(os.Stdout, "INFO ", log.Ldate|log.Ltime)
	StdoutWarningLogger = log.New(os.Stdout, "WARNING ", log.Ldate|log.Ltime)
	StdoutErrorLogger = log.New(os.Stdout, "ERROR ", log.Ldate|log.Ltime)

	if logType == "INFO" {
		StdoutInfoLogger.Println(hosts+"."+domain, message)
	} else if logType == "WARN" {
		StdoutWarningLogger.Println(hosts+"."+domain, message)
	} else if logType == "ERROR" {
		StdoutErrorLogger.Println(hosts+"."+domain, message)
	} else {
		fmt.Println(hosts+"."+domain, message)
	}
}

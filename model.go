package main

import (
	"fmt"
	"time"
)

type CustomError struct {
	ErrorCode int
	Err       error
}

func (err *CustomError) Error() string {
	return fmt.Sprintf("Error: %v, StatusCode: %d", err.Err, err.ErrorCode)
}

var (
	version          string        = "1.3.1-go1.23"
	daemon_poll_time time.Duration = 1 * time.Minute // Time in minute
	gitrepo          string        = "https://github.com/navilg/namecheap-ddns-docker"
	httpTimeout      time.Duration = 30 * time.Second
	expiryTime       float64       = 86400 // Ip env timeout in seconds (24hrs.)
	// expiryTime float64 = 600
)

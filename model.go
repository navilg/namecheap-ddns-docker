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
	version          string        = "1.0.0-go1.17"
	daemon_poll_time time.Duration = 1 * time.Minute // Time in minute
	gitrepo          string        = "https://github.com/navilg/namecheap-ddns-docker"
)

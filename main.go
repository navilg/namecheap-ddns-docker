package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Namecheap Dynamic DNS client Version", version)
	fmt.Println("Git Repo:", gitrepo)

	domain := flag.String("domain", "", "Domain name e.g. example.com")
	host := flag.String("host", "", "Subdomain or hostname e.g. www")
	password := flag.String("password", "", "Dynamic DNS Password from Namecheap")
	custom_ipcheck_url := flag.String("custom-ipcheck-url", "", "Custom IP check URL. Script always falls back to default ones")
	logLevelFlag := flag.String("log-level", InformationLog, "Log level (DEBUG, INFO, WARN, ERROR) - default is INFO")

	flag.Parse()

	// Set log level based on flag
	logLevel = *logLevelFlag

	if logLevel == "" {
		logLevel = InformationLog
	}

	// Validate log level
	validLogLevels := map[string]bool{
		DebugLog:       true,
		InformationLog: true,
		WarningLog:     true,
		ErrorLog:       true,
	}

	if !validLogLevels[logLevel] {
		fmt.Println("ERROR: Invalid log level. Supported values are DEBUG, INFO, WARN, ERROR")
		os.Exit(1)
	}

	if *domain == "" || *host == "" || *password == "" {
		fmt.Println("ERROR: domain, host and Dynamic DDNS password are mandatory")
		fmt.Printf("\nUsage of %s:\n", os.Args[1])
		flag.PrintDefaults()
		os.Exit(1)
	}

	pubIp, err := getPubIP(*custom_ipcheck_url)
	if err != nil {
		DDNSLogger(ErrorLog, *host, *domain, err.Error())
	} else {
		if err = setDNSRecord(*host, *domain, *password, pubIp); err != nil {
			DDNSLogger(ErrorLog, *host, *domain, err.Error())
			DDNSLogger(WarningLog, *host, *domain, "Ignoring above error. If this is not right, Re-run the process after fixing the error")
		} else {
			DDNSLogger(InformationLog, *host, *domain, "Record updated. "+pubIp)
		}
	}

	updateRecord(*domain, *host, *password, *custom_ipcheck_url)
}
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Namecheap Dynamic DNS client Version", version)
	fmt.Println("Git Repo:", gitrepo)

	domain := flag.String("domain", "", "Domain name e.g. example.com")
	hosts := flag.String("host", "", "Subdomain or hostname e.g. www")
	password := flag.String("password", "", "Dynamic DNS Password from Namecheap")

	flag.Parse()
	if *domain == "" || *hosts == "" || *password == "" {
		fmt.Println("ERROR domain, host and Dynamic DDNS password are mandatory")
		fmt.Printf("\nUsage of %s:\n", os.Args[1])
		flag.PrintDefaults()
		os.Exit(1)
	}

	pubIp, err := getPubIP()
	if err != nil {
		DDNSLogger(ErrorLog, *hosts, *domain, err.Error())
	} else {
		DDNSLogger(InformationLog, *hosts, *domain, "Updating all hosts.")
		for _, host := range strings.Split(hosts, ",") {
			if err = setDNSRecord(*host, *domain, *password, pubIp); err != nil {
				DDNSLogger(ErrorLog, *hosts, *domain, err.Error())
				DDNSLogger(WarningLog, *hosts, *domain, "Above error occured while updating host " + host + ", ignoring. If this is not right, Re-run the process after fixing the error")
			} else {
				DDNSLogger(InformationLog, *hosts, *domain, "Record for " + host + " updated. "+pubIp)
			}
		}
	}

	updateRecord(*domain, *hosts, *password)
}

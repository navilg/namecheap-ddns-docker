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

	flag.Parse()
	if *domain == "" || *host == "" || *password == "" {
		fmt.Println("ERROR domain, host and Dynamic DDNS password are mandatory")
		fmt.Printf("\nUsage of %s:\n", os.Args[1])
		flag.PrintDefaults()
		os.Exit(1)
	}

	pubIp, err := getPubIP()
	if err != nil {
		DDNSLogger(ErrorLog, *host, *domain, "Failed to get public Ip of your machine. "+err.Error())
	} else {
		setDNSRecord(*host, *domain, *password, pubIp)
		DDNSLogger(InformationLog, *host, *domain, "Record updated.")
	}

	updateRecord(*domain, *host, *password)
}

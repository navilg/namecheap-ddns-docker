package main

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"
)

func updateRecord(domain, host, password string) {

	DDNSLogger(InformationLog, "", "", "Started daemon process")

	ticker := time.NewTicker(daemon_poll_time)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return

			case <-ticker.C:
				pubIp, err := getPubIP()
				if err != nil {
					DDNSLogger(ErrorLog, host, domain, err.Error())
				}

				if parsedIp := net.ParseIP(pubIp); parsedIp == nil {
					DDNSLogger(ErrorLog, host, domain, "Invalid pubIp - "+pubIp+" (This could be due to non-existent of public IP or internet issue)")
				}

				currentIp := os.Getenv("NC_PUB_IP")
				lastIpUpdatedStr := os.Getenv("NC_PUB_IP_TIME")
				var lastIpUpdatedDuration float64

				fmt.Println(lastIpUpdatedStr)
				lastIpUpdated, err := time.Parse("2006-01-02 15:04:05", lastIpUpdatedStr)
				if err != nil {
					DDNSLogger(WarningLog, host, domain, "Not able to fetch last IP updated time. "+err.Error())
					lastIpUpdatedDuration = 0
				} else {
					currentTime := time.Now().Format("2006-01-02 15:04:05")
					currentTimeF, err := time.Parse("2006-01-02 15:04:05", currentTime)
					if err != nil {
						DDNSLogger(WarningLog, host, domain, "Not able to fetch last IP updated time. "+err.Error())
						lastIpUpdatedDuration = 0
					} else {
						lastIpUpdatedDuration = currentTimeF.Sub(lastIpUpdated).Seconds()
					}
				}

				if currentIp == pubIp && lastIpUpdatedDuration < expiryTime {
					// If currentIp is same as whats set in env var NC_PUB_IP AND last time IP updated in NC was less than 24 hrs ago.
					DDNSLogger(InformationLog, host, domain, "DNS record is same as current IP "+pubIp+". Last record update request made "+fmt.Sprintf("%f", lastIpUpdatedDuration)+" seconds ago.")
				} else {
					err = setDNSRecord(host, domain, password, pubIp)
					if err != nil {
						DDNSLogger(ErrorLog, host, domain, err.Error())
					} else {
						DDNSLogger(InformationLog, host, domain, "Record updated (ip: "+currentIp+"->"+pubIp+")")
					}
				}

			}
		}

	}()

	// Handle signal interrupt

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			DDNSLogger(InformationLog, "", "", "Interrupt signal received. Exiting")
			ticker.Stop()
			done <- true
			os.Exit(0)
		}
	}()

	time.Sleep(8760 * time.Hour) // Sleep for 365 days
	ticker.Stop()
	done <- true
}

func getPubIP() (string, error) {

	type GetIPBody struct {
		IP string `json:"ip"`
	}

	var ipbody GetIPBody

	apiclient := &http.Client{Timeout: httpTimeout}

	response, err := apiclient.Get("https://api.ipify.org?format=json")
	if err != nil {
		response, err = apiclient.Get("https://ipinfo.io/json")
		if err != nil {
			return "", nil
		}
	}

	defer response.Body.Close()
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		// fmt.Println(err.Error())
		return "", &CustomError{ErrorCode: response.StatusCode, Err: errors.New("IP could not be fetched." + err.Error())}
	}

	err = json.Unmarshal(bodyBytes, &ipbody)
	if err != nil {
		// fmt.Println(err.Error())
		return "", &CustomError{ErrorCode: response.StatusCode, Err: errors.New("IP could not be fetched." + err.Error())}
	}

	if ipbody.IP == "" {
		return "", &CustomError{ErrorCode: response.StatusCode, Err: errors.New("IP could not be fetched. Empty IP value detected")}
	}

	return ipbody.IP, nil

}

func setDNSRecord(host, domain, password, pubIp string) error {

	type InterfaceError struct {
		Err1 string `xml:"Err1"`
	}

	type InterfaceResponse struct {
		ErrorCount int            `xml:"ErrCount"`
		Errors     InterfaceError `xml:"errors"`
	}

	var interfaceResponse InterfaceResponse

	// Link from Namecheap knowledge article.
	// https://www.namecheap.com/support/knowledgebase/article.aspx/29/11/how-to-dynamically-update-the-hosts-ip-with-an-http-request/

	ncURL := "https://dynamicdns.park-your-domain.com/update?host=" + host + "&domain=" + domain + "&password=" + password + "&ip=" + pubIp

	apiclient := &http.Client{Timeout: httpTimeout}

	req, err := http.NewRequest("GET", ncURL, nil)
	if err != nil {
		// fmt.Println(1, err.Error())
		return err
	}

	// req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*")
	// req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	// req.Header.Add("Connection", "keep-alive")

	response, err := apiclient.Do(req)
	if err != nil {
		// fmt.Println(2, err.Error())
		return err
	}

	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	// Below function removes first line (below line) from response body because golang xml encoder does not support utf-16
	// <?xml version="1.0" encoding="utf-16"?>
	modifyBodyBytes := func(bodyBytes []byte) []byte {

		bodyString := string(bodyBytes)

		read_lines := strings.Split(bodyString, "\n")

		var updatedString string

		for i, line := range read_lines {
			if i != 0 {
				updatedString = fmt.Sprintf("%s%s\n", updatedString, line)
			}
		}

		return []byte(updatedString)
	}

	err = xml.Unmarshal(modifyBodyBytes(bodyBytes), &interfaceResponse)
	if err != nil {
		return err
	}

	if interfaceResponse.ErrorCount != 0 {
		return &CustomError{ErrorCode: -1, Err: errors.New(interfaceResponse.Errors.Err1)}
	}

	currentTime := time.Now()

	os.Setenv("NC_PUB_IP", pubIp)
	os.Setenv("NC_PUB_IP_TIME", currentTime.Format("2006-01-02 15:04:05"))

	return nil
}

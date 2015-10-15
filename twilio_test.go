package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func TestCallLog(t *testing.T) {
	//    http.Get("https://api.twilio.com/2010-04-01/Accounts/AC91873a310817d94c635634653ec6dc92.json")
	client := &http.Client{}
	req, _ := http.NewRequest("get", "https://api.twilio.com/2010-04-01/Accounts/AC91873a310817d94c635634653ec6dc92/Usage/Records.json?Category=calls", nil)
	//    req, _ := http.NewRequest("get", "https://api.twilio.com/2010-04-01/Accounts/AC91873a310817d94c635634653ec6dc92/Usage/Records", nil)
	//    req, _ := http.NewRequest("get", "https://api.twilio.com/2010-04-01/Accounts/AC91873a310817d94c635634653ec6dc92/Usage/Records.json", nil)

	req.SetBasicAuth("AC91873a310817d94c635634653ec6dc92", "247871c64430fd010c32359c070727cd")
	//    req.Header.Set("Authorization", "AC91873a310817d94c635634653ec6dc92:247871c64430fd010c32359c070727cd")

	response, _ := client.Do(req)
	log.Println("response:", response)
	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		bodystr := string(body)
		fmt.Println(bodystr)
	} else {
		fmt.Println(response)
	}
}

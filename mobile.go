package main

import (
	"fmt"
	"net/http"
	// "net/url"
	"io/ioutil"
	"log"
)

func RequestMobile(config conf) []string {
	fmt.Println("Checking mobile ...")

	// open client
	client := &http.Client{}
	// base URL
	req, err := http.NewRequest("GET", "https://jonathanmh.com/web-scraping-golang-goquery", nil)
	if err != nil {
		log.Print(err)
		return []string{}
	}
	fmt.Println(req)
	// add query strings
	q := req.URL.Query()
	q.Add("api_key", "myapikey")
	req.URL.RawQuery = q.Encode()
	fmt.Println(req.URL.String())


	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return []string{}
	}

	 defer resp.Body.Close()
	 respBody, _ := ioutil.ReadAll(resp.Body)

	 fmt.Println(resp.Status)
	 fmt.Println(string(respBody))

	results := []string{"www.mobile.de/0", "www.mobile.de/1"}
	return results
}

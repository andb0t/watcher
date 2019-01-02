package main

import "fmt"

func RequestEbay(config conf) []string {
	fmt.Println("Checking ebay ...")
	results := []string{"www.ebay.de/0", "www.ebay.de/1"}
	return results
}

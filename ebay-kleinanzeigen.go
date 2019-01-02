package main

import "fmt"

func RequestEbayKleinanzeigen(config conf) []string {
	fmt.Println("Checking ebay-kleinanzeigen ...")
	results := []string{"www.ebay-kleinanzeigen.de/0", "www.ebay-kleinanzeigen.de/1"}
	return results
}

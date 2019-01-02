package main

import "fmt"

func RequestMobile(config conf) []string {
	fmt.Println("Checking mobile ...")
	results := []string{"www.mobile.de/0", "www.mobile.de/1"}
	return results
}

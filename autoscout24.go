package main

import "fmt"

func RequestAutoScout24(config conf) []string {
	fmt.Println("Checking autoscout24 ...")
	results := []string{"www.autoscout24.de/0", "www.autoscout24.de/1"}
	return results
}

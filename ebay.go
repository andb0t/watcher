package main

import "fmt"


func RequestEbay(config conf) ([]string) {
  fmt.Println("Checking ebay ...")
  results := []string{"www.url0.de", "www.url1.de"}
  return results
}

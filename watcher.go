package main

import "fmt"


func main() {
    fmt.Println("Print checking for offers ...")
    keywords := []string {"Gebrauchtwagen", "BMW", "Audi"}
    vetowords := []string {"Neuwagen"}
    websites := []string {"www.ebay.de", "www.ebay-kleinanzeigen.de"}
    fmt.Println("Looking for:", keywords)
    fmt.Println(" ... while vetoing:", vetowords)
    fmt.Println(" ... on:", websites)

    RequestEbay()
}

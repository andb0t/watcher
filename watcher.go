package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type conf struct {
	Keywords  []string `yaml:"keywords"`
	Vetowords []string `yaml:"vetowords"`
	Websites  []string `yaml:"websites"`
}

func (config *conf) getConf() *conf {
	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return config
}

func main() {
	var config conf
	config.getConf()
	fmt.Println("Print checking for offers ...")
	fmt.Println("Looking for:", config.Keywords)
	fmt.Println(" ... while vetoing:", config.Vetowords)
	fmt.Println(" ... on:", config.Websites)

	var results []string
	if contains(config.Websites, "ebay") {
		results = append(append(RequestEbay(config), results...))
	}
	if contains(config.Websites, "autoscout24") {
		results = append(append(RequestAutoScout24(config), results...))
	}
	if contains(config.Websites, "ebay-kleinanzeigen") {
		results = append(append(RequestEbayKleinanzeigen(config), results...))
	}
	fmt.Println(results)
	SendEmail(results)
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

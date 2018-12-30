package main

import (
  "fmt"
  "gopkg.in/yaml.v2"
  "io/ioutil"
  "log"
  )


type conf struct {
    Keywords []string `yaml:"keywords"`
    Vetowords []string `yaml:"vetowords"`
    Websites []string `yaml:"websites"`
}

func (c *conf) getConf() *conf {
    yamlFile, err := ioutil.ReadFile("conf.yaml")
    if err != nil {
        log.Printf("yamlFile.Get err   #%v ", err)
    }
    err = yaml.Unmarshal(yamlFile, c)
    if err != nil {
        log.Fatalf("Unmarshal: %v", err)
    }
    return c
}

func main() {
    var c conf
    c.getConf()
    fmt.Println("Print checking for offers ...")
    fmt.Println("Looking for:", c.Keywords)
    fmt.Println(" ... while vetoing:", c.Vetowords)
    fmt.Println(" ... on:", c.Websites)

    RequestEbay()
}

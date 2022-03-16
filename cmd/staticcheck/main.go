package main

import (
	"encoding/json"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/multichecker"
	"honnef.co/go/tools/staticcheck"
	"log"
	"os"
)

type Config struct {
	Checks []string `json:"staticcheck"`
}

func checkErr(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}

func main() {
	data, err := os.ReadFile("config.json")
	checkErr(err)

	c := &Config{}

	if err = json.Unmarshal(data, c); err != nil {
		checkErr(err)
	}

	// определим map подключаемых правил
	checks := map[string]bool{}

	for _, check := range c.Checks {
		checks[check] = true
	}

	log.Println(checks)

	var mychecks []*analysis.Analyzer
	for _, v := range staticcheck.Analyzers {
		// добавляем в массив нужные проверки
		if checks[v.Analyzer.Name] {
			mychecks = append(mychecks, v.Analyzer)
		}
	}
	multichecker.Main(
		mychecks...,
	)
}

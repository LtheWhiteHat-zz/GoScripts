package main

import (
	"../../shodan"
	"fmt"
	"log"
	"os"
)

//USAGE: SHODAN_API_KEY={your key here} go run main.go target searchmethod
func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage: shodan searchterm")
	}
	apiKey := os.Getenv("SHODAN_API_KEY")
	s := shodan.New(apiKey)
	info, err := s.APIInfo()
	if err != nil {
		log.Panicln(err)
	}
	fmt.Printf(
		"Query Credits: %d\nScanCredits: %d\n\n",
		info.QueryCredits,
		info.ScanCredits)
	hostSearch, err := s.HostSearch(os.Args[1])
	if err != nil {
		log.Panicln(err)
	}
	for _, host := range hostSearch.Matches {
		fmt.Printf("%18s%8d -> %s %s %s %s\n", host.IPString, host.Port, host.Location.CountryName, host.Location.City, host.Org, host.OS)
	}
}

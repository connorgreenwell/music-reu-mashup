package main

import (
	"fmt"
	"os"
)

func main() {
	xml, err := os.Open("nsf.xml")
	if err != nil {
		panic(err)
	}
	defer xml.Close()

	sites := ParseSites(xml)

	fmt.Println(SitesWithKeyword(sites, "computer vision"))
}

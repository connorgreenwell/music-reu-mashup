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

	ml := SitesWithKeyword(sites, "machine learning")
	fmt.Println(len(ml))

	dm := SitesWithKeyword(sites, "data mining")
	fmt.Println(len(dm))

	cv := SitesWithKeyword(sites, "computer vision")
	fmt.Println(len(cv))

	combo := SitesWithKeyword(sites, "computer vision", "machine learning", "data mining")
	fmt.Println(len(combo))

	nc := SitesInState(sites, "California")
	for _, site := range nc {
		fmt.Println(site.Keywords)
	}
	fmt.Println("-----")
	nc_cv := SitesWithKeyword(nc, "computer vision", "machine learning", "data mining", "artifical intelligence")
	for _, site := range nc_cv {
		fmt.Println(site.Location, site.Keywords)
	}
}

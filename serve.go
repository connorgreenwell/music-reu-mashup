package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type SiteList struct {
	Sites []Site `json:"sites"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	xml, err := os.Open("nsf.xml")
	if err != nil {
		panic(err)
	}
	defer xml.Close()

	sites := ParseSites(xml)

	if r.Form["state"] != nil {
		fmt.Println("states!", r.Form["state"])
		sites = SitesInState(sites, r.Form["state"]...)
	}

	if r.Form["keyword"] != nil {
		fmt.Println("keywords!", r.Form["keyword"])
		sites = SitesWithKeyword(sites, r.Form["keyword"]...)
	}

	j, err := json.Marshal(SiteList{Sites: sites})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprint(w, (string)(j))
}

func main() {
	http.HandleFunc("/query", handler)
	http.ListenAndServe(":8080", nil)
}

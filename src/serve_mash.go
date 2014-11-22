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

	if r.Form["artist"] != nil {
		fmt.Println("artists!", r.Form["artist"])

		var states []string
		for _, artist := range r.Form["artist"] {
			states = append(states, StatesShowingIn(artist)...)
		}

		sites = SitesInState(sites, states...)
	}

	if r.Form["keyword"] != nil {
		fmt.Println("keywords!", r.Form["keyword"])
		sites = SitesWithKeyword(sites, r.Form["keyword"]...)
	}

	j, err := json.MarshalIndent(SiteList{Sites: sites}, "", "  ")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(r.Form, len(sites))
	fmt.Fprint(w, (string)(j))
}

func main() {
	http.HandleFunc("/query", handler)
	http.ListenAndServe(":8080", nil)
}

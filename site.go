package main

import (
	"strings"
)

type Site struct {
	SiteName   string    `json:"site_name"`
	SiteUrl    string    `json:"site_url"`
	Department string    `json:"department"`
	Location   Location  `json:"location"`
	Contacts   []Contact `json:"contacts"`
	Keywords   []string  `json:"keywords"`
	Comments   string    `json:"comments"`
	Award      string    `json:"award"`
	Cofunded   []string  `json:"confunded"`
}

type Location struct {
	City    string `json:"city"`
	State   string `json:"state"`
	Zipcode string `json:"zipcode"`
	Country string `json:"country"`
}

type Contact struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

func SitesWithKeyword(siteList []Site, targets ...string) (sites []Site) {
	for _, site := range siteList {
		found := false
		for _, keyword := range site.Keywords {
			for _, target := range targets {
				if strings.Contains(keyword, target) {
					sites = append(sites, site)
					found = true
					break
				}
			}
			if found {
				break
			}
		}
	}
	return
}

func SitesInState(siteList []Site, states ...string) (sites []Site) {
	for _, site := range siteList {
		for _, state := range states {
			if strings.Contains(site.Location.State, state) {
				sites = append(sites, site)
				break
			}
		}
	}
	return
}

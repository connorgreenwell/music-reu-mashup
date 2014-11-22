package main

import (
	"strings"
)

type Site struct {
	SiteName   string
	SiteUrl    string
	Department string
	Location   Location
	Contacts   []Contact
	Keywords   []string
	Comments   string
	Award      string
	Cofunded   []string
}

type Location struct {
	City    string
	State   string
	Zipcode string
	Country string
}

type Contact struct {
	Name  string
	Phone string
	Email string
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

package main

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"strings"
)

type Table struct {
	Rows []Row `xml:"row"`
}

type Row struct {
	Columns []string `xml:"column"`
}

func ParseSite(row Row) (site Site) {
	site.SiteName = strings.TrimSpace(row.Columns[0])
	site.SiteUrl = strings.TrimSpace(row.Columns[1])
	site.Department = strings.TrimSpace(row.Columns[2])

	loc := Location{
		City:    strings.TrimSpace(row.Columns[3]),
		State:   strings.TrimSpace(row.Columns[4]),
		Zipcode: strings.TrimSpace(row.Columns[5]),
		Country: strings.TrimSpace(row.Columns[6]),
	}
	site.Location = loc

	var contacts []Contact
	if row.Columns[7] != "" {
		con := Contact{
			Name:  strings.TrimSpace(row.Columns[7]),
			Phone: strings.TrimSpace(row.Columns[8]),
			Email: strings.TrimSpace(row.Columns[9]),
		}
		contacts = append(contacts, con)
	}

	if row.Columns[10] != "" {
		con := Contact{
			Name:  strings.TrimSpace(row.Columns[10]),
			Phone: strings.TrimSpace(row.Columns[11]),
			Email: strings.TrimSpace(row.Columns[12]),
		}
		contacts = append(contacts, con)
	}
	site.Contacts = contacts

	var keys []string
	for _, key := range strings.Split(row.Columns[13], ",") {
		keys = append(keys, strings.ToLower(strings.TrimSpace(key)))
	}
	site.Keywords = keys

	site.Comments = strings.TrimSpace(row.Columns[14])
	site.Award = strings.TrimSpace(row.Columns[15])

	var cofunds []string
	if row.Columns[16] != "" {
		cofunds = append(cofunds, strings.TrimSpace(row.Columns[16]))
	}
	if row.Columns[17] != "" {
		cofunds = append(cofunds, strings.TrimSpace(row.Columns[17]))
	}
	if row.Columns[18] != "" {
		cofunds = append(cofunds, strings.TrimSpace(row.Columns[18]))
	}
	site.Cofunded = cofunds

	return
}

func ParseSites(xml_file *os.File) (sites []Site) {
	xml_data, err := ioutil.ReadAll(xml_file)
	if err != nil {
		panic(err)
	}

	table := Table{}
	err = xml.Unmarshal(xml_data, &table)
	if err != nil {
		panic(err)
	}

	for _, row := range table.Rows[1:] {
		sites = append(sites, ParseSite(row))
	}

	return
}

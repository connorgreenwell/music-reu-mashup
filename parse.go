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
	site.SiteName = row.Columns[0]
	site.SiteUrl = row.Columns[1]
	site.Department = row.Columns[2]

	loc := Location{
		City:    row.Columns[3],
		State:   row.Columns[4],
		Zipcode: row.Columns[5],
		Country: row.Columns[6],
	}
	site.Location = loc

	var contacts []Contact
	if row.Columns[7] != "" {
		con := Contact{
			Name:  row.Columns[7],
			Phone: row.Columns[8],
			Email: row.Columns[9],
		}
		contacts = append(contacts, con)
	}

	if row.Columns[10] != "" {
		con := Contact{
			Name:  row.Columns[10],
			Phone: row.Columns[11],
			Email: row.Columns[12],
		}
		contacts = append(contacts, con)
	}
	site.Contacts = contacts

	var keys []string
	for _, key := range strings.Split(row.Columns[13], ",") {
		keys = append(keys, strings.ToLower(strings.TrimSpace(key)))
	}
	site.Keywords = keys

	site.Comments = row.Columns[14]
	site.Award = row.Columns[15]

	var cofunds []string
	if row.Columns[16] != "" {
		cofunds = append(cofunds, row.Columns[16])
	}
	if row.Columns[17] != "" {
		cofunds = append(cofunds, row.Columns[17])
	}
	if row.Columns[18] != "" {
		cofunds = append(cofunds, row.Columns[18])
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

package main

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

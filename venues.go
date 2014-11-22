package main

import (
	"fmt"
	"github.com/nf/geocode"
	"github.com/shkh/lastfm-go/lastfm"
)

const API_KEY = "adb8db28f9a1fe68937b7199cad6d554"
const API_SECRET = "02a672eda2e8999a9d5191da308189ed"

func StateFromZip(zipcode string) (state string) {
	req := &geocode.Request{
		Region:   "us",
		Provider: geocode.GOOGLE,
		Address:  zipcode,
	}

	resp, err := req.Lookup(nil)
	if err != nil {
		panic(err)
	}
	if s := resp.Status; s != "OK" {
		fmt.Println(s)
		return
	}

	for _, part := range resp.GoogleResponse.Results[0].AddressParts {
		if part.Types[0] == "administrative_area_level_1" {
			state = part.Name
			break
		}
	}

	return
}

func StatesShowingIn(artist string) (states []string) {
	api := lastfm.New(API_KEY, API_SECRET)

	res, err := api.Artist.GetEvents(lastfm.P{"artist": artist})
	if err != nil {
		panic(err)
	}
	for _, event := range res.Events {
		if event.Venue.Location.Country != "United States" {
			continue
		}
		if event.Venue.Location.Postalcode != "" {
			states = append(states, StateFromZip(event.Venue.Location.Postalcode))
		}
	}
	return
}

func main() {
	fmt.Println("Stars:", StatesShowingIn("Stars"))
	fmt.Println("Yo La Tengo:", StatesShowingIn("Yo La Tengo"))
	fmt.Println("Run The Jewels:", StatesShowingIn("Run The Jewels"))
}

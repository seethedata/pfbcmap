// pfbcapi responds to requests with a JSON object that lists all locations of Pathfinder Bootcamps
package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"googlemaps.github.io/maps"
	"io"
	"log"
	"net/http"
	"os"
)

var bootcampLocations = []bootcamp{
	{City: "Johannesburg, South Africa"},
	{City: "Los Angeles, CA"},
	{City: "Irvine, CA"},
	{City: "Salt Lake City, UT"},
	{City: "Tampa, FL"},
	{City: "San Diego, CA"},
	{City: "Atlanta, GA"},
	{City: "Phoenix, AZ"},
	{City: "Philadelphia, PA"},
	{City: "Houston, TX"},
	{City: "Dallas, TX"},
	{City: "Frankfurt, Germany"},
	{City: "Kansas City, MO"},
	{City: "Paris, France"},
	{City: "Tampa, FL"},
	{City: "Tulsa, OK"},
	{City: "Chicago, IL"},
	{City: "London, England"},
	{City: "Boston, MA"},
	{City: "Seattle, WA"},
	{City: "Charlotte, NC"},
	{City: "Asheville, NC"},
	{City: "Rockville, MD"},
	{City: "Denver, CO"},
	{City: "San Francisco, CA"},
	{City: "Oakland, CA"},
	{City: "Des Moines, IA"}}

type bootcamp struct {
	City     string      `json:"city"`
	Position maps.LatLng `json:"position"`
}

type templateData struct {
	Data string
}

func check(function string, e error) {
	if e != nil {
		log.Fatal(function, e)
	}
}

func responseHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	Data, err := json.Marshal(bootcampLocations)
	check("Marshal", err)
	io.WriteString(w, fmt.Sprintf("%s", Data))
}

func main() {
	c, err := maps.NewClient(maps.WithAPIKey(os.Getenv("googleAPIKey")))
	check("New Client", err)

	for val, location := range bootcampLocations {
		r := &maps.GeocodingRequest{Address: location.City}

		resp, err := c.Geocode(context.Background(), r)
		check("Geocode", err)

		bootcampLocations[val].Position = resp[0].Geometry.Location
	}

	http.HandleFunc("/data", responseHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

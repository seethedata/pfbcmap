// pfbcmap shows the locations of Pathfinder Bootcamps
package main

import (
	"encoding/json"
	"golang.org/x/net/context"
	"googlemaps.github.io/maps"
	"html/template"
	"log"
	"net/http"
	"os"
	"fmt"
)

type Bootcamp struct {
	City     string      `json:"city"`
	Position maps.LatLng `json:"position"`
}

type templateData  struct {
	Data string
}

var bootcampLocations = []Bootcamp{
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



func check(function string, e error) {
	if e != nil {
		log.Fatal(function, e)
	}
}

func responseHandler(w http.ResponseWriter, r *http.Request) {
	Data, err := json.Marshal(bootcampLocations)
	check("Marshal", err)
	var d templateData
	d.Data=fmt.Sprintf("%s",Data)
	t, err := template.ParseFiles("templates/index.tmpl")
	check("Parse template", err)
	t.Execute(w, d)
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


	http.HandleFunc("/", responseHandler)
	http.Handle("/images/", http.FileServer(http.Dir("")))
	http.Handle("/css/", http.FileServer(http.Dir("")))
	http.Handle("/fonts/", http.FileServer(http.Dir("")))
	http.Handle("/js/", http.FileServer(http.Dir("")))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

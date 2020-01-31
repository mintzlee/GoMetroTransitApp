package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

//PLAN
// on start, check for preexisting datafiles for route data (and stops for each route?)
// if existing and up to date (update every day?), continue, else go fetch new route data and save to file with iso8601 timestamp
// give usage output to user. be friendly :)
// give user options to print route names and stop names for a given route.
// check for user input, compare route and stop names to those in file. If not existing, print error and ask to try again.
// check that input direction is valid for given route.
// if valid input, go getTimepointDepartures data, return departureTime and departureText to user

func main() {
	currentTime := time.Now().Format(time.RFC3339)
	fmt.Printf("Current time: " + currentTime)

	fmt.Printf("\n\n" + hello() + "\n\nRoutes\n")
	// getRoutes()
	// fmt.Printf("\n\nProviders")
	// getProviders()
	// fmt.Printf("\n\nDirections")

	readSampleData()

}

// convert into greeting / usage print
func hello() string {
	return ("Hello, world\n")
}

// read in sample files as byte array, convert into structs for use. okay? okay.
func readSampleData() {
	jsonbytes, e := ioutil.ReadFile("../sampleData/departuresSample.json")
	if e != nil {
		fmt.Printf("error")
		os.Exit(1)
	}

	data := NexTripDepartures{}

	_ = json.Unmarshal([]byte(jsonbytes), &data)

	for i := 0; i < len(data.NexTripDepartures); i++ {
		fmt.Printf("%d DepartureText: %s", i, data.NexTripDepartures[i].DepartureText)
	}
}

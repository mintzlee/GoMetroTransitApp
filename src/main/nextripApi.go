package main

import (
	"fmt"
	"os"
)

func getRoutes() {
	routesJson, e := getDataFromUrl(routesJsonUrl)

	if e != nil {
		fmt.Printf(e.Error())
		os.Exit(1)
	}
	fmt.Println(routesJson)
}

func getProviders() {
	providersJson, e := getDataFromUrl(providersJsonUrl)

	if e != nil {
		fmt.Printf(e.Error())
		os.Exit(1)
	}
	fmt.Println(providersJson)
}

func getDirections(route int) {
	directionsJson, e := getDataFromUrl(getDirectionsJsonUrl(string(route)))
	if e != nil {
		fmt.Printf(e.Error())
		os.Exit(1)
	}
	fmt.Println(directionsJson)
}

// set up direction enum?
func getStops(route int, direction string) {
	stopsJson, e := getDataFromUrl(getStopsJsonUrl(string(route), direction))
	if e != nil {
		fmt.Printf(e.Error())
		os.Exit(1)
	}
	fmt.Println(stopsJson)
}

func getDepartures(stopId int) {
	departuresJson, e := getDataFromUrl(getDeparturesJsonUrl(string(stopId)))
	if e != nil {
		fmt.Printf(e.Error())
		os.Exit(1)
	}
	fmt.Println(departuresJson)
}

func getTimepointDepartures(route int, direction string, stop int) {
	timepointDeparturesJson, e := getDataFromUrl(getTimepointDeparturesJsonUrl(string(route), direction, string(stop)))
	if e != nil {
		fmt.Printf(e.Error())
		os.Exit(1)
	}
	fmt.Println(timepointDeparturesJson)
}

func getVehicleLocationsJson(route int) {
	vehicleLocationsJson, e := getDataFromUrl(getVehicleLocationsJsonUrl(string(route)))
	if e != nil {
		fmt.Printf(e.Error())
		os.Exit(1)
	}
	fmt.Println(vehicleLocationsJson)
}

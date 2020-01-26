package main

import (
	"fmt"
)

// move format=json query param to the getJsonfromurl function
func getRoutes() {
	routesJson := getJsonFromUrl("https://svc.metrotransit.org/NexTrip/Routes?format=json")
	fmt.Println(routesJson)
}

func getProviders() {
	providersJson := getJsonFromUrl("https://svc.metrotransit.org/NexTrip/Providers?format=json")
	fmt.Println(providersJson)
}

// TODO: make sure route string is correctly formatted before passing in
func getDirections(route string) {
	directionsJson := getJsonFromUrl("https://svc.metrotransit.org/NexTrip/Directions/" + route + "?format=json")
	fmt.Println(directionsJson)
}

// TODO: make sure route string is correctly formatted before passing in, set up direction enum?
func getStops(route string, direction string) {
	stopsJson := getJsonFromUrl("https://svc.metrotransit.org/NexTrip/Stops/" + route + "/" + direction + "?format=json")
	fmt.Println(stopsJson)
}

func getDepartures(stopId int) {
	departuresJson := getJsonFromUrl("https://svc.metrotransit.org/NexTrip/" + string(stopId) + "?format=json")
	fmt.Println(departuresJson)
}

// doublecheck api documentation for stop formatting for this call. ensure route formatted, direction enum?
func getTimepointDepartures(route string, direction string, stop string) {
	timepointDeparturesJson := getJsonFromUrl("https://svc.metrotransit.org/NexTrip/" + route + "/" + direction + "/" + stop + "?format=json")
	fmt.Println(timepointDeparturesJson)
}

func getVehicleLocationsJson(route string) {
	vehicleLocationsJson := getJsonFromUrl("https://svc.metrotransit.org/NexTrip/VehicleLocations/" + route + "?format=json")
	fmt.Println(vehicleLocationsJson)
}

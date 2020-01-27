package main

const jsonQueryParam = "?format=json"
const nexTripBaseUrl = "https://svc.metrotransit.org/NexTrip"

const routesUrl = nexTripBaseUrl + "/Routes"
const routesJsonUrl = routesUrl + jsonQueryParam

const providersUrl = nexTripBaseUrl + "/Providers"
const providersJsonUrl = providersUrl + jsonQueryParam

func getDirectionsJsonUrl(routeNum string) string {
	url := nexTripBaseUrl + "/Directions/" + routeNum + jsonQueryParam
	return url
}

func getStopsJsonUrl(routeNum string, direction string) string {
	url := nexTripBaseUrl + "/Stops/" + routeNum + "/" + direction + jsonQueryParam
	return url
}

func getDeparturesJsonUrl(stopId string) string {
	url := nexTripBaseUrl + "/" + stopId + jsonQueryParam
	return url
}

func getTimepointDeparturesJsonUrl(route string, direction string, stop string) string {
	url := nexTripBaseUrl + "/" + route + "/" + direction + "/" + stop + jsonQueryParam
	return url
}

func getVehicleLocationsJsonUrl(route string) string {
	url := nexTripBaseUrl + "/VehicleLocations/" + route + jsonQueryParam
	return url
}

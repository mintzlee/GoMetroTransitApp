package main

type apiOps string

const (
	GetProviders           string = "GET_PROVIDERS"
	GetRoutes              string = "GET_ROUTES"
	GetDirections          string = "GET_DIRECTIONS"
	GetStops               string = "GET_STOPS"
	GetDepartures          string = "GET_DEPARTURES"
	GetTimepointDepartures string = "GET_TIMEPOINT_DEPARTURES"
	GetVehicleLocations    string = "GET_VEHICLE_LOCATIONS"
)

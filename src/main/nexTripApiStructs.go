package main

type NexTripDeparture struct {
	Actual           bool
	BlockNumber      int
	DepartureText    string // can be nil, only returned when actual = true
	DepartureTime    string // iso format?
	Description      string
	Gate             string  // can be nil
	Route            string  // typically int for bus, possible string for trains? nope its an int for sure
	RouteDirection   string  // NORTHBOUND, SOUTHBOUND etc, make enum?
	Terminal         string  // can be nil
	VehicleHeading   int     // always 0, awaiting future development
	VehicleLatitude  float64 // (-)xx.xxxxx format. can be nil, only returned when actual = true
	VehicleLongitude float64 // (-)xx.xxxxx format. can be nil, only returned when actual = true
}

type NexTripDepartures struct {
	NexTripDepartures []NexTripDeparture
}

type NexTripRoute struct {
	Description string
	ProviderID  int
	Route       int
}

type NexTripRoutes struct {
	NexTripRoutes []NexTripRoute
}

type TextValuePair struct {
	Text  string
	Value string
}

type TextValuePairs struct {
	TextValuePairs []TextValuePair
}

type VehicleLocation struct {
	BlockNumber      int
	Direction        int
	LocationTime     string // iso format?
	Route            int
	Terminal         string
	VehicleLatitude  float64 // (-)xx.xxxxx format. can be nil, only returned when actual = true
	VehicleLongitude float64 // (-)xx.xxxxx format. can be nil, only returned when actual = true
	Bearing          int     // always 0, awaiting future development
	Odometer         int     // always 0, awaiting future development
	Speed            float64 // always 0, awaiting future development, samples show float
}

type VehicleLocations struct {
	VehicleLocations []VehicleLocation
}

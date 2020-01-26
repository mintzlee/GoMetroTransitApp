package main

type nexTripDeparture struct {
	actual           bool
	blockNumber      int
	departureText    string // can be null, only returned when actual = true
	departureTime    string // iso format?
	description      string
	gate             string  // can be nil
	route            string  // typically int for bus, possible string for trains?
	routeDirection   string  // NORTHBOUND, SOUTHBOUND etc, make enum?
	terminal         rune    // can be nil
	vehicleHeading   int     // always 0, awaiting future development
	vehicleLatitude  float32 // (-)xx.xxxxx format. can be nil, only returned when actual = true
	vehicleLongitude float32 // (-)xx.xxxxx format. can be nil, only returned when actual = true
}

type nexTripRoute struct {
	description string
	providerId  int // assumed always int. not sure.
	route       int
}

type textValuePair struct {
	text  string
	value string
}

type vehicleLocation struct {
	blockNumber      int
	direction        int
	locationTime     string // iso format?
	route            int
	terminal         rune    // char
	vehicleLatitude  float32 // (-)xx.xxxxx format. can be nil, only returned when actual = true
	vehicleLongitude float32 // (-)xx.xxxxx format. can be nil, only returned when actual = true
	bearing          int     // always 0, awaiting future development
	odometer         int     // always 0, awaiting future development
	speed            int     // always 0, awaiting future development

}

NexTrip api information taken from svc.metrotransit.org/nextrip

metrotransit.org web services
The NexTrip API is a real-time transit departure data web service for third-party application developers using Metro Transit information.

NexTrip departure information updates every 30 seconds. Help conserve our bandwidth and server resources by writing your applications responsibly. Third party applications should not update departure information more frequently than every 30 seconds. Applications making excessive calls and updating more frequently than 30 seconds will be subject to restriction.

There are seven requests with four different response schema.
See also http://svc.metrotransit.org/nextrip/help

Request operations
GetProviders operation - http://svc.metrotransit.org/NexTrip/Providers
Returns a list of area Transit providers. Providers are identified in the list of Routes allowing routes to be selected for a single provider.
GetRoutes operation - http://svc.metrotransit.org/NexTrip/Routes
Returns a list of Transit routes that are in service on the current day.
GetDirections operation - http://svc.metrotransit.org/NexTrip/Directions/{ROUTE}
Returns the two directions that are valid for a given route. Either North/South or East/West. The result includes text/value pair with the direction name and an ID. Directions are identified with an ID value. 1 = South, 2 = East, 3 = West, 4 = North.
GetStops operation - http://svc.metrotransit.org/NexTrip/Stops/{ROUTE}/{DIRECTION}
Returns a list of Timepoint stops for the given Route/Direction. The result includes text/value pairs with the stop description and a 4 character stop (or node) identifier.
GetDepartures operation - http://svc.metrotransit.org/NexTrip/{STOPID}
This operation is used to return a list of departures scheduled for any given bus stop. A StopID is an integer value identifying any one of the many thousands of bus stops in the metro. Stop information can be derived from the GTFS schedule data updated weekly for public use. https://gisdata.mn.gov/dataset/us-mn-state-metc-trans-transit-schedule-google-fd
GetTimepointDepartures operation - http://svc.metrotransit.org/NexTrip/{ROUTE}/{DIRECTION}/{STOP}
Returns the scheduled departures for a selected route, direction and timepoint stop.
GetVehicleLocations operation - http://svc.metrotransit.org/NexTrip/VehicleLocations/{ROUTE}
This operation returns a list of vehicles currently in service that have recently (within 5 minutes) reported their locations. A route paramter is used to return results for the given route. Use "0" for the route parameter to return a list of all vehicles in service.
Response schemas
NexTripDeparture
An array of NexTripDeparture elements is returned by the GetDepartures and GetTimepointDepartures operations. Each element consists of the fields:
Actual - bool indicates the departure time is based on current information from the vehicle.
BlockNumber - indicates the work for a vehicle.
DepartureText - displays time format for scheduled time and countdown format for real-time departures. (Actual=true)
DepartureTime - datetime value of the departure time.
Description - describes the trip destination.
Gate - indicates the stop or gate identifier where applicable.
Route - the current route for this departure.
RouteDirection - the current trip direction.
Terminal - the route branch letter where applicable.
VehicleHeading - this value is currently not available and always returns 0. (maybe someday)
VehicleLatitude - last reported latitude. only included with real-time departures. (Actual=true)
VehicleLongitude - last reported longitude. only included with real-time departures. (Actual=true)
NexTripRoute
An array of NexTripRoute elements is returned by the GetRoutes operation. Each element consists of the fields:
Description - description of the route.
ProviderID - identifier that corresponds to elements returned by the GetProviders operation.
Route - route number or label.
TextValuePair
The GetProviders, GetDirections and GetStops operations return arrays of TextValuePair. These simply consist of a Text and Value element. The text elements are labels and value elements contain identifiers. The direction and stop values, for example, along with the route number are required for the GetTimepointDepartures operation.
VehicleLocation
BlockNumber - indicates the work for a vehicle.
Direction - the direction ID of the current trip direction.
LocationTime - the time the location was last reported by the vehicle.
Route - the current route for the vehicle.
Terminal - the route branch letter where applicable.
VehicleLatitude - last reported latitude.
VehicleLongitude - last reported longitude.
Bearing - this value is currently not available and always returns 0. (for future use)
Odometer - this value is currently not available and always returns 0. (for future use)
Speed - this value is currently not available and always returns 0. (for future use)
Format Parameters
XML is the default response format. JSON can be requested by using a format paramter svc.metrotransit.org/NexTrip/Stops/21/3?format=json

Callback: To add a callback to a JSON request, include the callback parameter: svc.metrotransit.org/NexTrip/Directions/21?format=json&callback=cb

JSON can also be requested using appropriate request headers:
Content-Type: application/json
Accept: application/json
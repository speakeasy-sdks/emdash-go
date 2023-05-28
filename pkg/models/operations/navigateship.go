// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"SpaceTraders/pkg/models/shared"
	"net/http"
)

type NavigateShipSecurity struct {
	AgentToken string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
}

type NavigateShipRequestBody struct {
	// The target destination.
	WaypointSymbol string `json:"waypointSymbol"`
}

type NavigateShipRequest struct {
	RequestBody *NavigateShipRequestBody `request:"mediaType=application/json"`
	// The ship symbol
	ShipSymbol string `pathParam:"style=simple,explode=false,name=shipSymbol"`
}

type NavigateShip200ApplicationJSONData struct {
	// Details of the ship's fuel tanks including how much fuel was consumed during the last transit or action.
	Fuel shared.ShipFuel `json:"fuel"`
	// The navigation information of the ship.
	Nav shared.ShipNav `json:"nav"`
}

// NavigateShip200ApplicationJSON - The successful transit information including the route details and changes to ship fuel, supplies, and crew wages paid. The route includes the expected time of arrival.
type NavigateShip200ApplicationJSON struct {
	Data NavigateShip200ApplicationJSONData `json:"data"`
}

type NavigateShipResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// The successful transit information including the route details and changes to ship fuel, supplies, and crew wages paid. The route includes the expected time of arrival.
	NavigateShip200ApplicationJSONObject *NavigateShip200ApplicationJSON
}

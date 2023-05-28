// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"SpaceTraders/pkg/models/shared"
	"net/http"
)

type GetJumpGateRequest struct {
	// The system symbol
	SystemSymbol string `pathParam:"style=simple,explode=false,name=systemSymbol"`
	// The waypoint symbol
	WaypointSymbol string `pathParam:"style=simple,explode=false,name=waypointSymbol"`
}

// GetJumpGate200ApplicationJSON - OK
type GetJumpGate200ApplicationJSON struct {
	Data shared.JumpGate `json:"data"`
}

type GetJumpGateResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	GetJumpGate200ApplicationJSONObject *GetJumpGate200ApplicationJSON
}

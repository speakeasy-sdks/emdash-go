// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"SpaceTraders/pkg/models/shared"
	"net/http"
)

type GetMyShipsSecurity struct {
	AgentToken string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
}

type GetMyShipsRequest struct {
	// How many entries to return per page
	Limit *int64 `queryParam:"style=form,explode=true,name=limit"`
	// What entry offset to request
	Page *int64 `queryParam:"style=form,explode=true,name=page"`
}

// GetMyShips200ApplicationJSON - OK
type GetMyShips200ApplicationJSON struct {
	Data []shared.Ship `json:"data"`
	Meta shared.Meta   `json:"meta"`
}

type GetMyShipsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	GetMyShips200ApplicationJSONObject *GetMyShips200ApplicationJSON
}

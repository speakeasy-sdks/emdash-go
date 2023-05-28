// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"SpaceTraders/pkg/models/shared"
	"net/http"
)

type DeliverContractSecurity struct {
	AgentToken string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
}

type DeliverContractRequestBody struct {
	ShipSymbol  string `json:"shipSymbol"`
	TradeSymbol string `json:"tradeSymbol"`
	Units       int64  `json:"units"`
}

type DeliverContractRequest struct {
	RequestBody *DeliverContractRequestBody `request:"mediaType=application/json"`
	// The ID of the contract
	ContractID string `pathParam:"style=simple,explode=false,name=contractId"`
}

type DeliverContract200ApplicationJSONData struct {
	Cargo    shared.ShipCargo `json:"cargo"`
	Contract shared.Contract  `json:"contract"`
}

// DeliverContract200ApplicationJSON - OK
type DeliverContract200ApplicationJSON struct {
	Data DeliverContract200ApplicationJSONData `json:"data"`
}

type DeliverContractResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	DeliverContract200ApplicationJSONObject *DeliverContract200ApplicationJSON
}

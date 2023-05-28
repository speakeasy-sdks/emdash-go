// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"SpaceTraders/pkg/models/shared"
	"net/http"
)

type PurchaseCargoSecurity struct {
	AgentToken string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
}

type PurchaseCargoPurchaseCargoRequest struct {
	Symbol string `json:"symbol"`
	Units  int64  `json:"units"`
}

type PurchaseCargoRequest struct {
	RequestBody *PurchaseCargoPurchaseCargoRequest `request:"mediaType=application/json"`
	ShipSymbol  string                             `pathParam:"style=simple,explode=false,name=shipSymbol"`
}

type PurchaseCargoPurchaseCargo201ResponseData struct {
	Agent       shared.Agent             `json:"agent"`
	Cargo       shared.ShipCargo         `json:"cargo"`
	Transaction shared.MarketTransaction `json:"transaction"`
}

// PurchaseCargoPurchaseCargo201Response - Created
type PurchaseCargoPurchaseCargo201Response struct {
	Data PurchaseCargoPurchaseCargo201ResponseData `json:"data"`
}

type PurchaseCargoResponse struct {
	ContentType string
	// Created
	PurchaseCargo201Response *PurchaseCargoPurchaseCargo201Response
	StatusCode               int
	RawResponse              *http.Response
}

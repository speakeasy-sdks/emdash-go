// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"SpaceTraders/pkg/models/shared"
	"net/http"
)

type SellCargoSecurity struct {
	AgentToken string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
}

type SellCargoSellCargoRequest struct {
	Symbol string `json:"symbol"`
	Units  int64  `json:"units"`
}

type SellCargoRequest struct {
	RequestBody *SellCargoSellCargoRequest `request:"mediaType=application/json"`
	ShipSymbol  string                     `pathParam:"style=simple,explode=false,name=shipSymbol"`
}

type SellCargoSellCargo201ResponseData struct {
	Agent       shared.Agent             `json:"agent"`
	Cargo       shared.ShipCargo         `json:"cargo"`
	Transaction shared.MarketTransaction `json:"transaction"`
}

// SellCargoSellCargo201Response - Created
type SellCargoSellCargo201Response struct {
	Data SellCargoSellCargo201ResponseData `json:"data"`
}

type SellCargoResponse struct {
	ContentType string
	// Created
	SellCargo201Response *SellCargoSellCargo201Response
	StatusCode           int
	RawResponse          *http.Response
}

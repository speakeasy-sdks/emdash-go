// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"SpaceTraders/pkg/models/shared"
	"net/http"
)

type CreateSurveySecurity struct {
	AgentToken string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
}

type CreateSurveyRequest struct {
	// The symbol of the ship
	ShipSymbol string `pathParam:"style=simple,explode=false,name=shipSymbol"`
}

type CreateSurvey201ApplicationJSONData struct {
	// A cooldown is a period of time in which a ship cannot perform certain actions.
	Cooldown shared.Cooldown `json:"cooldown"`
	Surveys  []shared.Survey `json:"surveys"`
}

// CreateSurvey201ApplicationJSON - Created
type CreateSurvey201ApplicationJSON struct {
	Data CreateSurvey201ApplicationJSONData `json:"data"`
}

type CreateSurveyResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Created
	CreateSurvey201ApplicationJSONObject *CreateSurvey201ApplicationJSON
}

// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ShipRequirements - The requirements for installation on a ship
type ShipRequirements struct {
	// The number of crew required for operation.
	Crew *int64 `json:"crew,omitempty"`
	// The amount of power required from the reactor.
	Power *int64 `json:"power,omitempty"`
	// The number of module slots required for installation.
	Slots *int64 `json:"slots,omitempty"`
}

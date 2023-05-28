// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// Ship - A ship
type Ship struct {
	Cargo ShipCargo `json:"cargo"`
	// The ship's crew service and maintain the ship's systems and equipment.
	Crew ShipCrew `json:"crew"`
	// The engine determines how quickly a ship travels between waypoints.
	Engine ShipEngine `json:"engine"`
	// The frame of the ship. The frame determines the number of modules and mounting points of the ship, as well as base fuel capacity. As the condition of the frame takes more wear, the ship will become more sluggish and less maneuverable.
	Frame ShipFrame `json:"frame"`
	// Details of the ship's fuel tanks including how much fuel was consumed during the last transit or action.
	Fuel    ShipFuel     `json:"fuel"`
	Modules []ShipModule `json:"modules"`
	Mounts  []ShipMount  `json:"mounts"`
	// The navigation information of the ship.
	Nav ShipNav `json:"nav"`
	// The reactor of the ship. The reactor is responsible for powering the ship's systems and weapons.
	Reactor ShipReactor `json:"reactor"`
	// The public registration information of the ship
	Registration ShipRegistration `json:"registration"`
	// The globally unique identifier of the ship in the following format: `[AGENT_SYMBOL]_[HEX_ID]`
	Symbol string `json:"symbol"`
}

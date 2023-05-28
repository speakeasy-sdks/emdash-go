// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ScannedWaypoint - A waypoint is a location that ships can travel to such as a Planet, Moon or Space Station.
type ScannedWaypoint struct {
	// The chart of a system or waypoint, which makes the location visible to other agents.
	Chart        *Chart            `json:"chart,omitempty"`
	Faction      *WaypointFaction  `json:"faction,omitempty"`
	Orbitals     []WaypointOrbital `json:"orbitals"`
	Symbol       string            `json:"symbol"`
	SystemSymbol string            `json:"systemSymbol"`
	// The traits of the waypoint.
	Traits []WaypointTrait `json:"traits"`
	// The type of waypoint.
	Type WaypointType `json:"type"`
	X    int64        `json:"x"`
	Y    int64        `json:"y"`
}

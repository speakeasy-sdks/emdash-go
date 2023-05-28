// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ContractDeliverGood - The details of a delivery contract. Includes the type of good, units needed, and the destination.
type ContractDeliverGood struct {
	// The destination where goods need to be delivered.
	DestinationSymbol string `json:"destinationSymbol"`
	// The symbol of the trade good to deliver.
	TradeSymbol string `json:"tradeSymbol"`
	// The number of units fulfilled on this contract.
	UnitsFulfilled int64 `json:"unitsFulfilled"`
	// The number of units that need to be delivered on this contract.
	UnitsRequired int64 `json:"unitsRequired"`
}

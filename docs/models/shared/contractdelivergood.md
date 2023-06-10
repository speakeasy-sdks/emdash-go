# ContractDeliverGood

The details of a delivery contract. Includes the type of good, units needed, and the destination.


## Fields

| Field                                                           | Type                                                            | Required                                                        | Description                                                     |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| `DestinationSymbol`                                             | *string*                                                        | :heavy_check_mark:                                              | The destination where goods need to be delivered.               |
| `TradeSymbol`                                                   | *string*                                                        | :heavy_check_mark:                                              | The symbol of the trade good to deliver.                        |
| `UnitsFulfilled`                                                | *int64*                                                         | :heavy_check_mark:                                              | The number of units fulfilled on this contract.                 |
| `UnitsRequired`                                                 | *int64*                                                         | :heavy_check_mark:                                              | The number of units that need to be delivered on this contract. |
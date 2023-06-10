# ShipyardTransaction


## Fields

| Field                                                        | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `AgentSymbol`                                                | *string*                                                     | :heavy_check_mark:                                           | The symbol of the agent that made the transaction.           |
| `Price`                                                      | *int64*                                                      | :heavy_check_mark:                                           | The price of the transaction.                                |
| `ShipSymbol`                                                 | *string*                                                     | :heavy_check_mark:                                           | The symbol of the ship that was purchased.                   |
| `Timestamp`                                                  | [time.Time](https://pkg.go.dev/time#Time)                    | :heavy_check_mark:                                           | The timestamp of the transaction.                            |
| `WaypointSymbol`                                             | *string*                                                     | :heavy_check_mark:                                           | The symbol of the waypoint where the transaction took place. |
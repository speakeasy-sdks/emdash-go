# ShipFuelConsumed


## Fields

| Field                                                             | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `Amount`                                                          | *int64*                                                           | :heavy_check_mark:                                                | The amount of fuel consumed by the most recent transit or action. |
| `Timestamp`                                                       | [time.Time](https://pkg.go.dev/time#Time)                         | :heavy_check_mark:                                                | The time at which the fuel was consumed.                          |
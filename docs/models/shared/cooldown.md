# Cooldown

A cooldown is a period of time in which a ship cannot perform certain actions.


## Fields

| Field                                                          | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `Expiration`                                                   | [*time.Time](https://pkg.go.dev/time#Time)                     | :heavy_minus_sign:                                             | The date and time when the cooldown expires in ISO 8601 format |
| `RemainingSeconds`                                             | *int64*                                                        | :heavy_check_mark:                                             | The remaining duration of the cooldown in seconds              |
| `ShipSymbol`                                                   | *string*                                                       | :heavy_check_mark:                                             | The symbol of the ship that is on cooldown                     |
| `TotalSeconds`                                                 | *int64*                                                        | :heavy_check_mark:                                             | The total duration of the cooldown in seconds                  |
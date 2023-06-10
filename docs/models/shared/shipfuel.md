# ShipFuel

Details of the ship's fuel tanks including how much fuel was consumed during the last transit or action.


## Fields

| Field                                                        | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `Capacity`                                                   | *int64*                                                      | :heavy_check_mark:                                           | The maximum amount of fuel the ship's tanks can hold.        |
| `Consumed`                                                   | [*ShipFuelConsumed](../../models/shared/shipfuelconsumed.md) | :heavy_minus_sign:                                           | N/A                                                          |
| `Current`                                                    | *int64*                                                      | :heavy_check_mark:                                           | The current amount of fuel in the ship's tanks.              |
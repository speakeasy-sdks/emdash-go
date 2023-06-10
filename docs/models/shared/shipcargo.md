# ShipCargo


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `Capacity`                                                    | *int64*                                                       | :heavy_check_mark:                                            | The max number of items that can be stored in the cargo hold. |
| `Inventory`                                                   | [][ShipCargoItem](../../models/shared/shipcargoitem.md)       | :heavy_check_mark:                                            | The items currently in the cargo hold.                        |
| `Units`                                                       | *int64*                                                       | :heavy_check_mark:                                            | The number of items currently stored in the cargo hold.       |
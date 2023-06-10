# ShipModule

A module can be installed in a ship and provides a set of capabilities such as storage space or quarters for crew. Module installations are permanent.


## Fields

| Field                                                       | Type                                                        | Required                                                    | Description                                                 |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `Capacity`                                                  | **int64*                                                    | :heavy_minus_sign:                                          | N/A                                                         |
| `Description`                                               | **string*                                                   | :heavy_minus_sign:                                          | N/A                                                         |
| `Name`                                                      | *string*                                                    | :heavy_check_mark:                                          | N/A                                                         |
| `Range`                                                     | **int64*                                                    | :heavy_minus_sign:                                          | N/A                                                         |
| `Requirements`                                              | [ShipRequirements](../../models/shared/shiprequirements.md) | :heavy_check_mark:                                          | The requirements for installation on a ship                 |
| `Symbol`                                                    | [ShipModuleSymbol](../../models/shared/shipmodulesymbol.md) | :heavy_check_mark:                                          | N/A                                                         |
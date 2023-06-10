# ScannedShip

The ship that was scanned. Details include information about the ship that could be detected by the scanner.


## Fields

| Field                                                            | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `Engine`                                                         | [ScannedShipEngine](../../models/shared/scannedshipengine.md)    | :heavy_check_mark:                                               | The engine of the ship.                                          |
| `Frame`                                                          | [*ScannedShipFrame](../../models/shared/scannedshipframe.md)     | :heavy_minus_sign:                                               | The frame of the ship.                                           |
| `Mounts`                                                         | [][ScannedShipMounts](../../models/shared/scannedshipmounts.md)  | :heavy_minus_sign:                                               | N/A                                                              |
| `Nav`                                                            | [ShipNav](../../models/shared/shipnav.md)                        | :heavy_check_mark:                                               | The navigation information of the ship.                          |
| `Reactor`                                                        | [*ScannedShipReactor](../../models/shared/scannedshipreactor.md) | :heavy_minus_sign:                                               | The reactor of the ship.                                         |
| `Registration`                                                   | [ShipRegistration](../../models/shared/shipregistration.md)      | :heavy_check_mark:                                               | The public registration information of the ship                  |
| `Symbol`                                                         | *string*                                                         | :heavy_check_mark:                                               | The globally unique identifier of the ship.                      |
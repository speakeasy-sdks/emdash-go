# Fleet

## Overview

Fleet

### Available Operations

* [CreateChart](#createchart) - Create Chart
* [CreateShipShipScan](#createshipshipscan) - Scan Ships
* [CreateShipSystemScan](#createshipsystemscan) - Scan Systems
* [CreateShipWaypointScan](#createshipwaypointscan) - Scan Waypoints
* [CreateSurvey](#createsurvey) - Create Survey
* [DockShip](#dockship) - Dock Ship
* [ExtractResources](#extractresources) - Extract Resources
* [GetMyShip](#getmyship) - Get Ship
* [GetMyShipCargo](#getmyshipcargo) - Get Ship Cargo
* [GetMyShips](#getmyships) - List Ships
* [GetShipCooldown](#getshipcooldown) - Get Ship Cooldown
* [GetShipNav](#getshipnav) - Get Ship Nav
* [Jettison](#jettison) - Jettison Cargo
* [JumpShip](#jumpship) - Jump Ship
* [NavigateShip](#navigateship) - Navigate Ship
* [NegotiateContract](#negotiatecontract) - Negotiate Contract
* [OrbitShip](#orbitship) - Orbit Ship
* [PatchShipNav](#patchshipnav) - Patch Ship Nav
* [PurchaseCargo](#purchasecargo) - Purchase Cargo
* [PurchaseShip](#purchaseship) - Purchase Ship
* [RefuelShip](#refuelship) - Refuel Ship
* [SellCargo](#sellcargo) - Sell Cargo
* [ShipRefine](#shiprefine) - Ship Refine
* [TransferCargo](#transfercargo) - Transfer Cargo
* [WarpShip](#warpship) - Warp Ship

## CreateChart

Command a ship to chart the current waypoint.

Waypoints in the universe are uncharted by default. These locations will not show up in the API until they have been charted by a ship.

Charting a location will record your agent as the one who created the chart.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"SpaceTraders"
	"SpaceTraders/pkg/models/operations"
)

func main() {
    s := spacetraders.New()

    ctx := context.Background()
    res, err := s.Fleet.CreateChart(ctx, operations.CreateChartRequest{
        ShipSymbol: "molestiae",
    }, operations.CreateChartSecurity{
        AgentToken: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateChart201ApplicationJSONObject != nil {
        // handle response
    }
}
```

## CreateShipShipScan

Activate your ship's sensor arrays to scan for ship information.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"SpaceTraders"
	"SpaceTraders/pkg/models/operations"
)

func main() {
    s := spacetraders.New()

    ctx := context.Background()
    res, err := s.Fleet.CreateShipShipScan(ctx, operations.CreateShipShipScanRequest{
        ShipSymbol: "minus",
    }, operations.CreateShipShipScanSecurity{
        AgentToken: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateShipShipScan201ApplicationJSONObject != nil {
        // handle response
    }
}
```

## CreateShipSystemScan

Activate your ship's sensor arrays to scan for system information.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"SpaceTraders"
	"SpaceTraders/pkg/models/operations"
)

func main() {
    s := spacetraders.New()

    ctx := context.Background()
    res, err := s.Fleet.CreateShipSystemScan(ctx, operations.CreateShipSystemScanRequest{
        ShipSymbol: "placeat",
    }, operations.CreateShipSystemScanSecurity{
        AgentToken: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateShipSystemScan201ApplicationJSONObject != nil {
        // handle response
    }
}
```

## CreateShipWaypointScan

Activate your ship's sensor arrays to scan for waypoint information.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"SpaceTraders"
	"SpaceTraders/pkg/models/operations"
)

func main() {
    s := spacetraders.New()

    ctx := context.Background()
    res, err := s.Fleet.CreateShipWaypointScan(ctx, operations.CreateShipWaypointScanRequest{
        ShipSymbol: "voluptatum",
    }, operations.CreateShipWaypointScanSecurity{
        AgentToken: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateShipWaypointScan201ApplicationJSONObject != nil {
        // handle response
    }
}
```

## CreateSurvey

If you want to target specific yields for an extraction, you can survey a waypoint, such as an asteroid field, and send the survey in the body of the extract request. Each survey may have multiple deposits, and if a symbol shows up more than once, that indicates a higher chance of extracting that resource.

Your ship will enter a cooldown between consecutive survey requests. Surveys will eventually expire after a period of time. Multiple ships can use the same survey for extraction.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"SpaceTraders"
	"SpaceTraders/pkg/models/operations"
)

func main() {
    s := spacetraders.New()

    ctx := context.Background()
    res, err := s.Fleet.CreateSurvey(ctx, operations.CreateSurveyRequest{
        ShipSymbol: "iusto",
    }, operations.CreateSurveySecurity{
        AgentToken: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateSurvey201ApplicationJSONObject != nil {
        // handle response
    }
}
```

## DockShip

Attempt to dock your ship at it's current location. Docking will only succeed if the waypoint is a dockable location, and your ship is capable of docking at the time of the request.

The endpoint is idempotent - successive calls will succeed even if the ship is already docked.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"SpaceTraders"
	"SpaceTraders/pkg/models/operations"
)

func main() {
    s := spacetraders.New()

    ctx := context.Background()
    res, err := s.Fleet.DockShip(ctx, operations.DockShipRequest{
        ShipSymbol: "excepturi",
    }, operations.DockShipSecurity{
        AgentToken: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DockShip200Response != nil {
        // handle response
    }
}
```

## ExtractResources

Extract resources from the waypoint into your ship. Send an optional survey as the payload to target specific yields.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"SpaceTraders"
	"SpaceTraders/pkg/models/operations"
	"SpaceTraders/pkg/models/shared"
	"SpaceTraders/pkg/types"
)

func main() {
    s := spacetraders.New()

    ctx := context.Background()
    res, err := s.Fleet.ExtractResources(ctx, operations.ExtractResourcesRequest{
        RequestBody: &operations.ExtractResourcesRequestBody{
            Survey: &shared.Survey{
                Deposits: []shared.SurveyDeposit{
                    shared.SurveyDeposit{
                        Symbol: "recusandae",
                    },
                    shared.SurveyDeposit{
                        Symbol: "temporibus",
                    },
                },
                Expiration: types.MustTimeFromString("2022-08-30T20:24:33.984Z"),
                Signature: "veritatis",
                Size: shared.SurveySizeModerate,
                Symbol: "perferendis",
            },
        },
        ShipSymbol: "ipsam",
    }, operations.ExtractResourcesSecurity{
        AgentToken: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ExtractResources201ApplicationJSONObject != nil {
        // handle response
    }
}
```

## GetMyShip

Retrieve the details of your ship.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"SpaceTraders"
	"SpaceTraders/pkg/models/operations"
)

func main() {
    s := spacetraders.New()

    ctx := context.Background()
    res, err := s.Fleet.GetMyShip(ctx, operations.GetMyShipRequest{
        ShipSymbol: "repellendus",
    }, operations.GetMyShipSecurity{
        AgentToken: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetMyShip200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## GetMyShipCargo

Retrieve the cargo of your ship.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"SpaceTraders"
	"SpaceTraders/pkg/models/operations"
)

func main() {
    s := spacetraders.New()

    ctx := context.Background()
    res, err := s.Fleet.GetMyShipCargo(ctx, operations.GetMyShipCargoRequest{
        ShipSymbol: "sapiente",
    }, operations.GetMyShipCargoSecurity{
        AgentToken: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetMyShipCargo200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## GetMyShips

Retrieve all of your ships.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"SpaceTraders"
	"SpaceTraders/pkg/models/operations"
)

func main() {
    s := spacetraders.New()

    ctx := context.Background()
    res, err := s.Fleet.GetMyShips(ctx, operations.GetMyShipsRequest{
        Limit: spacetraders.Int64(778157),
        Page: spacetraders.Int64(140350),
    }, operations.GetMyShipsSecurity{
        AgentToken: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetMyShips200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## GetShipCooldown

Retrieve the details of your ship's reactor cooldown. Some actions such as activating your jump drive, scanning, or extracting resources taxes your reactor and results in a cooldown.

Your ship cannot perform additional actions until your cooldown has expired. The duration of your cooldown is relative to the power consumption of the related modules or mounts for the action taken.

Response returns a 204 status code (no-content) when the ship has no cooldown.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"SpaceTraders"
	"SpaceTraders/pkg/models/operations"
)

func main() {
    s := spacetraders.New()

    ctx := context.Background()
    res, err := s.Fleet.GetShipCooldown(ctx, operations.GetShipCooldownRequest{
        ShipSymbol: "at",
    }, operations.GetShipCooldownSecurity{
        AgentToken: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetShipCooldown200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## GetShipNav

Get the current nav status of a ship.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"SpaceTraders"
	"SpaceTraders/pkg/models/operations"
)

func main() {
    s := spacetraders.New()

    ctx := context.Background()
    res, err := s.Fleet.GetShipNav(ctx, operations.GetShipNavRequest{
        ShipSymbol: "at",
    }, operations.GetShipNavSecurity{
        AgentToken: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetShipNav200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## Jettison

Jettison cargo from your ship's cargo hold.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"SpaceTraders"
	"SpaceTraders/pkg/models/operations"
)

func main() {
    s := spacetraders.New()

    ctx := context.Background()
    res, err := s.Fleet.Jettison(ctx, operations.JettisonRequest{
        RequestBody: &operations.JettisonRequestBody{
            Symbol: "maiores",
            Units: 473608,
        },
        ShipSymbol: "quod",
    }, operations.JettisonSecurity{
        AgentToken: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Jettison200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## JumpShip

Jump your ship instantly to a target system. When used while in orbit or docked to a jump gate waypoint, any ship can use this command. When used elsewhere, jumping requires a jump drive unit and consumes a unit of antimatter (which needs to be in your cargo).

### Example Usage

```go
package main

import(
	"context"
	"log"
	"SpaceTraders"
	"SpaceTraders/pkg/models/operations"
)

func main() {
    s := spacetraders.New()

    ctx := context.Background()
    res, err := s.Fleet.JumpShip(ctx, operations.JumpShipRequest{
        RequestBody: &operations.JumpShipRequestBody{
            SystemSymbol: "quod",
        },
        ShipSymbol: "esse",
    }, operations.JumpShipSecurity{
        AgentToken: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.JumpShip200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## NavigateShip

Navigate to a target destination. The destination must be located within the same system as the ship. Navigating will consume the necessary fuel and supplies from the ship's manifest, and will pay out crew wages from the agent's account.

The returned response will detail the route information including the expected time of arrival. Most ship actions are unavailable until the ship has arrived at it's destination.

To travel between systems, see the ship's warp or jump actions.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"SpaceTraders"
	"SpaceTraders/pkg/models/operations"
)

func main() {
    s := spacetraders.New()

    ctx := context.Background()
    res, err := s.Fleet.NavigateShip(ctx, operations.NavigateShipRequest{
        RequestBody: &operations.NavigateShipRequestBody{
            WaypointSymbol: "totam",
        },
        ShipSymbol: "porro",
    }, operations.NavigateShipSecurity{
        AgentToken: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.NavigateShip200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## NegotiateContract

Negotiate Contract

### Example Usage

```go
package main

import(
	"context"
	"log"
	"SpaceTraders"
	"SpaceTraders/pkg/models/operations"
)

func main() {
    s := spacetraders.New()

    ctx := context.Background()
    res, err := s.Fleet.NegotiateContract(ctx, operations.NegotiateContractRequest{
        RequestBody: spacetraders.String("dolorum"),
        ShipSymbol: "dicta",
    }, operations.NegotiateContractSecurity{
        AgentToken: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.NegotiateContract200Response != nil {
        // handle response
    }
}
```

## OrbitShip

Attempt to move your ship into orbit at it's current location. The request will only succeed if your ship is capable of moving into orbit at the time of the request.

The endpoint is idempotent - successive calls will succeed even if the ship is already in orbit.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"SpaceTraders"
	"SpaceTraders/pkg/models/operations"
)

func main() {
    s := spacetraders.New()

    ctx := context.Background()
    res, err := s.Fleet.OrbitShip(ctx, operations.OrbitShipRequest{
        ShipSymbol: "nam",
    }, operations.OrbitShipSecurity{
        AgentToken: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.OrbitShip200Response != nil {
        // handle response
    }
}
```

## PatchShipNav

Update the nav data of a ship, such as the flight mode.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"SpaceTraders"
	"SpaceTraders/pkg/models/operations"
	"SpaceTraders/pkg/models/shared"
)

func main() {
    s := spacetraders.New()

    ctx := context.Background()
    res, err := s.Fleet.PatchShipNav(ctx, operations.PatchShipNavRequest{
        RequestBody: &operations.PatchShipNavRequestBody{
            FlightMode: shared.ShipNavFlightModeCruise.ToPointer(),
        },
        ShipSymbol: "occaecati",
    }, operations.PatchShipNavSecurity{
        AgentToken: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PatchShipNav200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## PurchaseCargo

Purchase cargo.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"SpaceTraders"
	"SpaceTraders/pkg/models/operations"
)

func main() {
    s := spacetraders.New()

    ctx := context.Background()
    res, err := s.Fleet.PurchaseCargo(ctx, operations.PurchaseCargoRequest{
        RequestBody: &operations.PurchaseCargoPurchaseCargoRequest{
            Symbol: "fugit",
            Units: 537373,
        },
        ShipSymbol: "hic",
    }, operations.PurchaseCargoSecurity{
        AgentToken: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PurchaseCargo201Response != nil {
        // handle response
    }
}
```

## PurchaseShip

Purchase a ship

### Example Usage

```go
package main

import(
	"context"
	"log"
	"SpaceTraders"
	"SpaceTraders/pkg/models/operations"
	"SpaceTraders/pkg/models/shared"
)

func main() {
    s := spacetraders.New()

    ctx := context.Background()
    res, err := s.Fleet.PurchaseShip(ctx, operations.PurchaseShipRequestBody{
        ShipType: shared.ShipTypeShipLightShuttle,
        WaypointSymbol: "totam",
    }, operations.PurchaseShipSecurity{
        AgentToken: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PurchaseShip201ApplicationJSONObject != nil {
        // handle response
    }
}
```

## RefuelShip

Refuel your ship from the local market.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"SpaceTraders"
	"SpaceTraders/pkg/models/operations"
)

func main() {
    s := spacetraders.New()

    ctx := context.Background()
    res, err := s.Fleet.RefuelShip(ctx, operations.RefuelShipRequest{
        ShipSymbol: "beatae",
    }, operations.RefuelShipSecurity{
        AgentToken: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RefuelShip200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## SellCargo

Sell cargo.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"SpaceTraders"
	"SpaceTraders/pkg/models/operations"
)

func main() {
    s := spacetraders.New()

    ctx := context.Background()
    res, err := s.Fleet.SellCargo(ctx, operations.SellCargoRequest{
        RequestBody: &operations.SellCargoSellCargoRequest{
            Symbol: "commodi",
            Units: 473600,
        },
        ShipSymbol: "modi",
    }, operations.SellCargoSecurity{
        AgentToken: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SellCargo201Response != nil {
        // handle response
    }
}
```

## ShipRefine

Attempt to refine the raw materials on your ship. The request will only succeed if your ship is capable of refining at the time of the request.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"SpaceTraders"
	"SpaceTraders/pkg/models/operations"
)

func main() {
    s := spacetraders.New()

    ctx := context.Background()
    res, err := s.Fleet.ShipRefine(ctx, operations.ShipRefineRequest{
        RequestBody: &operations.ShipRefineRequestBody{
            Produce: operations.ShipRefineRequestBodyProduceCopper,
        },
        ShipSymbol: "impedit",
    }, operations.ShipRefineSecurity{
        AgentToken: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ShipRefine200Response != nil {
        // handle response
    }
}
```

## TransferCargo

Transfer cargo between ships.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"SpaceTraders"
	"SpaceTraders/pkg/models/operations"
)

func main() {
    s := spacetraders.New()

    ctx := context.Background()
    res, err := s.Fleet.TransferCargo(ctx, operations.TransferCargoRequest{
        RequestBody: &operations.TransferCargoTransferCargoRequest{
            ShipSymbol: "cum",
            TradeSymbol: "esse",
            Units: 216550,
        },
        ShipSymbol: "excepturi",
    }, operations.TransferCargoSecurity{
        AgentToken: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TransferCargo200Response != nil {
        // handle response
    }
}
```

## WarpShip

Warp your ship to a target destination in another system. Warping will consume the necessary fuel and supplies from the ship's manifest, and will pay out crew wages from the agent's account.

The returned response will detail the route information including the expected time of arrival. Most ship actions are unavailable until the ship has arrived at it's destination.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"SpaceTraders"
	"SpaceTraders/pkg/models/operations"
)

func main() {
    s := spacetraders.New()

    ctx := context.Background()
    res, err := s.Fleet.WarpShip(ctx, operations.WarpShipRequest{
        RequestBody: &operations.WarpShipRequestBody{
            WaypointSymbol: "aspernatur",
        },
        ShipSymbol: "perferendis",
    }, operations.WarpShipSecurity{
        AgentToken: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.WarpShip200ApplicationJSONObject != nil {
        // handle response
    }
}
```

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

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.CreateChartRequest](../../models/operations/createchartrequest.md)   | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `security`                                                                       | [operations.CreateChartSecurity](../../models/operations/createchartsecurity.md) | :heavy_check_mark:                                                               | The security requirements to use for the request.                                |


### Response

**[*operations.CreateChartResponse](../../models/operations/createchartresponse.md), error**


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

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.CreateShipShipScanRequest](../../models/operations/createshipshipscanrequest.md)   | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `security`                                                                                     | [operations.CreateShipShipScanSecurity](../../models/operations/createshipshipscansecurity.md) | :heavy_check_mark:                                                                             | The security requirements to use for the request.                                              |


### Response

**[*operations.CreateShipShipScanResponse](../../models/operations/createshipshipscanresponse.md), error**


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

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.CreateShipSystemScanRequest](../../models/operations/createshipsystemscanrequest.md)   | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `security`                                                                                         | [operations.CreateShipSystemScanSecurity](../../models/operations/createshipsystemscansecurity.md) | :heavy_check_mark:                                                                                 | The security requirements to use for the request.                                                  |


### Response

**[*operations.CreateShipSystemScanResponse](../../models/operations/createshipsystemscanresponse.md), error**


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

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.CreateShipWaypointScanRequest](../../models/operations/createshipwaypointscanrequest.md)   | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `security`                                                                                             | [operations.CreateShipWaypointScanSecurity](../../models/operations/createshipwaypointscansecurity.md) | :heavy_check_mark:                                                                                     | The security requirements to use for the request.                                                      |


### Response

**[*operations.CreateShipWaypointScanResponse](../../models/operations/createshipwaypointscanresponse.md), error**


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

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.CreateSurveyRequest](../../models/operations/createsurveyrequest.md)   | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `security`                                                                         | [operations.CreateSurveySecurity](../../models/operations/createsurveysecurity.md) | :heavy_check_mark:                                                                 | The security requirements to use for the request.                                  |


### Response

**[*operations.CreateSurveyResponse](../../models/operations/createsurveyresponse.md), error**


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

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.DockShipRequest](../../models/operations/dockshiprequest.md)   | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `security`                                                                 | [operations.DockShipSecurity](../../models/operations/dockshipsecurity.md) | :heavy_check_mark:                                                         | The security requirements to use for the request.                          |


### Response

**[*operations.DockShipResponse](../../models/operations/dockshipresponse.md), error**


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

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ExtractResourcesRequest](../../models/operations/extractresourcesrequest.md)   | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `security`                                                                                 | [operations.ExtractResourcesSecurity](../../models/operations/extractresourcessecurity.md) | :heavy_check_mark:                                                                         | The security requirements to use for the request.                                          |


### Response

**[*operations.ExtractResourcesResponse](../../models/operations/extractresourcesresponse.md), error**


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

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.GetMyShipRequest](../../models/operations/getmyshiprequest.md)   | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `security`                                                                   | [operations.GetMyShipSecurity](../../models/operations/getmyshipsecurity.md) | :heavy_check_mark:                                                           | The security requirements to use for the request.                            |


### Response

**[*operations.GetMyShipResponse](../../models/operations/getmyshipresponse.md), error**


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

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetMyShipCargoRequest](../../models/operations/getmyshipcargorequest.md)   | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `security`                                                                             | [operations.GetMyShipCargoSecurity](../../models/operations/getmyshipcargosecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |


### Response

**[*operations.GetMyShipCargoResponse](../../models/operations/getmyshipcargoresponse.md), error**


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

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.GetMyShipsRequest](../../models/operations/getmyshipsrequest.md)   | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `security`                                                                     | [operations.GetMyShipsSecurity](../../models/operations/getmyshipssecurity.md) | :heavy_check_mark:                                                             | The security requirements to use for the request.                              |


### Response

**[*operations.GetMyShipsResponse](../../models/operations/getmyshipsresponse.md), error**


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

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetShipCooldownRequest](../../models/operations/getshipcooldownrequest.md)   | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `security`                                                                               | [operations.GetShipCooldownSecurity](../../models/operations/getshipcooldownsecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |


### Response

**[*operations.GetShipCooldownResponse](../../models/operations/getshipcooldownresponse.md), error**


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

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.GetShipNavRequest](../../models/operations/getshipnavrequest.md)   | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `security`                                                                     | [operations.GetShipNavSecurity](../../models/operations/getshipnavsecurity.md) | :heavy_check_mark:                                                             | The security requirements to use for the request.                              |


### Response

**[*operations.GetShipNavResponse](../../models/operations/getshipnavresponse.md), error**


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

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.JettisonRequest](../../models/operations/jettisonrequest.md)   | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `security`                                                                 | [operations.JettisonSecurity](../../models/operations/jettisonsecurity.md) | :heavy_check_mark:                                                         | The security requirements to use for the request.                          |


### Response

**[*operations.JettisonResponse](../../models/operations/jettisonresponse.md), error**


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

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.JumpShipRequest](../../models/operations/jumpshiprequest.md)   | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `security`                                                                 | [operations.JumpShipSecurity](../../models/operations/jumpshipsecurity.md) | :heavy_check_mark:                                                         | The security requirements to use for the request.                          |


### Response

**[*operations.JumpShipResponse](../../models/operations/jumpshipresponse.md), error**


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

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.NavigateShipRequest](../../models/operations/navigateshiprequest.md)   | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `security`                                                                         | [operations.NavigateShipSecurity](../../models/operations/navigateshipsecurity.md) | :heavy_check_mark:                                                                 | The security requirements to use for the request.                                  |


### Response

**[*operations.NavigateShipResponse](../../models/operations/navigateshipresponse.md), error**


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

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.NegotiateContractRequest](../../models/operations/negotiatecontractrequest.md)   | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.NegotiateContractSecurity](../../models/operations/negotiatecontractsecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |


### Response

**[*operations.NegotiateContractResponse](../../models/operations/negotiatecontractresponse.md), error**


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

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.OrbitShipRequest](../../models/operations/orbitshiprequest.md)   | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `security`                                                                   | [operations.OrbitShipSecurity](../../models/operations/orbitshipsecurity.md) | :heavy_check_mark:                                                           | The security requirements to use for the request.                            |


### Response

**[*operations.OrbitShipResponse](../../models/operations/orbitshipresponse.md), error**


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

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.PatchShipNavRequest](../../models/operations/patchshipnavrequest.md)   | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `security`                                                                         | [operations.PatchShipNavSecurity](../../models/operations/patchshipnavsecurity.md) | :heavy_check_mark:                                                                 | The security requirements to use for the request.                                  |


### Response

**[*operations.PatchShipNavResponse](../../models/operations/patchshipnavresponse.md), error**


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

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.PurchaseCargoRequest](../../models/operations/purchasecargorequest.md)   | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `security`                                                                           | [operations.PurchaseCargoSecurity](../../models/operations/purchasecargosecurity.md) | :heavy_check_mark:                                                                   | The security requirements to use for the request.                                    |


### Response

**[*operations.PurchaseCargoResponse](../../models/operations/purchasecargoresponse.md), error**


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

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.PurchaseShipRequestBody](../../models/operations/purchaseshiprequestbody.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `security`                                                                               | [operations.PurchaseShipSecurity](../../models/operations/purchaseshipsecurity.md)       | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |


### Response

**[*operations.PurchaseShipResponse](../../models/operations/purchaseshipresponse.md), error**


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

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.RefuelShipRequest](../../models/operations/refuelshiprequest.md)   | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `security`                                                                     | [operations.RefuelShipSecurity](../../models/operations/refuelshipsecurity.md) | :heavy_check_mark:                                                             | The security requirements to use for the request.                              |


### Response

**[*operations.RefuelShipResponse](../../models/operations/refuelshipresponse.md), error**


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

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.SellCargoRequest](../../models/operations/sellcargorequest.md)   | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `security`                                                                   | [operations.SellCargoSecurity](../../models/operations/sellcargosecurity.md) | :heavy_check_mark:                                                           | The security requirements to use for the request.                            |


### Response

**[*operations.SellCargoResponse](../../models/operations/sellcargoresponse.md), error**


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

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.ShipRefineRequest](../../models/operations/shiprefinerequest.md)   | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `security`                                                                     | [operations.ShipRefineSecurity](../../models/operations/shiprefinesecurity.md) | :heavy_check_mark:                                                             | The security requirements to use for the request.                              |


### Response

**[*operations.ShipRefineResponse](../../models/operations/shiprefineresponse.md), error**


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

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.TransferCargoRequest](../../models/operations/transfercargorequest.md)   | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `security`                                                                           | [operations.TransferCargoSecurity](../../models/operations/transfercargosecurity.md) | :heavy_check_mark:                                                                   | The security requirements to use for the request.                                    |


### Response

**[*operations.TransferCargoResponse](../../models/operations/transfercargoresponse.md), error**


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

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.WarpShipRequest](../../models/operations/warpshiprequest.md)   | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `security`                                                                 | [operations.WarpShipSecurity](../../models/operations/warpshipsecurity.md) | :heavy_check_mark:                                                         | The security requirements to use for the request.                          |


### Response

**[*operations.WarpShipResponse](../../models/operations/warpshipresponse.md), error**


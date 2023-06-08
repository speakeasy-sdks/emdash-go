# Systems

## Overview

Systems

### Available Operations

* [GetJumpGate](#getjumpgate) - Get Jump Gate
* [GetMarket](#getmarket) - Get Market
* [GetShipyard](#getshipyard) - Get Shipyard
* [GetSystem](#getsystem) - Get System
* [GetSystemWaypoints](#getsystemwaypoints) - List Waypoints
* [GetSystems](#getsystems) - List Systems
* [GetWaypoint](#getwaypoint) - Get Waypoint

## GetJumpGate

Get jump gate details for a waypoint.

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
    s := spacetraders.New(
        spacetraders.WithSecurity(shared.Security{
            AgentToken: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Systems.GetJumpGate(ctx, operations.GetJumpGateRequest{
        SystemSymbol: "ad",
        WaypointSymbol: "natus",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetJumpGate200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## GetMarket

Retrieve imports, exports and exchange data from a marketplace. Imports can be sold, exports can be purchased, and exchange goods can be purchased or sold. Send a ship to the waypoint to access trade good prices and recent transactions.

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
    s := spacetraders.New(
        spacetraders.WithSecurity(shared.Security{
            AgentToken: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Systems.GetMarket(ctx, operations.GetMarketRequest{
        SystemSymbol: "sed",
        WaypointSymbol: "iste",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetMarket200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## GetShipyard

Get the shipyard for a waypoint. Send a ship to the waypoint to access ships that are currently available for purchase and recent transactions.

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
    s := spacetraders.New(
        spacetraders.WithSecurity(shared.Security{
            AgentToken: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Systems.GetShipyard(ctx, operations.GetShipyardRequest{
        SystemSymbol: "dolor",
        WaypointSymbol: "natus",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetShipyard200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## GetSystem

Get the details of a system.

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
    s := spacetraders.New(
        spacetraders.WithSecurity(shared.Security{
            AgentToken: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Systems.GetSystem(ctx, operations.GetSystemRequest{
        SystemSymbol: "laboriosam",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetSystem200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## GetSystemWaypoints

Fetch all of the waypoints for a given system. System must be charted or a ship must be present to return waypoint details.

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
    res, err := s.Systems.GetSystemWaypoints(ctx, operations.GetSystemWaypointsRequest{
        Limit: spacetraders.Int64(943749),
        Page: spacetraders.Int64(902599),
        SystemSymbol: "fuga",
    }, operations.GetSystemWaypointsSecurity{
        AgentToken: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetSystemWaypoints200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## GetSystems

Return a list of all systems.

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
    s := spacetraders.New(
        spacetraders.WithSecurity(shared.Security{
            AgentToken: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Systems.GetSystems(ctx, operations.GetSystemsRequest{
        Limit: spacetraders.Int64(449950),
        Page: spacetraders.Int64(359508),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetSystems200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## GetWaypoint

View the details of a waypoint.

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
    s := spacetraders.New(
        spacetraders.WithSecurity(shared.Security{
            AgentToken: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Systems.GetWaypoint(ctx, operations.GetWaypointRequest{
        SystemSymbol: "iste",
        WaypointSymbol: "iure",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetWaypoint200ApplicationJSONObject != nil {
        // handle response
    }
}
```

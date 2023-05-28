# Factions

## Overview

Factions

### Available Operations

* [GetFaction](#getfaction) - Get Faction
* [GetFactions](#getfactions) - List Factions

## GetFaction

View the details of a faction.

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
            AgentToken: "YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Factions.GetFaction(ctx, operations.GetFactionRequest{
        FactionSymbol: "delectus",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetFaction200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## GetFactions

List all discovered factions in the game.

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
            AgentToken: "YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Factions.GetFactions(ctx, operations.GetFactionsRequest{
        Limit: spacetraders.Int64(272656),
        Page: spacetraders.Int64(383441),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetFactions200ApplicationJSONObject != nil {
        // handle response
    }
}
```

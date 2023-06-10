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
            AgentToken: "",
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

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.GetFactionRequest](../../models/operations/getfactionrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.GetFactionResponse](../../models/operations/getfactionresponse.md), error**


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
            AgentToken: "",
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

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.GetFactionsRequest](../../models/operations/getfactionsrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.GetFactionsResponse](../../models/operations/getfactionsresponse.md), error**


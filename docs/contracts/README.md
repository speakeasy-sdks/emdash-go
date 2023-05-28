# Contracts

## Overview

Contracts

### Available Operations

* [AcceptContract](#acceptcontract) - Accept Contract
* [DeliverContract](#delivercontract) - Deliver Contract
* [FulfillContract](#fulfillcontract) - Fulfill Contract
* [GetContract](#getcontract) - Get Contract
* [GetContracts](#getcontracts) - List Contracts

## AcceptContract

Accept a contract.

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
    res, err := s.Contracts.AcceptContract(ctx, operations.AcceptContractRequest{
        ContractID: "illum",
    }, operations.AcceptContractSecurity{
        AgentToken: "YOUR_BEARER_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AcceptContract200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## DeliverContract

Deliver cargo on a given contract.

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
    res, err := s.Contracts.DeliverContract(ctx, operations.DeliverContractRequest{
        RequestBody: &operations.DeliverContractRequestBody{
            ShipSymbol: "vel",
            TradeSymbol: "error",
            Units: 645894,
        },
        ContractID: "suscipit",
    }, operations.DeliverContractSecurity{
        AgentToken: "YOUR_BEARER_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeliverContract200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## FulfillContract

Fulfill a contract

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
    res, err := s.Contracts.FulfillContract(ctx, operations.FulfillContractRequest{
        ContractID: "iure",
    }, operations.FulfillContractSecurity{
        AgentToken: "YOUR_BEARER_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.FulfillContract200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## GetContract

Get the details of a contract by ID.

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
    res, err := s.Contracts.GetContract(ctx, operations.GetContractRequest{
        ContractID: "magnam",
    }, operations.GetContractSecurity{
        AgentToken: "YOUR_BEARER_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetContract200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## GetContracts

List all of your contracts.

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
    res, err := s.Contracts.GetContracts(ctx, operations.GetContractsRequest{
        Limit: spacetraders.Int64(891773),
        Page: spacetraders.Int64(56713),
    }, operations.GetContractsSecurity{
        AgentToken: "YOUR_BEARER_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetContracts200ApplicationJSONObject != nil {
        // handle response
    }
}
```

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
        AgentToken: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AcceptContract200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.AcceptContractRequest](../../models/operations/acceptcontractrequest.md)   | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `security`                                                                             | [operations.AcceptContractSecurity](../../models/operations/acceptcontractsecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |


### Response

**[*operations.AcceptContractResponse](../../models/operations/acceptcontractresponse.md), error**


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
        AgentToken: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeliverContract200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.DeliverContractRequest](../../models/operations/delivercontractrequest.md)   | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `security`                                                                               | [operations.DeliverContractSecurity](../../models/operations/delivercontractsecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |


### Response

**[*operations.DeliverContractResponse](../../models/operations/delivercontractresponse.md), error**


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
        AgentToken: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.FulfillContract200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.FulfillContractRequest](../../models/operations/fulfillcontractrequest.md)   | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `security`                                                                               | [operations.FulfillContractSecurity](../../models/operations/fulfillcontractsecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |


### Response

**[*operations.FulfillContractResponse](../../models/operations/fulfillcontractresponse.md), error**


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
        AgentToken: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetContract200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.GetContractRequest](../../models/operations/getcontractrequest.md)   | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `security`                                                                       | [operations.GetContractSecurity](../../models/operations/getcontractsecurity.md) | :heavy_check_mark:                                                               | The security requirements to use for the request.                                |


### Response

**[*operations.GetContractResponse](../../models/operations/getcontractresponse.md), error**


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
        AgentToken: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetContracts200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetContractsRequest](../../models/operations/getcontractsrequest.md)   | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `security`                                                                         | [operations.GetContractsSecurity](../../models/operations/getcontractssecurity.md) | :heavy_check_mark:                                                                 | The security requirements to use for the request.                                  |


### Response

**[*operations.GetContractsResponse](../../models/operations/getcontractsresponse.md), error**


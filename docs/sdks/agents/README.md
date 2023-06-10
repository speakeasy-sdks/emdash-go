# Agents

## Overview

Agents

### Available Operations

* [GetMyAgent](#getmyagent) - My Agent Details

## GetMyAgent

Fetch your agent's details.

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
    res, err := s.Agents.GetMyAgent(ctx, operations.GetMyAgentSecurity{
        AgentToken: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetMyAgent200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `security`                                                                     | [operations.GetMyAgentSecurity](../../models/operations/getmyagentsecurity.md) | :heavy_check_mark:                                                             | The security requirements to use for the request.                              |


### Response

**[*operations.GetMyAgentResponse](../../models/operations/getmyagentresponse.md), error**


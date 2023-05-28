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
        AgentToken: "YOUR_BEARER_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetMyAgent200ApplicationJSONObject != nil {
        // handle response
    }
}
```

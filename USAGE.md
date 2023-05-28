<!-- Start SDK Example Usage -->
```go
package main

import(
	"context"
	"log"
	"SpaceTraders"
)

func main() {
    s := spacetraders.New(
        spacetraders.WithSecurity(shared.Security{
            AgentToken: "YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.GetStatus(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetStatus200ApplicationJSONObject != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->
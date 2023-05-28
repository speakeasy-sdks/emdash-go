# SpaceTraders

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/emdash-go
```
<!-- End SDK Installation -->

## SDK Example Usage
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

<!-- Start SDK Available Operations -->
## Available Resources and Operations

### [SpaceTraders SDK](docs/spacetraders/README.md)

* [GetStatus](docs/spacetraders/README.md#getstatus) - Get Status
* [Register](docs/spacetraders/README.md#register) - Register New Agent

### [Agents](docs/agents/README.md)

* [GetMyAgent](docs/agents/README.md#getmyagent) - My Agent Details

### [Contracts](docs/contracts/README.md)

* [AcceptContract](docs/contracts/README.md#acceptcontract) - Accept Contract
* [DeliverContract](docs/contracts/README.md#delivercontract) - Deliver Contract
* [FulfillContract](docs/contracts/README.md#fulfillcontract) - Fulfill Contract
* [GetContract](docs/contracts/README.md#getcontract) - Get Contract
* [GetContracts](docs/contracts/README.md#getcontracts) - List Contracts

### [Factions](docs/factions/README.md)

* [GetFaction](docs/factions/README.md#getfaction) - Get Faction
* [GetFactions](docs/factions/README.md#getfactions) - List Factions

### [Fleet](docs/fleet/README.md)

* [CreateChart](docs/fleet/README.md#createchart) - Create Chart
* [CreateShipShipScan](docs/fleet/README.md#createshipshipscan) - Scan Ships
* [CreateShipSystemScan](docs/fleet/README.md#createshipsystemscan) - Scan Systems
* [CreateShipWaypointScan](docs/fleet/README.md#createshipwaypointscan) - Scan Waypoints
* [CreateSurvey](docs/fleet/README.md#createsurvey) - Create Survey
* [DockShip](docs/fleet/README.md#dockship) - Dock Ship
* [ExtractResources](docs/fleet/README.md#extractresources) - Extract Resources
* [GetMyShip](docs/fleet/README.md#getmyship) - Get Ship
* [GetMyShipCargo](docs/fleet/README.md#getmyshipcargo) - Get Ship Cargo
* [GetMyShips](docs/fleet/README.md#getmyships) - List Ships
* [GetShipCooldown](docs/fleet/README.md#getshipcooldown) - Get Ship Cooldown
* [GetShipNav](docs/fleet/README.md#getshipnav) - Get Ship Nav
* [Jettison](docs/fleet/README.md#jettison) - Jettison Cargo
* [JumpShip](docs/fleet/README.md#jumpship) - Jump Ship
* [NavigateShip](docs/fleet/README.md#navigateship) - Navigate Ship
* [NegotiateContract](docs/fleet/README.md#negotiatecontract) - Negotiate Contract
* [OrbitShip](docs/fleet/README.md#orbitship) - Orbit Ship
* [PatchShipNav](docs/fleet/README.md#patchshipnav) - Patch Ship Nav
* [PurchaseCargo](docs/fleet/README.md#purchasecargo) - Purchase Cargo
* [PurchaseShip](docs/fleet/README.md#purchaseship) - Purchase Ship
* [RefuelShip](docs/fleet/README.md#refuelship) - Refuel Ship
* [SellCargo](docs/fleet/README.md#sellcargo) - Sell Cargo
* [ShipRefine](docs/fleet/README.md#shiprefine) - Ship Refine
* [TransferCargo](docs/fleet/README.md#transfercargo) - Transfer Cargo
* [WarpShip](docs/fleet/README.md#warpship) - Warp Ship

### [Systems](docs/systems/README.md)

* [GetJumpGate](docs/systems/README.md#getjumpgate) - Get Jump Gate
* [GetMarket](docs/systems/README.md#getmarket) - Get Market
* [GetShipyard](docs/systems/README.md#getshipyard) - Get Shipyard
* [GetSystem](docs/systems/README.md#getsystem) - Get System
* [GetSystemWaypoints](docs/systems/README.md#getsystemwaypoints) - List Waypoints
* [GetSystems](docs/systems/README.md#getsystems) - List Systems
* [GetWaypoint](docs/systems/README.md#getwaypoint) - Get Waypoint
<!-- End SDK Available Operations -->

### Maturity

This SDK is in beta and therefore, we recommend pinning usage to a specific package version.
This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated and maintained programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release !

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)

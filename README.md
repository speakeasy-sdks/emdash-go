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
            AgentToken: "",
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

### [SpaceTraders SDK](docs/sdks/spacetraders/README.md)

* [GetStatus](docs/sdks/spacetraders/README.md#getstatus) - Get Status
* [Register](docs/sdks/spacetraders/README.md#register) - Register New Agent

### [Agents](docs/sdks/agents/README.md)

* [GetMyAgent](docs/sdks/agents/README.md#getmyagent) - My Agent Details

### [Contracts](docs/sdks/contracts/README.md)

* [AcceptContract](docs/sdks/contracts/README.md#acceptcontract) - Accept Contract
* [DeliverContract](docs/sdks/contracts/README.md#delivercontract) - Deliver Contract
* [FulfillContract](docs/sdks/contracts/README.md#fulfillcontract) - Fulfill Contract
* [GetContract](docs/sdks/contracts/README.md#getcontract) - Get Contract
* [GetContracts](docs/sdks/contracts/README.md#getcontracts) - List Contracts

### [Factions](docs/sdks/factions/README.md)

* [GetFaction](docs/sdks/factions/README.md#getfaction) - Get Faction
* [GetFactions](docs/sdks/factions/README.md#getfactions) - List Factions

### [Fleet](docs/sdks/fleet/README.md)

* [CreateChart](docs/sdks/fleet/README.md#createchart) - Create Chart
* [CreateShipShipScan](docs/sdks/fleet/README.md#createshipshipscan) - Scan Ships
* [CreateShipSystemScan](docs/sdks/fleet/README.md#createshipsystemscan) - Scan Systems
* [CreateShipWaypointScan](docs/sdks/fleet/README.md#createshipwaypointscan) - Scan Waypoints
* [CreateSurvey](docs/sdks/fleet/README.md#createsurvey) - Create Survey
* [DockShip](docs/sdks/fleet/README.md#dockship) - Dock Ship
* [ExtractResources](docs/sdks/fleet/README.md#extractresources) - Extract Resources
* [GetMyShip](docs/sdks/fleet/README.md#getmyship) - Get Ship
* [GetMyShipCargo](docs/sdks/fleet/README.md#getmyshipcargo) - Get Ship Cargo
* [GetMyShips](docs/sdks/fleet/README.md#getmyships) - List Ships
* [GetShipCooldown](docs/sdks/fleet/README.md#getshipcooldown) - Get Ship Cooldown
* [GetShipNav](docs/sdks/fleet/README.md#getshipnav) - Get Ship Nav
* [Jettison](docs/sdks/fleet/README.md#jettison) - Jettison Cargo
* [JumpShip](docs/sdks/fleet/README.md#jumpship) - Jump Ship
* [NavigateShip](docs/sdks/fleet/README.md#navigateship) - Navigate Ship
* [NegotiateContract](docs/sdks/fleet/README.md#negotiatecontract) - Negotiate Contract
* [OrbitShip](docs/sdks/fleet/README.md#orbitship) - Orbit Ship
* [PatchShipNav](docs/sdks/fleet/README.md#patchshipnav) - Patch Ship Nav
* [PurchaseCargo](docs/sdks/fleet/README.md#purchasecargo) - Purchase Cargo
* [PurchaseShip](docs/sdks/fleet/README.md#purchaseship) - Purchase Ship
* [RefuelShip](docs/sdks/fleet/README.md#refuelship) - Refuel Ship
* [SellCargo](docs/sdks/fleet/README.md#sellcargo) - Sell Cargo
* [ShipRefine](docs/sdks/fleet/README.md#shiprefine) - Ship Refine
* [TransferCargo](docs/sdks/fleet/README.md#transfercargo) - Transfer Cargo
* [WarpShip](docs/sdks/fleet/README.md#warpship) - Warp Ship

### [Systems](docs/sdks/systems/README.md)

* [GetJumpGate](docs/sdks/systems/README.md#getjumpgate) - Get Jump Gate
* [GetMarket](docs/sdks/systems/README.md#getmarket) - Get Market
* [GetShipyard](docs/sdks/systems/README.md#getshipyard) - Get Shipyard
* [GetSystem](docs/sdks/systems/README.md#getsystem) - Get System
* [GetSystemWaypoints](docs/sdks/systems/README.md#getsystemwaypoints) - List Waypoints
* [GetSystems](docs/sdks/systems/README.md#getsystems) - List Systems
* [GetWaypoint](docs/sdks/systems/README.md#getwaypoint) - Get Waypoint
<!-- End SDK Available Operations -->

### Maturity

This SDK is in beta and therefore, we recommend pinning usage to a specific package version.
This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated and maintained programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release !

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)

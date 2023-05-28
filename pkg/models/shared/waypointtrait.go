// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// WaypointTraitSymbol - The unique identifier of the trait.
type WaypointTraitSymbol string

const (
	WaypointTraitSymbolUncharted             WaypointTraitSymbol = "UNCHARTED"
	WaypointTraitSymbolMarketplace           WaypointTraitSymbol = "MARKETPLACE"
	WaypointTraitSymbolShipyard              WaypointTraitSymbol = "SHIPYARD"
	WaypointTraitSymbolOutpost               WaypointTraitSymbol = "OUTPOST"
	WaypointTraitSymbolScatteredSettlements  WaypointTraitSymbol = "SCATTERED_SETTLEMENTS"
	WaypointTraitSymbolSprawlingCities       WaypointTraitSymbol = "SPRAWLING_CITIES"
	WaypointTraitSymbolMegaStructures        WaypointTraitSymbol = "MEGA_STRUCTURES"
	WaypointTraitSymbolOvercrowded           WaypointTraitSymbol = "OVERCROWDED"
	WaypointTraitSymbolHighTech              WaypointTraitSymbol = "HIGH_TECH"
	WaypointTraitSymbolCorrupt               WaypointTraitSymbol = "CORRUPT"
	WaypointTraitSymbolBureaucratic          WaypointTraitSymbol = "BUREAUCRATIC"
	WaypointTraitSymbolTradingHub            WaypointTraitSymbol = "TRADING_HUB"
	WaypointTraitSymbolIndustrial            WaypointTraitSymbol = "INDUSTRIAL"
	WaypointTraitSymbolBlackMarket           WaypointTraitSymbol = "BLACK_MARKET"
	WaypointTraitSymbolResearchFacility      WaypointTraitSymbol = "RESEARCH_FACILITY"
	WaypointTraitSymbolMilitaryBase          WaypointTraitSymbol = "MILITARY_BASE"
	WaypointTraitSymbolSurveillanceOutpost   WaypointTraitSymbol = "SURVEILLANCE_OUTPOST"
	WaypointTraitSymbolExplorationOutpost    WaypointTraitSymbol = "EXPLORATION_OUTPOST"
	WaypointTraitSymbolMineralDeposits       WaypointTraitSymbol = "MINERAL_DEPOSITS"
	WaypointTraitSymbolCommonMetalDeposits   WaypointTraitSymbol = "COMMON_METAL_DEPOSITS"
	WaypointTraitSymbolPreciousMetalDeposits WaypointTraitSymbol = "PRECIOUS_METAL_DEPOSITS"
	WaypointTraitSymbolRareMetalDeposits     WaypointTraitSymbol = "RARE_METAL_DEPOSITS"
	WaypointTraitSymbolMethanePools          WaypointTraitSymbol = "METHANE_POOLS"
	WaypointTraitSymbolIceCrystals           WaypointTraitSymbol = "ICE_CRYSTALS"
	WaypointTraitSymbolExplosiveGases        WaypointTraitSymbol = "EXPLOSIVE_GASES"
	WaypointTraitSymbolStrongMagnetosphere   WaypointTraitSymbol = "STRONG_MAGNETOSPHERE"
	WaypointTraitSymbolVibrantAuroras        WaypointTraitSymbol = "VIBRANT_AURORAS"
	WaypointTraitSymbolSaltFlats             WaypointTraitSymbol = "SALT_FLATS"
	WaypointTraitSymbolCanyons               WaypointTraitSymbol = "CANYONS"
	WaypointTraitSymbolPerpetualDaylight     WaypointTraitSymbol = "PERPETUAL_DAYLIGHT"
	WaypointTraitSymbolPerpetualOvercast     WaypointTraitSymbol = "PERPETUAL_OVERCAST"
	WaypointTraitSymbolDrySeabeds            WaypointTraitSymbol = "DRY_SEABEDS"
	WaypointTraitSymbolMagmaSeas             WaypointTraitSymbol = "MAGMA_SEAS"
	WaypointTraitSymbolSupervolcanoes        WaypointTraitSymbol = "SUPERVOLCANOES"
	WaypointTraitSymbolAshClouds             WaypointTraitSymbol = "ASH_CLOUDS"
	WaypointTraitSymbolVastRuins             WaypointTraitSymbol = "VAST_RUINS"
	WaypointTraitSymbolMutatedFlora          WaypointTraitSymbol = "MUTATED_FLORA"
	WaypointTraitSymbolTerraformed           WaypointTraitSymbol = "TERRAFORMED"
	WaypointTraitSymbolExtremeTemperatures   WaypointTraitSymbol = "EXTREME_TEMPERATURES"
	WaypointTraitSymbolExtremePressure       WaypointTraitSymbol = "EXTREME_PRESSURE"
	WaypointTraitSymbolDiverseLife           WaypointTraitSymbol = "DIVERSE_LIFE"
	WaypointTraitSymbolScarceLife            WaypointTraitSymbol = "SCARCE_LIFE"
	WaypointTraitSymbolFossils               WaypointTraitSymbol = "FOSSILS"
	WaypointTraitSymbolWeakGravity           WaypointTraitSymbol = "WEAK_GRAVITY"
	WaypointTraitSymbolStrongGravity         WaypointTraitSymbol = "STRONG_GRAVITY"
	WaypointTraitSymbolCrushingGravity       WaypointTraitSymbol = "CRUSHING_GRAVITY"
	WaypointTraitSymbolToxicAtmosphere       WaypointTraitSymbol = "TOXIC_ATMOSPHERE"
	WaypointTraitSymbolCorrosiveAtmosphere   WaypointTraitSymbol = "CORROSIVE_ATMOSPHERE"
	WaypointTraitSymbolBreathableAtmosphere  WaypointTraitSymbol = "BREATHABLE_ATMOSPHERE"
	WaypointTraitSymbolJovian                WaypointTraitSymbol = "JOVIAN"
	WaypointTraitSymbolRocky                 WaypointTraitSymbol = "ROCKY"
	WaypointTraitSymbolVolcanic              WaypointTraitSymbol = "VOLCANIC"
	WaypointTraitSymbolFrozen                WaypointTraitSymbol = "FROZEN"
	WaypointTraitSymbolSwamp                 WaypointTraitSymbol = "SWAMP"
	WaypointTraitSymbolBarren                WaypointTraitSymbol = "BARREN"
	WaypointTraitSymbolTemperate             WaypointTraitSymbol = "TEMPERATE"
	WaypointTraitSymbolJungle                WaypointTraitSymbol = "JUNGLE"
	WaypointTraitSymbolOcean                 WaypointTraitSymbol = "OCEAN"
	WaypointTraitSymbolStripped              WaypointTraitSymbol = "STRIPPED"
)

func (e WaypointTraitSymbol) ToPointer() *WaypointTraitSymbol {
	return &e
}

func (e *WaypointTraitSymbol) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "UNCHARTED":
		fallthrough
	case "MARKETPLACE":
		fallthrough
	case "SHIPYARD":
		fallthrough
	case "OUTPOST":
		fallthrough
	case "SCATTERED_SETTLEMENTS":
		fallthrough
	case "SPRAWLING_CITIES":
		fallthrough
	case "MEGA_STRUCTURES":
		fallthrough
	case "OVERCROWDED":
		fallthrough
	case "HIGH_TECH":
		fallthrough
	case "CORRUPT":
		fallthrough
	case "BUREAUCRATIC":
		fallthrough
	case "TRADING_HUB":
		fallthrough
	case "INDUSTRIAL":
		fallthrough
	case "BLACK_MARKET":
		fallthrough
	case "RESEARCH_FACILITY":
		fallthrough
	case "MILITARY_BASE":
		fallthrough
	case "SURVEILLANCE_OUTPOST":
		fallthrough
	case "EXPLORATION_OUTPOST":
		fallthrough
	case "MINERAL_DEPOSITS":
		fallthrough
	case "COMMON_METAL_DEPOSITS":
		fallthrough
	case "PRECIOUS_METAL_DEPOSITS":
		fallthrough
	case "RARE_METAL_DEPOSITS":
		fallthrough
	case "METHANE_POOLS":
		fallthrough
	case "ICE_CRYSTALS":
		fallthrough
	case "EXPLOSIVE_GASES":
		fallthrough
	case "STRONG_MAGNETOSPHERE":
		fallthrough
	case "VIBRANT_AURORAS":
		fallthrough
	case "SALT_FLATS":
		fallthrough
	case "CANYONS":
		fallthrough
	case "PERPETUAL_DAYLIGHT":
		fallthrough
	case "PERPETUAL_OVERCAST":
		fallthrough
	case "DRY_SEABEDS":
		fallthrough
	case "MAGMA_SEAS":
		fallthrough
	case "SUPERVOLCANOES":
		fallthrough
	case "ASH_CLOUDS":
		fallthrough
	case "VAST_RUINS":
		fallthrough
	case "MUTATED_FLORA":
		fallthrough
	case "TERRAFORMED":
		fallthrough
	case "EXTREME_TEMPERATURES":
		fallthrough
	case "EXTREME_PRESSURE":
		fallthrough
	case "DIVERSE_LIFE":
		fallthrough
	case "SCARCE_LIFE":
		fallthrough
	case "FOSSILS":
		fallthrough
	case "WEAK_GRAVITY":
		fallthrough
	case "STRONG_GRAVITY":
		fallthrough
	case "CRUSHING_GRAVITY":
		fallthrough
	case "TOXIC_ATMOSPHERE":
		fallthrough
	case "CORROSIVE_ATMOSPHERE":
		fallthrough
	case "BREATHABLE_ATMOSPHERE":
		fallthrough
	case "JOVIAN":
		fallthrough
	case "ROCKY":
		fallthrough
	case "VOLCANIC":
		fallthrough
	case "FROZEN":
		fallthrough
	case "SWAMP":
		fallthrough
	case "BARREN":
		fallthrough
	case "TEMPERATE":
		fallthrough
	case "JUNGLE":
		fallthrough
	case "OCEAN":
		fallthrough
	case "STRIPPED":
		*e = WaypointTraitSymbol(v)
		return nil
	default:
		return fmt.Errorf("invalid value for WaypointTraitSymbol: %v", v)
	}
}

type WaypointTrait struct {
	// A description of the trait.
	Description string `json:"description"`
	// The name of the trait.
	Name string `json:"name"`
	// The unique identifier of the trait.
	Symbol WaypointTraitSymbol `json:"symbol"`
}

package main

import (
	"fmt"
	"math"
	"time"
)

// Telemetry holds the multi-variable environmental inputs for our predictive matrix
type Telemetry struct {
	DayOfWeek       time.Weekday
	HourOfDay       int
	IsPeakSeason    bool
	CurrentWeather  string  // "Blizzard", "Sunny", "Rain", "Clear"
	LiveUserTraffic int
}

// SystemState tracks the active infrastructure footprint and live deployment costs
type SystemState struct {
	ActiveNodes       int
	WholesaleSpendRate float64 // Simulated cost in USD per hour
}

// CalculatePredictiveRisk computes Risk Coefficient R based on real-time variables
func CalculatePredictiveRisk(t Telemetry) float64 {
	coefficient := 1.0

	// 1. Seasonal Operational Multipliers
	if t.IsPeakSeason {
		coefficient *= 2.5
	}

	// 2. Calendar and Time-of-Day Fluctuations
	if t.DayOfWeek == time.Saturday || t.DayOfWeek == time.Sunday {
		coefficient *= 1.8 // Weekend surge
	}
	if t.HourOfDay >= 8 && t.HourOfDay <= 11 {
		coefficient *= 1.4 // Morning booking rush window
	} else if t.HourOfDay >= 17 && t.HourOfDay <= 21 {
		coefficient *= 1.3 // Evening planning window
	}

	// 3. Live Weather Telemetry Dynamics
	switch t.CurrentWeather {
	case "Blizzard":
		coefficient *= 3.5 // Extreme lift for alpine properties
	case "Sunny":
		coefficient *= 1.6 // Boost for fair-weather leisure lines
	case "Rain":
		coefficient *= 0.4 // Traffic drop; immediately move into cost-mitigation mode
	default:
		coefficient *= 1.0
	}

	// 4. Live Traffic Volume Ingestion (Logarithmic scale prevents node oversaturating)
	if t.LiveUserTraffic > 0 {
		coefficient += math.Log10(float64(t.LiveUserTraffic))
	}

	return coefficient
}

// DeployInfrastructure Engine handles real-time node scaling and budget tracking
func DeployInfrastructure(riskFactor float64, state *SystemState) {
	// Dynamically compute target node matrix based on the calculated Risk Factor
	targetNodes := int(math.Ceil(riskFactor / 2.0))
	if targetNodes < 1 {
		targetNodes = 1 // Maintain permanent system availability baseline
	}

	fmt.Printf("[ENGINE LOG] Calculated Risk Coefficient (R): %.2f | Optimal Node Array: %d\n", riskFactor, targetNodes)

	// Evaluate discrepancies between live state and desired target state
	if targetNodes > state.ActiveNodes {
		nodesToSpawn := targetNodes - state.ActiveNodes
		fmt.Printf("[INFRASTRUCTURE UPGRADE] Provisioning %d wholesale compute nodes instantly...\n", nodesToSpawn)
		state.ActiveNodes = targetNodes
	} else if targetNodes < state.ActiveNodes {
		nodesToTerminate := state.ActiveNodes - targetNodes
		fmt.Printf("[COST MITIGATION] Terminating %d idle nodes to collapse infrastructure overhead...\n", nodesToTerminate)
		state.ActiveNodes = targetNodes
	}

	// Assume wholesale cloud bandwidth averages $0.024 per node/hour
	state.WholesaleSpendRate = float64(state.ActiveNodes) * 0.024
	fmt.Printf("[FINANCIAL REPORT] Predictive Spending Rate: $%.4f/hr ($%.2f/day baseline)\n\n", state.WholesaleSpendRate, state.WholesaleSpendRate*24)
}

func main() {
	fmt.Println("=== RUNNING ADVANCED MULTI-VARIABLE RESOURCE OPTIMIZER ===")
	
	// Stage our initial off-season baseline state
	fleetState := SystemState{ActiveNodes: 1, WholesaleSpendRate: 0.024}

	// Scenario A: A crisp Winter Saturday morning during a major snowfall event
	fmt.Println("--- TIME FRAME A: Winter Saturday Morning + Blizzard Warning ---")
	peakWinterWeekend := Telemetry{
		DayOfWeek:       time.Saturday,
		HourOfDay:       9, // Peak morning traffic
		IsPeakSeason:    true,
		CurrentWeather: "Blizzard",
		LiveUserTraffic: 25000,
	}
	riskA := CalculatePredictiveRisk(peakWinterWeekend)
	DeployInfrastructure(riskA, &fleetState)

	// Scenario B: A rainy Tuesday night during the same season (System sheds weight instantly)
	fmt.Println("--- TIME FRAME B: Winter Tuesday Night + Heavy Rain ---")
	rainyMidweekNight := Telemetry{
		DayOfWeek:       time.Tuesday,
		HourOfDay:       23, // Low-activity late night
		IsPeakSeason:    true,
		CurrentWeather: "Rain",
		LiveUserTraffic: 150,
	}
	riskB := CalculatePredictiveRisk(rainyMidweekNight)
	DeployInfrastructure(riskB, &fleetState)
}
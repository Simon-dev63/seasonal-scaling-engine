package api

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"

	"seasonal-scaling-engine/internal/database"
	"seasonal-scaling-engine/pkg/engine"
)

// IngestionHandler orchestrates the frontend CORS, the predictive math, and the database logging
func IngestionHandler(w http.ResponseWriter, r *http.Request) {
	// Enterprise CORS Implementation
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed. Use POST.", http.StatusMethodNotAllowed)
		return
	}

	var incomingData engine.Telemetry
	if err := json.NewDecoder(r.Body).Decode(&incomingData); err != nil {
		http.Error(w, "Invalid payload format.", http.StatusBadRequest)
		return
	}

	// Route data through the isolated mathematical engine
	riskFactor := engine.CalculatePredictiveRisk(incomingData)
	
	targetNodes := int(math.Ceil(riskFactor / 2.0))
	if targetNodes < 1 {
		targetNodes = 1
	}

	spendRate := float64(targetNodes) * 0.024
	action := fmt.Sprintf("System optimized to %d node(s).", targetNodes)

	response := engine.SystemResponse{
		RiskCoefficient: math.Round(riskFactor*100) / 100,
		TargetNodes:     targetNodes,
		HourlySpendRate: math.Round(spendRate*10000) / 10000,
		ActionTaken:     action,
	}

	// Trigger the database write asynchronously 
	go database.LogTransaction(incomingData, response)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	fmt.Printf("[NETWORK LOG] API request processed and routed to secure ledger.\n")
}
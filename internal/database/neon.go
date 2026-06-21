package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"seasonal-scaling-engine/pkg/engine"

	_ "github.com/lib/pq"
)

// DB exposes the active connection pool to the rest of the internal application
var DB *sql.DB

// Connect securely hooks into the Neon Cloud
func Connect() {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("[CRITICAL FAILURE] DATABASE_URL missing from .env vault.")
	}

	var err error
	DB, err = sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("[CRITICAL FAILURE] SQL initialization failed: %v", err)
	}
	
	// Enforce Enterprise Connection Pooling limits to prevent database crashing under high load
	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(5)

	if err = DB.Ping(); err != nil {
		log.Fatalf("[CRITICAL FAILURE] Failed to connect to Neon cloud: %v", err)
	}
	fmt.Println("[SYSTEM SECURE] Encrypted Database Layer Online.")
}

// LogTransaction handles the asynchronous writing of alpha metrics
func LogTransaction(t engine.Telemetry, r engine.SystemResponse) {
	telemetryQuery := `INSERT INTO telemetry_logs (day_of_week, current_weather, live_user_traffic, calculated_risk_coefficient) VALUES ($1, $2, $3, $4)`
	_, err := DB.Exec(telemetryQuery, t.DayOfWeek.String(), t.CurrentWeather, t.LiveUserTraffic, r.RiskCoefficient)
	if err != nil {
		log.Printf("[DB ERROR] Telemetry write failed: %v", err)
	}

	financialQuery := `INSERT INTO infrastructure_events (action_taken, nodes_active, wholesale_cost_per_hour) VALUES ($1, $2, $3)`
	_, err = DB.Exec(financialQuery, r.ActionTaken, r.TargetNodes, r.HourlySpendRate)
	if err != nil {
		log.Printf("[DB ERROR] Financial ledger write failed: %v", err)
	}
}
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"seasonal-scaling-engine/internal/api"
	"seasonal-scaling-engine/internal/database"

	"github.com/joho/godotenv"
)

func main() {
	// 1. Load the security vault
	if err := godotenv.Load(); err != nil {
		log.Println("[WARNING] No .env file found. Proceeding with system defaults.")
	}

	// 2. Ignite the database connection pool
	database.Connect()

	// 3. Register the API routing endpoints
	http.HandleFunc("/api/v1/optimize", api.IngestionHandler)

	// 4. Start the network listener
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("=== DOMAIN-DRIVEN INFRASTRUCTURE ENGINE ONLINE (PORT %s) ===\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// cmd/go-commerce/main.go

package main

import (
	"net/http"

	"github.com/Jagrit29/go-commerce/internal/api/handlers"
	"github.com/Jagrit29/go-commerce/internal/utils"
	"github.com/gorilla/mux"
)

func main() {
	// Initialize Logger
	logger := utils.NewLogger()

	// Initialize Router
	router := mux.NewRouter()

	// Initialize Handlers
	healthHandler := handlers.NewHealthHandler(logger)

	// Define API Routes
	router.HandleFunc("/health", healthHandler.HealthCheck).Methods("GET")

	// Start Server
	logger.Info("Go-Commerce server is running on port 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		logger.Fatal("Failed to start server:", err)
	}
}

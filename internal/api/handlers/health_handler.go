// internal/api/handlers/health_handler.go

package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Jagrit29/go-commerce/internal/models"
	"github.com/Jagrit29/go-commerce/internal/services"
	"github.com/Jagrit29/go-commerce/internal/utils"
)

type HealthHandler struct {
	service *services.HealthService
	logger  *utils.Logger
}

func NewHealthHandler(logger *utils.Logger) *HealthHandler {
	service := services.NewHealthService()
	return &HealthHandler{
		service: service,
		logger:  logger,
	}
}

func (h *HealthHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	status := h.service.CheckHealth()

	response := models.HealthResponse{
		Status: status,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		h.logger.Info("Failed to write response:")
	}
}

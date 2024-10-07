// internal/services/health_service.go

package services

type HealthService struct{}

func NewHealthService() *HealthService {
	return &HealthService{}
}

func (hs *HealthService) CheckHealth() string {
	return "OK"
}

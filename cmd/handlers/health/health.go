package health

import (
	"net/http"

	"github.com/juanMaAV92/go-server-template/internal/services/health"
	"github.com/labstack/echo/v4"
)

type Service interface {
	Check() health.HealthResponse
}

type Handler struct {
	Service
}

func NewHandler(s Service) *Handler {
	return &Handler{
		Service: s,
	}
}

// @Summary Health Check
// @Description Returns the health status
// @Tags health
// @Success 200 {object} health.HealthResponse
// @Router /health-check [get]
func (h *Handler) Check(ctx echo.Context) error {
	response := h.Service.Check()
	return ctx.JSON(http.StatusOK, response)
}

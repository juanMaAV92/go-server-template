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

func (h *Handler) Check(ctx echo.Context) error {
	response := h.Service.Check()
	return ctx.JSON(http.StatusOK, response)
}

package cmd

import (
	"github.com/juanMaAV92/go-server-template/cmd/handlers/health"
	utilsMiddleware "github.com/juanMaAV92/go-utils/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	apiV1Group      = "/v1"
	healthCheckPath = "/health-check"
)

type HealthHandler interface {
	Check(ctx echo.Context) error
}

type handlers struct {
	health HealthHandler
}

func configRoutes(inst *Instance, services *services) {
	configMiddleware(inst)

	handlers := initializeHandlers(services)

	baseGroup := inst.Server.Group(inst.config.ServerName)
	baseGroup.GET(healthCheckPath, handlers.health.Check)

	// v1 := baseGroup.Group(apiV1Group)
	// configureV1Routes(v1, handlers)
}

func initializeHandlers(services *services) *handlers {
	healthHandler := health.NewHandler(services.healthService)

	return &handlers{
		health: healthHandler,
	}
}

// func configureV1Routes(v1 *echo.Group, h *handlers) {
// }

func configMiddleware(inst *Instance) {
	inst.Server.Use(middleware.Recover())

	inst.Server.Use(utilsMiddleware.Tracing(inst.config.ServerName))
	inst.Server.Use(utilsMiddleware.TraceId())
	inst.Server.Use(utilsMiddleware.Logging(inst.Logger))
}

package cmd

import (
	healthHandler "github.com/juanMaAV92/go-server-template/cmd/handlers/health"
	"github.com/juanMaAV92/go-server-template/internal/services/health"
	"github.com/juanMaAV92/go-server-template/platform/config"
	"github.com/juanMaAV92/go-utils/env"
	"github.com/juanMaAV92/go-utils/errors"
	"github.com/juanMaAV92/go-utils/log"
	"github.com/juanMaAV92/go-utils/platform/server"
	"github.com/labstack/echo/v4"
	"github.com/shopspring/decimal"
)

type Instance struct {
	*server.Server
	log.Logger
	config *config.Config
}

type services struct {
	healthService healthHandler.Service
}

func NewServer(cfg *config.Config, logger log.Logger) (*Instance, error) {
	instance := echo.New()
	instance.HideBanner = true
	instance.HTTPErrorHandler = errors.CustomHTTPErrorHandler
	instance.Debug = cfg.Environment == env.LocalEnvironment
	decimal.MarshalJSONWithoutQuotes = true

	return &Instance{
		Server: &server.Server{
			Echo: instance,
		},
		Logger: logger,
		config: cfg,
	}, nil
}

func (inst Instance) initServices() (*services, error) {
	healthService := health.NewService()

	return &services{
		healthService: healthService,
	}, nil
}

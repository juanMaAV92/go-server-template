package cmd

import (
	"context"
	"os"

	"github.com/juanMaAV92/go-server-template/platform/config"
	"github.com/juanMaAV92/go-utils/env"
	"github.com/juanMaAV92/go-utils/log"
	"github.com/juanMaAV92/go-utils/tracing"
)

const (
	initServerStep = "init_server_step"
	shutdownStep   = "shutdown_server_step"

	errStartingMsg         = "Error starting server"
	errStartingServicesMsg = "Error starting services"
	errRunningMsg          = "Error while running server"
	errInitTracingMsg      = "Error initializing tracing"
)

const (
	exitCodeFailStartingServer = iota + 1
	exitPodeFailStartingServices
	exitCodeFailRunningServer
	exitCodeFailInitTracing
)

func Start() {
	ctx := context.Background()
	environment := env.GetEnviroment()

	cfg, err := config.Load(environment)
	if err != nil {
		panic("Failed to load configuration: " + err.Error())
	}

	logger := log.New(config.MicroserviceName, log.WithLevel(log.InfoLevel))

	tracingShutdown, err := StartInstrumentation(cfg)
	if err != nil {
		logger.Fatal(ctx, errInitTracingMsg, err.Error())
		os.Exit(exitCodeFailInitTracing)
	}
	defer StopInstrumentation(ctx, tracingShutdown)

	srv, err := NewServer(cfg, logger)
	if err != nil {
		logger.Fatal(ctx, errStartingMsg, err.Error())
		os.Exit(exitCodeFailStartingServer)
	}

	svc, err := srv.initServices()
	if err != nil {
		logger.Fatal(ctx, errStartingServicesMsg, err.Error())
		os.Exit(exitPodeFailStartingServices)
	}

	configRoutes(srv, svc)

	errC := srv.Run(cfg.Port, cfg.GracefulTime)

	logger.Info(ctx, "Server started successfully", "")

	if errS := <-errC; errS != nil {
		logger.Fatal(ctx, errRunningMsg, errS.Error())
		os.Exit(exitCodeFailRunningServer)
	}
}

func StartInstrumentation(config *config.Config) (func(context.Context) error, error) {
	tracingConfig := tracing.NewTracingConfig(config.ServerName, config.Telemetry.OTLPEndpoint, config.Environment)
	return tracing.InitTracing(tracingConfig)
}

func StopInstrumentation(ctx context.Context, shutdown func(context.Context) error) {
	shutdown(ctx)
}

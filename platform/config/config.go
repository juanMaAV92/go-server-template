package config

import (
	"time"

	"github.com/juanMaAV92/go-utils/env"
	platform "github.com/juanMaAV92/go-utils/platform/config"
)

const (
	MicroserviceName = "go-server-template"
	ServiceVersion   = "1.0.0"
)

var localConfig = Config{
	BasicConfig: &platform.BasicConfig{
		Port:         "8080",
		GracefulTime: 5 * time.Minute,
		Environment:  "local",
		ServerName:   MicroserviceName,
	},
	Telemetry: platform.TelemetryConfig{
		OTLPEndpoint: "localhost:4318",
	},
}

func deployConfig() Config {
	return Config{
		BasicConfig: platform.GetBasicServerConfig(MicroserviceName),
		Telemetry: platform.TelemetryConfig{
			OTLPEndpoint: env.GetEnv(env.OTLP_ENDPOINT),
		},
	}
}

func Load(enviroment string) (*Config, error) {
	var config Config
	if enviroment == env.LocalEnvironment {
		config = localConfig
	} else {
		config = deployConfig()
	}

	return &config, nil
}

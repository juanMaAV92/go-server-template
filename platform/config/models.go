package config

import (
	platform "github.com/juanMaAV92/go-utils/platform/config"
)

type Config struct {
	*platform.BasicConfig
	Telemetry platform.TelemetryConfig
}

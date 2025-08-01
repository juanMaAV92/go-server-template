# go-server-template

[![Go Version](https://img.shields.io/badge/Go-1.24+-blue.svg)](https://golang.org/dl/)

**Base template** to create microservices in Go with full observability, hexagonal architecture, and CI/CD configuration.

> ⚠️ **Important**: This is a project template. After cloning, you must change the name `go-server-template` to your new project name in all files (go.mod, imports, container names, etc.).

```bash
curl --location 'http://localhost:8080/go-server-template/health-check'
```

## 📋 Table of Contents

1. [🎯 Features](#-features)
2. [🏗️ Architecture](#-architecture)
3. [📂 Project Structure](#-project-structure)
4. [🔧 Customization](#-customization)
5. [⚙️ Configuration](#-configuration)
   - [Environment Variables](#environment-variables)
   - [GitHub Actions](#github-actions)
6. [🚀 Quick Start](#-quick-start)
7. [⚙️ Development](#-development)
8. [🔍 Observability](#-observability)
   - [Tool Stack](#tool-stack)
   - [Access to Tools](#access-to-tools)
   - [Querying Logs](#querying-logs)

## 🎯 Features

- **Structured Logging:** [zerolog](https://github.com/rs/zerolog) with trace correlation
- **Distributed Tracing:** Integrated OpenTelemetry
- **HTTP Framework:** [Echo](https://echo.labstack.com/) with custom middleware
- **Environment Configuration:** Using [go-utils](https://github.com/juanMaAV92/go-utils)
- **Testing:** Unit and integration tests with coverage
- **CI/CD:** Configured GitHub Actions
- **Hexagonal Architecture:** Clear separation of responsibilities
- **Observability Stack:** Jaeger, Loki, Grafana, OTel Collector

## 🏗️ Architecture

Implements a simplified hexagonal architecture:

```
┌─────────────────┐    ┌──────────────────┐    ┌─────────────────┐
│   Application   │    │     Domain       │    │ Infrastructure  │
│                 │    │                  │    │                 │
│ • HTTP Handlers │───▶│ • Services       │◀───│ • Configuration │
│ • Routing       │    │ • Domain Logic   │    │ • Database      │
│ • Middleware    │    │ • Models         │    │ • External Svcs │
└─────────────────┘    └──────────────────┘    └─────────────────┘
     cmd/              internal/services/       platform/ +
                       internal/domain/         services/
```

**Benefits:**
- Easy testing through mocking
- Flexibility to change implementations
- Maintainable and scalable code

## 📂 Project Structure

```
.
├── cmd/                        # 🚀 Application Layer
│   ├── main.go                 # Main configuration and startup
│   ├── server.go               # HTTP server configuration
│   ├── routing.go              # Route and middleware definition
│   └── handlers/               # HTTP handlers by domain
│       └── health/             # Health endpoints
├── internal/                   # 🧠 Domain Layer
│   ├── services/               # Application services
│   │   └── health/             # Health check logic
│   │       ├── health.go       # Service implementation
│   │       └── models.go       # Service models
│   └── domain/                 # Entities and business logic
│       └── [future domains]    # Business-specific domains
├── platform/                   # ⚙️ Infrastructure Layer
│   └── config/                 # Environment configuration
│       ├── config.go           # Configuration loading
│       └── models.go           # Configuration models
├── tests/                      # 🧪 Tests
│   ├── healthCheck_test.go     # Integration tests
│   └── helpers/                # Testing utilities
├── .github/workflows/          # 🔄 CI/CD
│   ├── test.yml                # Test pipeline
│   └── docker-publish.yml      # Build and publish
├── .vscode/                    # 🛠️ IDE configuration
├── Dockerfile                  # 🐳 Docker configuration
└── main.go                     # Entry point
```

## 🔧 Customization

After cloning this template, follow these steps to customize your project:

### 1. Change the Project Name

Replace `go-server-template` with your project name in the following files:

**📁 Files to modify:**
```bash
# 1. go.mod - Change the module name
module github.com/your-user/your-new-project

# 2. All imports in .go files
github.com/juanMaAV92/go-server-template → github.com/your-user/your-new-project

# 3. platform/config/config.go - Change MicroserviceName
const MicroserviceName = "your-new-project-ms"

# 4. Dockerfile - Change the microservice name in the health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD curl -f http://localhost:8080/your-new-project-ms/health-check || exit 1
```

## ⚙️ Configuration

### Environment Variables

The project uses `go-utils` for configuration. The following variables are available:

| Variable | Description | Default Value | Required |
|----------|-------------|---------------|----------|
| `ENVIRONMENT` | Execution environment (`local`, `development`, `staging`, `production`) | `local` | No |
| `PORT` | HTTP server port | `8080` | No |
| `GRACEFUL_TIME` | Graceful shutdown time (seconds) | `300` | No |
| `OTLP_ENDPOINT` | OpenTelemetry Collector endpoint | `localhost:4318` | No |


### GitHub Actions

To configure CI/CD workflows:

1. **Go to Settings → Secrets and variables → Actions**
2. **Set the following Repository secrets:**

| Secret Name | Description | Example |
|-------------|-------------|---------|
| `GITHUB_TOKEN` | Token for repository access during Docker build and for private repositories | `ghp_xxxxx` |

## 🚀 Quick Start

### Local Execution
```bash
# Run directly
go run main.go

# Or build and run
go build -o bin/go-server-template main.go
./bin/go-server-template
```


## ⚙️ Development

### Tests
```bash
# Run all tests with coverage
go test ./... -coverprofile=coverage.out -coverpkg=./...

# View coverage report
go tool cover -html=coverage.out
```


### Build
```bash
# Production build
go build -o bin/go-server-template main.go
```

## 🔍 Observability

### Tool Stack

The project includes a complete observability stack:

1. **OpenTelemetry Collector**
   - Receives telemetry from the application
   - Processes and routes to specific backends
   - Flexible and decoupled configuration

# go-server-template

[![Go Version](https://img.shields.io/badge/Go-1.24+-blue.svg)](https://golang.org/dl/)


**Plantilla base** para crear microservicios en Go con observabilidad completa, arquitectura hexagonal y configuración para CI/CD.

> ⚠️ **Importante**: Esta es una plantilla de proyecto. Después de clonar, debes cambiar el nombre `go-server-template` por el nombre de tu nuevo proyecto en todos los archivos (go.mod, imports, nombres de contenedores, etc.).


```bash
curl --location 'http://localhost:8080/go-server-template/health-check'
```


## 📋 Tabla de Contenido

1. [🎯 Características](#-características)
2. [🏗️ Arquitectura](#-arquitectura)
3. [📂 Estructura del Proyecto](#-estructura-del-proyecto)
4. [🔧 Personalización](#-personalización)
5. [⚙️ Configuración](#-configuración)
   - [Variables de Entorno](#variables-de-entorno)
   - [GitHub Actions](#github-actions)
6. [🚀 Inicio Rápido](#-inicio-rápido)
7. [⚙️ Desarrollo](#-desarrollo)
8. [🔍 Observabilidad](#-observabilidad)
   - [Stack de Herramientas](#stack-de-herramientas)
   - [Acceso a Herramientas](#acceso-a-herramientas)
   - [Consultando Logs](#consultando-logs)

## 🎯 Características

- **Logging Estructurado:** [zerolog](https://github.com/rs/zerolog) con correlación de trazas
- **Tracing Distribuido:** OpenTelemetry integrado
- **HTTP Framework:** [Echo](https://echo.labstack.com/) con middleware personalizado
- **Configuración por Entorno:** Usando [go-utils](https://github.com/juanMaAV92/go-utils)
- **Testing:** Tests unitarios e integración con coverage
- **CI/CD:** GitHub Actions configurado
- **Arquitectura Hexagonal:** Separación clara de responsabilidades
- **Stack de Observabilidad:** Jaeger, Loki, Grafana, OTel Collector

## 🏗️ Arquitectura

Implementa una arquitectura hexagonal simplificada:

```
┌─────────────────┐    ┌──────────────────┐    ┌─────────────────┐
│   Aplicación    │    │     Dominio      │    │ Infraestructura │
│                 │    │                  │    │                 │
│ • HTTP Handlers │───▶│ • Services       │◀───│ • Configuración │
│ • Routing       │    │ • Domain Logic   │    │ • Base de Datos │
│ • Middleware    │    │ • Models         │    │ • Servicios Ext │
└─────────────────┘    └──────────────────┘    └─────────────────┘
     cmd/              internal/services/       platform/ +
                       internal/domain/         services/
```

**Beneficios:**
- Testing fácil mediante mocking
- Flexibilidad para cambiar implementaciones
- Código mantenible y escalable

## 📂 Estructura del Proyecto

```
.
├── cmd/                        # 🚀 Capa de Aplicación
│   ├── main.go                 # Configuración principal y arranque
│   ├── server.go               # Configuración del servidor HTTP
│   ├── routing.go              # Definición de rutas y middleware
│   └── handlers/               # Handlers HTTP por dominio
│       └── health/             # Endpoints de salud
├── internal/                   # 🧠 Capa de Dominio
│   ├── services/               # Servicios de aplicación
│   │   └── health/             # Lógica de health checks
│   │       ├── health.go       # Implementación del servicio
│   │       └── models.go       # Modelos del servicio
│   └── domain/                 # Entidades y lógica de negocio
│       └── [future domains]    # Dominios específicos del negocio
├── platform/                   # ⚙️ Capa de Infraestructura
│   └── config/                 # Configuración por entorno
│       ├── config.go           # Carga de configuración
│       └── models.go           # Modelos de configuración
├── tests/                      # 🧪 Tests
│   ├── healthCheck_test.go     # Tests de integración
│   └── helpers/                # Utilidades para testing
├── .github/workflows/          # 🔄 CI/CD
│   ├── test.yml                # Pipeline de tests
│   └── docker-publish.yml      # Build y publicación
├── .vscode/                    # 🛠️ Configuración IDE
├── Dockerfile                  # 🐳 Configuración Docker
└── main.go                     # Punto de entrada
```

## 🔧 Personalización

Después de clonar esta plantilla, sigue estos pasos para personalizar tu proyecto:

### 1. Cambiar el Nombre del Proyecto

Reemplaza `go-server-template` por el nombre de tu proyecto en los siguientes archivos:

**📁 Archivos a modificar:**
```bash
# 1. go.mod - Cambiar el nombre del módulo
module github.com/tu-usuario/tu-nuevo-proyecto

# 2. Todos los imports en archivos .go
github.com/juanMaAV92/go-server-template → github.com/tu-usuario/tu-nuevo-proyecto

# 3. platform/config/config.go - Cambiar MicroserviceName
const MicroserviceName = "tu-nuevo-proyecto-ms"

# 4. Dockerfile - Cambiar el nombre del microservicio en el health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD curl -f http://localhost:8080/tu-nuevo-proyecto-ms/health-check || exit 1
```

## ⚙️ Configuración

### Variables de Entorno

El proyecto utiliza `go-utils` para configuración. Las siguientes variables están disponibles:

| Variable | Descripción | Valor por Defecto | Requerida |
|----------|-------------|-------------------|-----------|
| `ENVIRONMENT` | Entorno de ejecución (`local`, `development`, `staging`, `production`) | `local` | No |
| `PORT` | Puerto del servidor HTTP | `8080` | No |
| `GRACEFUL_TIME` | Tiempo de gracia para shutdown (segundos) | `300` | No |
| `OTLP_ENDPOINT` | Endpoint del OpenTelemetry Collector | `localhost:4318` | No |


### GitHub Actions

Para configurar los workflows de CI/CD:

1. **Ve a Settings → Secrets and variables → Actions**
2. **Configura los siguientes Repository secrets:**

| Secret Name | Descripción | Ejemplo |
|-------------|-------------|---------|
| `GITHUB_TOKEN` | Token para acceso a repositorios durante build de Docker y para repositorios privados | `ghp_xxxxx` |

## 🚀 Inicio Rápido

### Ejecución Local
```bash
# Ejecutar directamente
go run main.go

# O compilar y ejecutar
go build -o bin/go-server-template main.go
./bin/go-server-template
```


## ⚙️ Desarrollo

### Tests
```bash
# Ejecutar todos los tests con coverage
go test ./... -coverprofile=coverage.out -coverpkg=./...

# Ver reporte de coverage
go tool cover -html=coverage.out
```


### Build
```bash
# Build para producción
go build -o bin/go-server-template main.go
```

## 🔍 Observabilidad

### Stack de Herramientas

El proyecto incluye un stack completo de observabilidad:

1. **OpenTelemetry Collector**
   - Recibe telemetría de la aplicación
   - Procesa y enruta a backends específicos
   - Configuración flexible y desacoplada




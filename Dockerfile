# ================================
# STAGE 1: Build Stage
# ================================
FROM --platform=$BUILDPLATFORM golang:1.24-alpine AS builder

# Install git and ca-certificates (needed for downloading private modules)
RUN apk add --no-cache git ca-certificates tzdata

# Set working directory
WORKDIR /app

# if you want to configure Go for private modules
# ENV GOPRIVATE=github.com/juanMaAV92/*

# Configure git to use token-based authentication for GitHub
# Uses BuildKit secret (more secure than ARG)
RUN --mount=type=secret,id=github_token,env=GH_TOKEN_REPOSITORY \
    if [ -n "$GH_TOKEN_REPOSITORY" ]; then \
        git config --global url."https://${GH_TOKEN_REPOSITORY}@github.com/".insteadOf "https://github.com/" ; \
    fi

# Copy go mod files first (for better caching)
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build arguments for cross-compilation
ARG TARGETOS
ARG TARGETARCH

# Build the application
# CGO_ENABLED=0: Statically linked binary
# GOOS and GOARCH: Target platform from build args
# -ldflags="-w -s": Strip debugging info to reduce size
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build \
    -ldflags="-w -s" \
    -o main \
    ./main.go

# ================================
# STAGE 2: Production Stage
# ================================
FROM --platform=$TARGETPLATFORM alpine:latest

# Install necessary packages: ca-certificates, tzdata, and curl for health checks
RUN apk --no-cache add ca-certificates tzdata curl

# Create a non-root user for security
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# Set working directory
WORKDIR /app

# Copy the binary from builder stage
COPY --from=builder /app/main .

# Copy migration files (needed for database setup)
COPY --from=builder /app/migration ./migration

# Change ownership to non-root user
RUN chown -R appuser:appgroup /app

# Switch to non-root user
USER appuser

# Expose port 8080
EXPOSE 8080

# Health check using curl (more common and reliable)
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD curl -f http://localhost:8080/go-server-template/health-check || exit 1

# Run the application
CMD ["./main"] 
# Token Introspection Endpoint

[![Go Version](https://img.shields.io/badge/go-1.24-blue)]()
[![License: MIT](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](LICENSE)


A Go microservice implementing OAuth2 Token Introspection (RFC 7662)

## Overview

This microservice provides a `/introspect` endpoint that accepts an OAuth2 token (JWT) and returns a JSON response conforming to RFC 7662. Incoming JWTs are validated against a JWKS endpoint (e.g., https://issuer.example.com/certs), and public keys are cached in memory for a configurable interval.

Health checks (`/health/liveness`, `/health/readiness`) are provided for container orchestration readiness and liveness probes. The service uses slog for uniform, structured logging in JSON format.

## Features

- **Token Introspection (POST /introspect)**: Validates token signature, expiration, and returns structured introspection response.
- **JWKS Caching**: Fetches public keys from a JWKS URL and caches them for a configurable duration.
- **Health Endpoints**: 
  - `/health/liveness`: Returns HTTP 200 if the service is up.
  - `/health/readiness`: Returns HTTP 200 if the service is ready to receive traffic.
- **Graceful Shutdown**: Listens for SIGINT/SIGTERM and shuts down the HTTP server cleanly.
- **Structured Logging**: Uses slog (JSON) for all log statements, with standard fields (timestamp, level, message, key/value pairs).

## Installation 
```bash
# 1. Clone the repository
git clone https://github.com/usuario/token-introspect.git
cd token-introspect

# 2. Fetch dependencies
go mod tidy
```

## Running the Service

1. **Run directly** (for development):
```bash
go run ./cmd/introspect/main.go
```

2. **Build a binary** and execute:
```bash
go build -o bin/introspect ./cmd/introspect
./bin/introspect
```


# GoPick: Passenger Service

The passenger-facing backend service for GoPick, a ride-hailing platform.
Handles passenger identity, profiles and ride requests.

Part of the GoPick system:
- **go-passenger** (this repo) : passenger API
- **go-driver** : driver API
- **go-infra** : shared infrastructure and architecture docs

## Tech stack

- **Go** : service implementation
- **PostgreSQL** : primary datastore

## Configuration

All configuration comes from environment variables. For local development, copy the template and edit it:

    cp .env.example .env

| Variable | Default | Description |
| --- | --- | --- |
| `PORT` | `8080` | Port the HTTP server listens on |
| `APP_ENV` | `development` | Runtime environment: development, staging or production |

Real environment variables take precedence over values in `.env`. The `.env` file is a local development convenience only. In staging and production, configuration is supplied by the platform or a secret manager, never by a file in the repo.

## API

| Method | Path | Description |
| --- | --- | --- |
| `GET` | `/health` | Liveness check. Returns 200 with {"status":"ok"} |

## Shutdown behaviour

The service handles SIGINT and SIGTERM. On either signal it stops accepting new
connections and waits up to 10 seconds for in-flight requests to finish before
exiting. Deployment platforms must allow a termination grace period longer than
10 seconds, or shutdown will be cut short by SIGKILL.

Request timeouts: 5s for headers, 15s read, 15s write, 60s idle.

## Status

Early development. Nothing is stable yet.

## Getting started

Requirements: Go 1.26+, PostgreSQL 16+

```bash
git clone git@github.com:Olamilekan-12/go-passenger.git
cd go-passenger
cp .env.example .env   # then fill in your own values
go run ./cmd/api
```

## Tests

    go test ./...                 # run all tests
    go test ./... -v              # with per-test output
    go test ./... -cover          # with coverage summary

## License

MIT


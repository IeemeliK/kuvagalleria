# Kuvagalleria

Image gallery application built with Go, HTMX, and TailwindCSS.

## Prerequisites

| Tool | Version | Install |
|------|---------|---------|
| Go | 1.25+ | [go.dev](https://go.dev/dl/) |
| Docker | Latest | [docker.com](https://docs.docker.com/engine/install/) |
| Bun | Latest | `curl -fsSL https://bun.sh/install \| bash` |
| goose | Latest | `go install github.com/pressly/goose/v3/cmd/goose@latest` |

Optional tools:
- `golangci-lint` — `go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest`
- `air` (hot-reload in dev container — already in Dockerfile.dev)

## Getting Started

```bash
# 1. Clone and enter
git clone <repo> && cd kuvagalleria

# 2. Environment config
cp configs/.env.example .env
# Edit .env — at minimum change COOKIESTORE_SECRET:
#   openssl rand -base64 32

# 3. Start dev environment
make dev_up
```

This starts Postgres + pgAdmin + the app, and runs pending migrations automatically.

Open http://localhost:8080.

## Migrations

SQL files live in `migrations/`. Commands load env vars from `.env` automatically.

```bash
make migrate-status        # check current version
make migrate-create name=add_images  # new migration
make migrate-up            # apply pending
make migrate-down          # rollback last
make migrate-reset         # rollback all
```

## Commands

| Command | Description |
|---------|-------------|
| `make dev_up` | Start Postgres + app in Docker (runs migrations) |
| `make dev_down` | Stop all containers |
| `make build` | Production binary to `./bin/app` |
| `make test` | Run all tests with race detection |
| `make lint` | Run golangci-lint |
| `make fmt` | Format Go source |
| `make watch_css` | Watch and rebuild Tailwind CSS |
| `make dev_css` | Build dev CSS once |

## Production Build

```bash
make build
./bin/app
```

## Project Structure

```
├── cmd/server/main.go      # Entry point
├── internal/               # Go packages
│   ├── api/                # HTTP handlers
│   ├── config/             # Env config
│   ├── middleware/          # Auth + logging
│   ├── repository/         # Database layer
│   ├── service/            # Business logic
│   └── templates/          # Template cache
├── web/                    # Frontend assets
│   ├── static/             # Built CSS (embedded)
│   └── templates/          # HTML templates
├── migrations/             # DB migrations (goose)
└── deployments/            # Docker + Compose
```

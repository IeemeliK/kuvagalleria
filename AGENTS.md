# Agent Guidelines

## Architecture & Tech Stack
- **Backend**: Go 1.25+ stdlib `http.NewServeMux` routing. No Gin/Echo.
- **Database**: PostgreSQL via `database/sql` + `pgx/v5` driver.
- **Auth**: `gorilla/sessions` cookie store, bcrypt password verification.
- **Frontend**: HTMX, TailwindCSS v4 + DaisyUI v5.
- **Templates**: Go `html/template` with custom cache (`internal/templates/`). HTML in `web/templates/`, layouts and pages subdirectories. Embedded via `web/embed.go`.
- **Static assets**: Built CSS at `web/static/css/output.css` (embedded, served at `/static/`). Source CSS at `assets/input.css`.

## Project Layout
```
cmd/server/main.go          # Entry point
internal/
  api/                      # HTTP handlers
  config/                   # Env-based config struct
  middleware/               # Auth + logging middleware
  repository/               # Data access layer (PostgreSQL)
  service/                  # Business logic
  templates/                # Template cache + render
web/
  embed.go                  # //go:embed for static + templates
  static/css/output.css     # Built Tailwind CSS
  templates/layouts/        # Base layout templates
  templates/pages/          # Page templates
deployments/                # Docker + Compose files
configs/.env.example        # Env template
assets/input.css            # Source CSS
```

## Setup & Workflow
- **Prerequisites**: `.env` file is required (loaded via `godotenv`). Copy `configs/.env.example` as template.
- **Database**: `make dev_up` starts Postgres + pgAdmin via Compose.
- **Dependencies**: `go.mod` (Go), `bun.lock` (Node — Tailwind/Prettier).
- **Docker**: Production `Dockerfile` (multi-stage, alpine). Dev `Dockerfile.dev` (hot-reload with `air`).

## Developer Commands
- **Build**: `make build` (binary to `./bin/app`)
- **Dev (local)**: `make dev` (requires `air` globally installed — `go install github.com/air-verse/air@latest`)
- **Dev (container)**: `make dev_up`
- **Test**: `make test` (runs `go test -race` on all packages)
- **Lint**: `make lint` (requires `golangci-lint`)
- **CSS watch**: `make watch_css`
- **Formatting**: `make fmt` for Go. `bunx prettier --write "**/*.html"` for templates.
- **Env**: `COOKIESTORE_SECRET` is required (generate with `openssl rand -base64 32`).

## Agent Rules
- Run `make build` after Go changes to verify compilation.
- No new backend/frontend frameworks — stick to stdlib `http` + HTMX + Tailwind.
- Respect Prettier `go-template` plugin formatting when editing HTML templates.
- `.gitignore` excludes `**/output.css`, `.env`, `node_modules`, `bin/`, `tmp/`.
- `.dockerignore` excludes `.air.toml` and `data/`.

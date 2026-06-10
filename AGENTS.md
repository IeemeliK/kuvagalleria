# Agent Guidelines

## Architecture & Tech Stack
- **Backend**: Go 1.25+ stdlib `http.NewServeMux` routing. No Gin/Echo.
- **Database**: PostgreSQL via `database/sql` + `pgx/v5` driver.
- **Auth**: `gorilla/sessions` cookie store, bcrypt password verification.
- **Frontend**: HTMX, TailwindCSS v4 + DaisyUI v5.
- **Templates**: Go `html/template` with custom cache. Layouts in `internal/templates/layouts/`, pages in `internal/templates/pages/`. All embedded via `//go:embed`.
- **Static assets**: Built CSS at `internal/static/css/output.css` (embedded, served at `/static/`). Source CSS at `assets/input.css`.

## Setup & Workflow
- **Prerequisites**: `.env` file is required (loaded via `godotenv`). Use `.env.example` as template.
- **Database**: `docker compose -f docker-compose.yaml -f dev_compose.yaml up -d` starts Postgres + pgAdmin.
- **Dependencies**: `go.mod` (Go), `bun.lock` (Node — Tailwind/Prettier).
- **Docker**: Production `Dockerfile` (multi-stage, alpine). Dev `Dockerfile.dev` (hot-reload with `air`).

## Developer Commands
- **Build**: `make build` (binary to `./bin/app`)
- **Dev (local)**: `make dev` (requires `air` globally installed — `go install github.com/air-verse/air@latest`)
- **Dev (container)**: `make dev_up` (runs `docker compose -f docker-compose.yaml -f dev_compose.yaml up --build -d`)
- **CSS watch**: `make watch_css` (uses `bun` + `@tailwindcss/cli`)
- **Formatting**: `gofmt` for Go. `bunx prettier --write "**/*.html"` for templates (uses `prettier-plugin-go-template`).
- **Test**: No tests exist yet. `go test -v ./...` produces nothing.
- **Env**: `COOKIESTORE_SECRET` is required (generate with `openssl rand -base64 32`).

## Agent Rules
- Run `make build` after Go changes to verify compilation.
- No new backend/frontend frameworks — stick to stdlib `http` + HTMX + Tailwind.
- Respect Prettier `go-template` plugin formatting when editing HTML templates.
- `.gitignore` excludes `**/output.css`, `.env`, `node_modules`, `bin/`, `tmp/`.
- `.dockerignore` excludes `.air.toml` and `data/`.

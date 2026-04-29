# Agent Guidelines

## Architecture & Tech Stack
- **Backend**: Go 1.22+ standard library `http.NewServeMux` routing. No heavy frameworks (Gin/Echo).
- **Database**: PostgreSQL via `database/sql` and `pgx/v5` driver.
- **Frontend**: HTMX for interactions, TailwindCSS v4 + DaisyUI v5 for styling.
- **Structure**:
  - `internal/`: Application logic (`routes/`, `db/`, `middleware/`).
  - `internal/templates/`: Go HTML templates.
  - `assets/`: Raw static assets (e.g. `input.css`).
  - `internal/static/`: Built assets (e.g. `css/output.css`). Served at `/static/`.

## Setup & Workflow
- **Prerequisites**: A `.env` file is required to start (loaded via `godotenv`).
- **Database setup**: Run `docker compose -f dev_compose.yaml up -d` to start a local Postgres database.
- **Dependencies**: Managed via `go.mod` (Go) and `bun.lock` (Node/Tailwind/Prettier).

## Developer Commands
- **Build**: `make build` (outputs binary to `./bin/app`)
- **Dev Server**: `make dev` (requires `air` installed globally. Runs server with hot reload)
- **CSS**: `make watch_css` (uses `bun` & `@tailwindcss/cli` to watch changes)
- **Formatting**: Go files use standard `gofmt`. HTML templates are formatted using Prettier (`bunx prettier --write "**/*.html"`).
- **Test**: `go test -v ./...`
  - **Single Test**: `go test -v -run TestName ./internal/pkg/...`

## Agent Rules
- Always run `make build` to verify compilation after making Go changes.
- Do not introduce new backend or frontend frameworks. Stick to standard library `http` and HTMX/Tailwind.
- When modifying HTML templates, respect the Prettier `go-template` plugin formatting.

## Plan Mode
- Make the plan extremely concise. Sacrifice grammar for the sake of concision.
- At the end of each plan, give me a list of unresolved questions to answer, if any.

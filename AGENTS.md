# Agent Guidelines

## Build & Test Commands
- **Build**: `make build` (outputs binary to `./bin/app`)
- **Dev Server**: `make dev` (runs `air` with hot reload + CSS watch)
- **CSS**: `make watch_css` (uses `bun` & `@tailwindcss/cli`)
- **Test**: `go test -v ./...` (Run all tests)
  - **Single Test**: `go test -v -run TestName ./internal/pkg/...`

## Code Style & Conventions
- **Language**: Go (backend), standard `gofmt` formatting.
- **Frontend**: HTMX for interactivity, TailwindCSS for styling.
- **Structure**:
  - `internal/`: Private application code (routes, db, middleware).
  - `assets/`: Raw static assets (input.css).
  - `internal/templates/`: Go HTML templates.
- **Dependencies**: Managed via `go.mod` (backend) and `bun.lock` (frontend).

## Agent Rules
- Always run `make build` to verify compilation after changes.
- Do not introduce new heavy frameworks; stick to existing patterns.
- Check `Makefile` for available commands before running custom scripts.

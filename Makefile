composeFile := deployments/docker-compose.yaml
composeDevFile := deployments/dev_compose.yaml
composeFlags := -f $(composeFile) -f $(composeDevFile)
cssBuild := bunx @tailwindcss/cli -i ./public/input.css -o ./public/static/css/output.css --minify

include .env
export

.PHONY: build
build: init
	$(cssBuild)
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -v -o ./bin/app ./cmd/server

.PHONY: init
init:
	go mod tidy

.PHONY: dev_up
dev_up: dev_css
	docker compose --env-file .env $(composeFlags) up --build -d

.PHONY: dev_down
dev_down:
	docker compose $(composeFlags) down

.PHONY: dev_restart
dev_restart:
	docker compose $(composeFlags) restart

.PHONY: migrate-up
migrate-up:
	go tool goose up

.PHONY: migrate-down
migrate-down:
	go tool goose down

.PHONY: migrate-status
migrate-status:
	go tool goose status

.PHONY: migrate-create
migrate-create:
	go tool goose create $(name) sql

.PHONY: migrate-reset
migrate-reset:
	go tool goose reset

.PHONY: seed
seed:
	go run ./scripts/seed/main.go

.PHONY: dev_css
dev_css:
	bun i
	$(cssBuild)

.PHONY: watch_css
watch_css:
	$(cssBuild) --watch

.PHONY: test
test:
	go test -v -race -count=1 ./...

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: clean
clean:
	rm -rf tmp bin
	go clean

composeFile := deployments/docker-compose.yaml
composeDevFile := deployments/dev_compose.yaml
composeFlags := -f $(composeFile) -f $(composeDevFile)
cssBuild := bunx @tailwindcss/cli -i ./assets/input.css -o ./web/static/css/output.css
MIGRATIONS_DIR := ./migrations

include .env
export

.PHONY: build
build: init prod_css
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -v -o ./bin/app ./cmd/server

.PHONY: init
init:
	go mod tidy

.PHONY: dev_up
dev_up:
	docker compose $(composeFlags) up --build -d
	$(MAKE) migrate-up

.PHONY: dev_down
dev_down:
	docker compose $(composeFlags) down

.PHONY: migrate-up
migrate-up:
	goose -dir $(MIGRATIONS_DIR) postgres "$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(HOST):$(PORT)/$(POSTGRES_DB)" up

.PHONY: migrate-down
migrate-down:
	goose -dir $(MIGRATIONS_DIR) postgres "$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(HOST):$(PORT)/$(POSTGRES_DB)" down

.PHONY: migrate-status
migrate-status:
	goose -dir $(MIGRATIONS_DIR) postgres "$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(HOST):$(PORT)/$(POSTGRES_DB)" status

.PHONY: migrate-create
migrate-create:
	goose -dir $(MIGRATIONS_DIR) create $(name) sql

.PHONY: migrate-reset
migrate-reset:
	goose -dir $(MIGRATIONS_DIR) postgres "$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(HOST):$(PORT)/$(POSTGRES_DB)" reset

.PHONY: prod_css
prod_css:
	$(cssBuild) --minify

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

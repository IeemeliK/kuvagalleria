composeFile := docker-compose.yaml
composeDevFile := dev_compose.yaml
composeFlags := -f $(composeFile) -f $(composeDevFile)
cssBuild := bunx @tailwindcss/cli -i ./assets/input.css -o ./internal/static/css/output.css
.PHONY: build
build: init prod_css
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -v -o ./bin/app

.PHONY: init
init:
	go mod tidy

.PHONY: dev
dev: init dev_open
	$(GOPATH)/bin/air

.PHONY: dev_up
dev_up:
	docker compose $(composeFlags) up --build -d

.PHONY: dev_down
dev_down:
	docker compose $(composeFlags) down

.PHONY: dev_open
dev_open:
	open http://localhost:8080

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

.PHONY: clean
clean:
	rm -rf tmp bin
	go clean

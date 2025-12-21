cssBuild := bunx @tailwindcss/cli -i ./assets/input.css -o ./internal/static/css/output.css
.PHONY: build init dev dev_open prod_css dev_css watch_css clean
build: init prod_css
	go build -v -o ./bin/app

init:
	go mod tidy

dev: init dev_css dev_open
	$(GOPATH)/bin/air

dev_open:
	open http://localhost:8080

prod_css:
	$(cssBuild) --minify

dev_css:
	$(cssBuild)

watch_css:
	$(cssBuild) --watch

clean:
	rm -rf tmp bin
	go clean

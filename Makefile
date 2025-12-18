cssBuild := bunx @tailwindcss/cli -i ./assets/input.css -o ./static/css/output.css
.PHONY: build init prod_css dev dev_css watch_css clean dev_open
build: clean init prod_css
	go build -o ./bin/main ./cmd/web/

init:
	go mod tidy

dev: init dev_css dev_open
	air

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
	rm ./static/css/output.css

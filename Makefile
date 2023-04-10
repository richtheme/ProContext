.PHONY:
.SILENT:

build:
	go mod tidy
	go build -o .bin/bot main.go

run: build
	./.bin/bot


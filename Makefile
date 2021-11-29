build:
	go build -o bin/gobroke-server main.go

run: build
	bin/gobroke-server

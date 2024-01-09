test:
	go test ./... -coverprofile=coverage.out && go tool cover -html=coverage.out

start:
	go run cmd/main.go

build:
	 docker build -t matheussalencar/desafio-metaplane:latest -f Dockerfile.prod .

run:
	docker run --rm -p 8080:8080 matheussalencar/desafio-metaplane:latest

enter-container:
	docker exec -it metaplane bash

.PHONY: test start build run enter-container
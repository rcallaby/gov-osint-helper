BINARY=gov-osint-tool

all: build

build:
	go build -o bin/$(BINARY) ./cmd/gov-osint-tool

run:
	go run ./cmd/gov-osint-tool

docker-build:
	docker build -t gov-osint-tool .

docker-run:
	docker run -p 8080:8080 gov-osint-tool

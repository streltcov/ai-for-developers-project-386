.PHONY: run build clean setup test

run:
	cd api && go run ./cmd/server

build:
	cd api && go build -o ../bin/server ./cmd/server

clean:
	rm -rf bin/ api/booking.db

setup:
	cd api && go mod download

test:
	cd api && go test ./...

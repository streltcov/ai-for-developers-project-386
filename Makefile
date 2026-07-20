.PHONY: run build clean

run:
	cd api && go run ./cmd/server

build:
	cd api && go build -o ../bin/server ./cmd/server

clean:
	rm -rf bin/ api/booking.db

package main

import (
	"log"
	"net/http"
	"os"

	"call-booking/internal/handler"
	"call-booking/internal/repository"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "booking.db"
	}

	repo, err := repository.New(dbPath)
	if err != nil {
		log.Fatalf("failed to initialize repository: %v", err)
	}
	defer func() { _ = repo.Close() }()

	h := handler.New(repo)
	addr := ":" + port
	log.Printf("server starting on %s", addr)
	log.Fatal(http.ListenAndServe(addr, h.Routes()))
}

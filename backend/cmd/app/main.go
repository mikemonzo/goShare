package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080" // Default port
	}

	log.Printf("Starting server on port %s...", port)
	// TODO: Initialize routes, servicies, DB, middlewares, etc.

}

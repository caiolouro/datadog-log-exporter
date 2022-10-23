package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error while trying to load env vars: %v", err)
	}
	os.Exit(0)
	// ctx := datadog.NewDefaultContext(context.Background())
}

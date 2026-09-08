package main

import (
	"log"
	"storage-api/core/startup"
)

func main() {
	if err := startup.Startup(); err != nil {
		log.Fatalf("Application startup failed: %v", err)
	}
}

package main

import (
	"fmt"
	"log"

	"github.com/0xRichardL/otel-prom-practice/game/internal"
	"github.com/0xRichardL/otel-prom-practice/game/internal/services"
)

func main() {
	port := "8080"
	log.Printf("Starting game server on port %s", port)

	// Setup server with Gin router
	server := internal.NewServer(services.NewDice(), services.NewRoulette())
	router := server.SetupRouter()

	// Start server
	addr := fmt.Sprintf(":%s", port)
	log.Printf("Server listening on %s", addr)
	if err := router.Run(addr); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

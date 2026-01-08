package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/0xRichardL/otel-prom-practice/game/internal"
	"github.com/0xRichardL/otel-prom-practice/game/internal/metrics"
	"github.com/0xRichardL/otel-prom-practice/game/internal/services"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}
func run() error {
	// Handle SIGINT (CTRL+C) gracefully.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	otelConfig := &OtelConfig{
		CollectorEndpoint: "otel-collector:4318",
	}
	otelShutdown, err := otelConfig.Setup(ctx)
	if err != nil {
		return fmt.Errorf("failed to set up OpenTelemetry SDK: %v", err)
	}
	// Handle shutdown properly so nothing leaks.
	defer func() {
		err = errors.Join(err, otelShutdown(context.Background()))
	}()

	// Initialize metrics
	metrics, err := metrics.NewAppMetrics()
	if err != nil {
		return fmt.Errorf("failed to initialize metrics: %v", err)
	}

	// Setup server with Gin router
	app := internal.NewApp(services.NewDice(), services.NewRoulette(), metrics)
	router := app.SetupRouter()

	addr := ":8080"
	srv := &http.Server{
		Addr:    addr,
		Handler: metrics.WrapHandler(router),
	}

	srvErr := make(chan error, 1)
	// Start server
	go func() {
		log.Printf("Server listening on %s", addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			srvErr <- fmt.Errorf("Server failed to start: %v", err)
		}
	}()
	select {
	case err := <-srvErr:
		return err
	case <-ctx.Done():
		log.Println("Shutting down server...")
		stop()
	}

	// Shutdown server gracefully
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		return fmt.Errorf("server shutdown failed: %v", err)
	}
	log.Println("Server exited properly")
	return nil
}

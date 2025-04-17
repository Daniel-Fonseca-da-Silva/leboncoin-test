package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dafon/projects/leboncoin-test/internal/handler"
	"github.com/dafon/projects/leboncoin-test/internal/repository"
	"github.com/dafon/projects/leboncoin-test/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	statsRepo := repository.GetInstance()
	calculator := service.NewDefaultFizzBuzzCalculator()
	fizzBuzzService := service.NewFizzBuzzService(calculator, statsRepo)
	fizzBuzzHandler := handler.NewFizzBuzzHandler(fizzBuzzService)
	healthHandler := handler.NewHealthHandler()

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)
	r.Use(middleware.RequestID)
	r.Use(middleware.Timeout(60 * time.Second))

	fizzBuzzHandler.RegisterRoutes(r)

	r.Route("/api", func(r chi.Router) {
		healthHandler.RegisterRoutes(r)
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		log.Println("Server starting on port 8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Server shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited properly")
}

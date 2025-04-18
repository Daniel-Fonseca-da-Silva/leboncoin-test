package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dafon/projects/leboncoin-test/internal/config"
	"github.com/dafon/projects/leboncoin-test/internal/handler"
	"github.com/dafon/projects/leboncoin-test/internal/repository"
	"github.com/dafon/projects/leboncoin-test/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	logger := config.NewLogger("[FizzBuzz]")
	logger.SetLevel(config.INFO)

	statsRepo := repository.GetInstance()
	calculator := service.NewDefaultFizzBuzzCalculator()
	fizzBuzzService := service.NewFizzBuzzService(calculator, statsRepo)
	fizzBuzzHandler := handler.NewFizzBuzzHandler(fizzBuzzService)
	healthHandler := handler.NewHealthHandler()

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := middleware.GetReqID(r.Context())
			logger.Info("Request started", map[string]interface{}{
				"request_id": requestID,
				"method":     r.Method,
				"path":       r.URL.Path,
				"remote_ip":  r.RemoteAddr,
			})
			next.ServeHTTP(w, r)
		})
	})

	fizzBuzzHandler.RegisterRoutes(r)

	r.Route("/api", func(r chi.Router) {
		healthHandler.RegisterRoutes(r)
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		logger.Info("Server starting", map[string]interface{}{
			"port": 8080,
		})
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("Server failed to start", map[string]interface{}{
				"error": err.Error(),
			})
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("Server shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("Server forced to shutdown", map[string]interface{}{
			"error": err.Error(),
		})
	}

	logger.Info("Server exited properly")
}

package handler

import (
	"encoding/json"
	"net/http"

	"github.com/dafon/projects/leboncoin-test/internal/config"
	"github.com/dafon/projects/leboncoin-test/internal/model"
	"github.com/dafon/projects/leboncoin-test/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-playground/validator/v10"
)

// FizzBuzzHandler define a interface para o handler
type FizzBuzzHandler interface {
	RegisterRoutes(r chi.Router)
	CalculateFizzBuzz(w http.ResponseWriter, r *http.Request)
	GetStats(w http.ResponseWriter, r *http.Request)
}

// DefaultFizzBuzzHandler implementa FizzBuzzHandler
type DefaultFizzBuzzHandler struct {
	service  service.FizzBuzzService
	validate *validator.Validate
	logger   *config.Logger
}

func NewFizzBuzzHandler(service service.FizzBuzzService) FizzBuzzHandler {
	return &DefaultFizzBuzzHandler{
		service:  service,
		validate: validator.New(),
		logger:   config.NewLogger("[FizzBuzzHandler]"),
	}
}

// Aqui eu defino as rotas para o FizzBuzz
func (h *DefaultFizzBuzzHandler) RegisterRoutes(r chi.Router) {
	r.Post("/fizzbuzz", h.CalculateFizzBuzz)
	r.Get("/stats", h.GetStats)
}

// Aqui eu calculo o FizzBuzz junto com validacoes para erros caso o corpo da requisicao seja invalido
func (h *DefaultFizzBuzzHandler) CalculateFizzBuzz(w http.ResponseWriter, r *http.Request) {
	requestID := middleware.GetReqID(r.Context())

	var req model.FizzBuzzRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error("Invalid request body", map[string]interface{}{
			"request_id": requestID,
			"error":      err.Error(),
		})
		SendError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := h.validate.Struct(req); err != nil {
		h.logger.Error("Invalid request parameters", map[string]interface{}{
			"request_id": requestID,
			"error":      err.Error(),
			"request":    req,
		})
		SendError(w, http.StatusBadRequest, "Invalid request parameters")
		return
	}

	response := h.service.CalculateFizzBuzz(req)
	h.logger.Info("FizzBuzz calculation successful", map[string]interface{}{
		"request_id": requestID,
		"request":    req,
	})
	SendSuccess(w, "calculate_fizzbuzz", response)
}

// Aqui eu retorno as estatísticas da requisição mais frequente
func (h *DefaultFizzBuzzHandler) GetStats(w http.ResponseWriter, r *http.Request) {
	requestID := middleware.GetReqID(r.Context())
	stats := h.service.GetStats()

	h.logger.Info("Stats retrieved successfully", map[string]interface{}{
		"request_id": requestID,
		"stats":      stats,
	})
	SendSuccess(w, "get_stats", stats)
}

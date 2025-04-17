package handler

import (
	"encoding/json"
	"net/http"

	"github.com/dafon/projects/leboncoin-test/internal/model"
	"github.com/dafon/projects/leboncoin-test/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type FizzBuzzHandler struct {
	service  *service.FizzBuzzService
	validate *validator.Validate
}

func NewFizzBuzzHandler(service *service.FizzBuzzService) *FizzBuzzHandler {
	return &FizzBuzzHandler{
		service:  service,
		validate: validator.New(),
	}
}

// Aqui eu defino as rotas para o FizzBuzz
func (h *FizzBuzzHandler) RegisterRoutes(r chi.Router) {
	r.Post("/fizzbuzz", h.CalculateFizzBuzz)
	r.Get("/stats", h.GetStats)
}

// Aqui eu calculo o FizzBuzz junto com validacoes para erros caso o corpo da requisicao seja invalido
func (h *FizzBuzzHandler) CalculateFizzBuzz(w http.ResponseWriter, r *http.Request) {
	var req model.FizzBuzzRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.validate.Struct(req); err != nil {
		http.Error(w, "Invalid request parameters", http.StatusBadRequest)
		return
	}

	response := h.service.CalculateFizzBuzz(req)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Aqui eu retorno as estatísticas da requisição mais frequente
func (h *FizzBuzzHandler) GetStats(w http.ResponseWriter, r *http.Request) {
	stats := h.service.GetStats()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}

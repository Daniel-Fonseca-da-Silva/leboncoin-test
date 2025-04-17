package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// Aqui o RegisterRoutes registra as rotas de verificação da saúde do sistema
func (h *HealthHandler) RegisterRoutes(r chi.Router) {
	r.Get("/health", h.HealthCheck)
}

// Aqui eu retorno uma resposta simples OK para indicar que o serviço está em execução
func (h *HealthHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "ok"}`))
}

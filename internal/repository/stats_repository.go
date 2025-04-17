package repository

import (
	"fmt"
	"sync"

	"github.com/dafon/projects/leboncoin-test/internal/model"
	"github.com/dafon/projects/leboncoin-test/internal/service"
)

// DefaultStatsRepository implementa service.StatsRepository
type DefaultStatsRepository struct {
	mu       sync.RWMutex
	stats    map[string]int
	requests map[string]model.FizzBuzzRequest
}

var _ service.StatsRepository = (*DefaultStatsRepository)(nil)

var (
	instance *DefaultStatsRepository
	once     sync.Once
)

// GetInstance retorna o singleton da instancia do repositório
func GetInstance() service.StatsRepository {
	once.Do(func() {
		instance = &DefaultStatsRepository{
			stats:    make(map[string]int),
			requests: make(map[string]model.FizzBuzzRequest),
		}
	})
	return instance
}

// ResetInstance reseta a instancia do repositório
func ResetInstance() {
	instance = nil
	once = sync.Once{}
}

// IncrementStats incrementa o contador de requisições
func (r *DefaultStatsRepository) IncrementStats(req model.FizzBuzzRequest) {
	r.mu.Lock()
	defer r.mu.Unlock()

	key := r.generateKey(req)
	r.stats[key]++
	r.requests[key] = req
}

// GetMostFrequentRequest retorna a requisição mais frequente
func (r *DefaultStatsRepository) GetMostFrequentRequest() (model.FizzBuzzRequest, int) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var maxHits int
	var mostFrequentKey string

	// Aqui eu percorro o mapa de estatísticas e encontro a requisição mais frequente
	for key, hits := range r.stats {
		if hits > maxHits {
			maxHits = hits
			mostFrequentKey = key
		}
	}

	// Aqui eu verifico se a chave mais frequente é vazia
	if mostFrequentKey == "" {
		return model.FizzBuzzRequest{}, 0
	}

	return r.requests[mostFrequentKey], maxHits
}

// generateKey cria uma chave única para uma requisição
func (r *DefaultStatsRepository) generateKey(req model.FizzBuzzRequest) string {
	return fmt.Sprintf("%d-%d-%d-%s-%s", req.Int1, req.Int2, req.Limit, req.Str1, req.Str2)
}

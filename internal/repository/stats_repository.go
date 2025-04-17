package repository

import (
	"fmt"
	"sync"

	"github.com/dafon/projects/leboncoin-test/internal/model"
	"github.com/dafon/projects/leboncoin-test/internal/service"
)

type StatsRepository struct {
	mu       sync.RWMutex
	stats    map[string]int
	requests map[string]model.FizzBuzzRequest
}

var _ service.StatsRepository = (*StatsRepository)(nil)

var (
	instance *StatsRepository
	once     sync.Once
)

// Aqui eu retorno o singleton da instancia do repositório
func GetInstance() *StatsRepository {
	once.Do(func() {
		instance = &StatsRepository{
			stats:    make(map[string]int),
			requests: make(map[string]model.FizzBuzzRequest),
		}
	})
	return instance
}

// Aqui eu reseto a instancia do repositório para que seja criada uma nova instancia (principalmente para fins de teste)
func ResetInstance() {
	instance = nil
	once = sync.Once{}
}

// Aqui eu incremento o contador de requisições para uma requisição específica
func (r *StatsRepository) IncrementStats(req model.FizzBuzzRequest) {
	r.mu.Lock()
	defer r.mu.Unlock()

	key := r.generateKey(req)
	r.stats[key]++
	r.requests[key] = req
}

// Aqui eu retorno a requisição mais frequente e o número de hits
// Aqui eu uso um mutex para garantir que a leitura e escrita sejam seguras
// A função retorna a requisição mais frequente e o número de hits
// Aqui eu uso um mutex para garantir que a leitura e escrita sejam seguras
func (r *StatsRepository) GetMostFrequentRequest() (model.FizzBuzzRequest, int) {
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

// Aqui eu crio uma chave única para uma requisição do FizzBuzz
func (r *StatsRepository) generateKey(req model.FizzBuzzRequest) string {
	return fmt.Sprintf("%d-%d-%d-%s-%s", req.Int1, req.Int2, req.Limit, req.Str1, req.Str2)
}

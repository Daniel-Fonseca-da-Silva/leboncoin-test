package service

import (
	"strconv"

	"github.com/dafon/projects/leboncoin-test/internal/model"
)

type FizzBuzzCalculator interface {
	Calculate(req model.FizzBuzzRequest) []string
}

type StatsRepository interface {
	IncrementStats(req model.FizzBuzzRequest)
	GetMostFrequentRequest() (model.FizzBuzzRequest, int)
}

type FizzBuzzService struct {
	calculator FizzBuzzCalculator
	statsRepo  StatsRepository
}

type DefaultFizzBuzzCalculator struct{}

func NewDefaultFizzBuzzCalculator() *DefaultFizzBuzzCalculator {
	return &DefaultFizzBuzzCalculator{}
}

// Funcao que calcula o FizzBuzz recebendo um modelo FizzBuzzRequest com os parametros int1, int2, limit, str1 e str2
// E retorna um slice de strings com o resultado da calculacao
func (c *DefaultFizzBuzzCalculator) Calculate(req model.FizzBuzzRequest) []string {
	result := make([]string, req.Limit)

	for i := 1; i <= req.Limit; i++ {
		// Verifica se o número atual é múltiplo de int1 e/ou int2
		isInt1Multiple := i%req.Int1 == 0
		isInt2Multiple := i%req.Int2 == 0

		// Se int1 for 1, só consideramos múltiplo quando for exatamente igual a int1
		if req.Int1 == 1 {
			isInt1Multiple = i == req.Int1
		}

		// Aqui eu verifico se o número atual é múltiplo de int1 e/ou int2
		if isInt1Multiple && isInt2Multiple {
			result[i-1] = req.Str1 + req.Str2
		} else if isInt1Multiple {
			result[i-1] = req.Str1
		} else if isInt2Multiple {
			result[i-1] = req.Str2
		} else {
			result[i-1] = strconv.Itoa(i)
		}
	}

	return result
}

func NewFizzBuzzService(calculator FizzBuzzCalculator, statsRepo StatsRepository) *FizzBuzzService {
	return &FizzBuzzService{
		calculator: calculator,
		statsRepo:  statsRepo,
	}
}

// CalculateFizzBuzz performs the FizzBuzz calculation from 1 to limit
// Tenho uma funcao chamada CalculateFizzBuzz que recebe um modelo FizzBuzzRequest com os parametros int1, int2, limit, str1 e str2
// E retorna um modelo FizzBuzzResponse com o resultado da calculacao
func (s *FizzBuzzService) CalculateFizzBuzz(req model.FizzBuzzRequest) model.FizzBuzzResponse {
	result := s.calculator.Calculate(req)
	s.statsRepo.IncrementStats(req)

	return model.FizzBuzzResponse{
		Result: result,
	}
}

// Aqui eu retorno as estatísticas da requisição mais frequente
func (s *FizzBuzzService) GetStats() model.Stats {
	req, hits := s.statsRepo.GetMostFrequentRequest()
	return model.Stats{
		Request: req,
		Hits:    hits,
	}
}

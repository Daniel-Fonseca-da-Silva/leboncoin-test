package service

import (
	"strconv"

	"github.com/dafon/projects/leboncoin-test/internal/model"
	"github.com/dafon/projects/leboncoin-test/internal/repository"
)

type FizzBuzzService struct {
	statsRepo *repository.StatsRepository
}

func NewFizzBuzzService(statsRepo *repository.StatsRepository) *FizzBuzzService {
	return &FizzBuzzService{
		statsRepo: statsRepo,
	}
}

// CalculateFizzBuzz performs the FizzBuzz calculation from 1 to limit
// Tenho uma funcao chamada CalculateFizzBuzz que recebe um modelo FizzBuzzRequest com os parametros int1, int2, limit, str1 e str2
// E retorna um modelo FizzBuzzResponse com o resultado da calculacao
func (s *FizzBuzzService) CalculateFizzBuzz(req model.FizzBuzzRequest) model.FizzBuzzResponse {
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

	// Aqui eu registro a requisição para as estatísticas
	s.statsRepo.IncrementStats(req)

	// Aqui eu retorno o resultado da calculacao
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

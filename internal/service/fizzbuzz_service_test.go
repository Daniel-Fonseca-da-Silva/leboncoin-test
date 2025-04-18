package service

import (
	"testing"

	"github.com/dafon/projects/leboncoin-test/internal/model"
)

// Esta implementação é simples e não persiste os dados.
// Ela é usada apenas para testar o serviço FizzBuzz.
type SimpleStatsRepository struct {
	stats map[model.FizzBuzzRequest]int
}

func NewSimpleStatsRepository() *SimpleStatsRepository {
	return &SimpleStatsRepository{
		stats: make(map[model.FizzBuzzRequest]int),
	}
}

func (r *SimpleStatsRepository) IncrementStats(req model.FizzBuzzRequest) {
	r.stats[req]++
}

func (r *SimpleStatsRepository) GetMostFrequentRequest() (model.FizzBuzzRequest, int) {
	var mostFreq model.FizzBuzzRequest
	maxHits := 0

	for req, hits := range r.stats {
		if hits > maxHits {
			mostFreq = req
			maxHits = hits
		}
	}

	return mostFreq, maxHits
}

func TestDefaultFizzBuzzCalculator_Calculate(t *testing.T) {
	calculator := NewDefaultFizzBuzzCalculator()

	tests := []struct {
		name     string
		request  model.FizzBuzzRequest
		expected []string
	}{
		{
			name: "Standard FizzBuzz",
			request: model.FizzBuzzRequest{
				Int1:  3,
				Int2:  5,
				Limit: 15,
				Str1:  "Fizz",
				Str2:  "Buzz",
			},
			expected: []string{
				"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz",
				"11", "Fizz", "13", "14", "FizzBuzz",
			},
		},
		{
			name: "Custom Strings",
			request: model.FizzBuzzRequest{
				Int1:  2,
				Int2:  3,
				Limit: 6,
				Str1:  "Even",
				Str2:  "Three",
			},
			expected: []string{
				"1", "Even", "Three", "Even", "5", "EvenThree",
			},
		},
		{
			name: "Int1 equals 1",
			request: model.FizzBuzzRequest{
				Int1:  1,
				Int2:  2,
				Limit: 4,
				Str1:  "One",
				Str2:  "Two",
			},
			expected: []string{
				"One", "Two", "3", "Two",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calculator.Calculate(tt.request)
			if len(result) != len(tt.expected) {
				t.Errorf("Expected length %d, got %d", len(tt.expected), len(result))
				return
			}
			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("At index %d: expected %s, got %s", i, tt.expected[i], result[i])
				}
			}
		})
	}
}

func TestDefaultFizzBuzzService_CalculateFizzBuzz(t *testing.T) {
	calculator := NewDefaultFizzBuzzCalculator()
	statsRepo := NewSimpleStatsRepository()
	service := NewFizzBuzzService(calculator, statsRepo)

	request := model.FizzBuzzRequest{
		Int1:  3,
		Int2:  5,
		Limit: 5,
		Str1:  "Fizz",
		Str2:  "Buzz",
	}

	response := service.CalculateFizzBuzz(request)
	expected := []string{"1", "2", "Fizz", "4", "Buzz"}

	if len(response.Result) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(response.Result))
		return
	}

	for i := range response.Result {
		if response.Result[i] != expected[i] {
			t.Errorf("At index %d: expected %s, got %s", i, expected[i], response.Result[i])
		}
	}
}

func TestDefaultFizzBuzzService_GetStats(t *testing.T) {
	calculator := NewDefaultFizzBuzzCalculator()
	statsRepo := NewSimpleStatsRepository()
	service := NewFizzBuzzService(calculator, statsRepo)

	request := model.FizzBuzzRequest{
		Int1:  3,
		Int2:  5,
		Limit: 5,
		Str1:  "Fizz",
		Str2:  "Buzz",
	}

	service.CalculateFizzBuzz(request)
	service.CalculateFizzBuzz(request)

	stats := service.GetStats()
	if stats.Hits != 2 {
		t.Errorf("Expected 2 hits, got %d", stats.Hits)
	}

	if stats.Request != request {
		t.Error("Expected request to match the input request")
	}
}

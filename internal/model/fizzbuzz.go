package model

// Aqui eu defino o modelo de requisição para o FizzBuzz com validacoes e requisitos minimos
type FizzBuzzRequest struct {
	Int1  int    `json:"int1" validate:"required,min=1"`
	Int2  int    `json:"int2" validate:"required,min=1"`
	Limit int    `json:"limit" validate:"required,min=1"`
	Str1  string `json:"str1" validate:"required"`
	Str2  string `json:"str2" validate:"required"`
}

// Aqui eu defino o modelo de resposta para o FizzBuzz
type FizzBuzzResponse struct {
	Result []string `json:"result"`
}

// Aqui eu defino o modelo de resposta para as estatísticas
type Stats struct {
	Request FizzBuzzRequest `json:"request"`
	Hits    int             `json:"hits"`
}

// internal/infra/api/brasil_api_client_test.go

package api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jeancarlosdanese/go-temp-service/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestBrasilApiClient_FetchAddress(t *testing.T) {
	// Mock server para simular a resposta da API
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"cep": "01001-000",
			"state": "SP",
			"city": "São Paulo",
			"neighborhood": "Sé",
			"street": "Praça da Sé"
		}`))
	}))
	defer mockServer.Close()

	// Crie o cliente com a URL base do servidor de teste
	client := &BrasilApiClient{
		BaseURL: mockServer.URL,
	}

	ctx := context.Background()
	address, err := client.FetchAddress(ctx, "01001000")

	assert.NoError(t, err)
	assert.Equal(t, &entity.Address{
		Cep:        "01001-000",
		Logradouro: "Praça da Sé",
		Bairro:     "Sé",
		Cidade:     "São Paulo",
		Uf:         "SP",
	}, address)
}

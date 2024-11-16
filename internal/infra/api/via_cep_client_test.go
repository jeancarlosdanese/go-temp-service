package api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jeancarlosdanese/go-temp-service/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestViaCepClient_FetchAddress(t *testing.T) {
	// Mock server para simular a resposta da API
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"cep": "01001-000",
			"logradouro": "Praça da Sé",
			"bairro": "Sé",
			"localidade": "São Paulo",
			"uf": "SP"
		}`))
	}))
	defer mockServer.Close()

	// Modifique a URL usada no teste para apontar para o mock server
	client := &ViaCepClient{
		BaseURL: mockServer.URL, // Adicione uma propriedade BaseURL à struct ViaCepClient
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

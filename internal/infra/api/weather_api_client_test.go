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

func TestWeatherApiClient_FetchWeather(t *testing.T) {
	// Mock server para simular a resposta da API
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"current": {
				"temp_c": 25.0
			}
		}`))
	}))
	defer mockServer.Close()

	client := &WeatherApiClient{
		BaseURL: mockServer.URL,
	}

	ctx := context.Background()
	weather, err := client.FetchWeather(ctx, "São Paulo", "SP")

	assert.NoError(t, err)
	assert.Equal(t, &entity.WeatherData{
		TempC: 25.0,
		TempF: 77.0,
		TempK: 298.15,
	}, weather)
}

package web

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/jeancarlosdanese/go-temp-service/internal/usecase"
)

func WeatherHandler(ctx context.Context, addressUsecase *usecase.AddressUsecase, weatherUsecase *usecase.WeatherUsecase, w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Query().Get("cep")
	if len(cep) != 8 {
		http.Error(w, `{"message": "invalid zipcode"}`, http.StatusUnprocessableEntity)
		return
	}

	// Busca o endereço com base no CEP
	address, err := addressUsecase.GetAddress(ctx, cep)
	if err != nil {
		http.Error(w, `{"message": "can not find zipcode"}`, http.StatusNotFound)
		return
	}

	weather, err := weatherUsecase.GetWeather(ctx, address.Cidade, address.Uf)
	if err != nil {
		http.Error(w, `{"message": "error fetching weather data"}`, http.StatusInternalServerError)
		return
	}

	// Configura o cabeçalho e retorna a resposta
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(weather)
}

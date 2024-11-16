package entity_test

import (
	"testing"

	"github.com/jeancarlosdanese/go-temp-service/internal/entity"
)

func TestAddress(t *testing.T) {
	address := entity.Address{
		Cep:        "12345678",
		Logradouro: "Rua Teste",
		Bairro:     "Centro",
		Cidade:     "São Paulo",
		Uf:         "SP",
	}

	if address.Cep != "12345678" {
		t.Errorf("Expected CEP '12345678', got '%s'", address.Cep)
	}
}

func TestWeatherData(t *testing.T) {
	weatherData := entity.WeatherData{
		TempC: 25.0,
		TempF: 77.0,
		TempK: 298.0,
	}

	if weatherData.TempC != 25.0 {
		t.Errorf("Expected TempC 25.0, got '%f'", weatherData.TempC)
	}

	if weatherData.TempF != 77.0 {
		t.Errorf("Expected TempF 77.0, got '%f'", weatherData.TempF)
	}

	if weatherData.TempK != 298.0 {
		t.Errorf("Expected TempK 298.0, got '%f'", weatherData.TempK)
	}
}

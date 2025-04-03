// internal/infra/web/web_test.go

package web_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/jeancarlosdanese/go-temp-service/internal/entity"
	"github.com/jeancarlosdanese/go-temp-service/internal/infra/web"
	"github.com/jeancarlosdanese/go-temp-service/internal/usecase"
	"github.com/stretchr/testify/mock"
)

// MockAddressClient é um mock da interface AddressClientInterface
type MockAddressClient struct {
	mock.Mock
}

func (m *MockAddressClient) FetchAddress(ctx context.Context, cep string) (*entity.Address, error) {
	args := m.Called(ctx, cep)
	return args.Get(0).(*entity.Address), args.Error(1)
}

// MockWeatherClient é um mock da interface WeatherClientInterface
type MockWeatherClient struct {
	mock.Mock
}

func (m *MockWeatherClient) FetchWeather(ctx context.Context, city, state string) (*entity.WeatherData, error) {
	args := m.Called(ctx, city, state)
	return args.Get(0).(*entity.WeatherData), args.Error(1)
}

func TestWeatherHandler(t *testing.T) {
	// Criando mocks
	mockAddressClient := new(MockAddressClient)
	mockWeatherClient := new(MockWeatherClient)

	// Configurando os retornos dos mocks
	mockAddressClient.On("FetchAddress", mock.Anything, "12345678").Return(&entity.Address{
		Cep:    "12345678",
		Cidade: "São Paulo",
		Uf:     "SP",
	}, nil)

	mockWeatherClient.On("FetchWeather", mock.Anything, "São Paulo", "SP").Return(&entity.WeatherData{
		TempC: 25.0,
		TempF: 77.0,
		TempK: 298.15,
	}, nil)

	// Criando as usecases com os mocks
	addressUsecase := usecase.NewAddressUsecase(mockAddressClient)
	weatherUsecase := usecase.NewWeatherUsecase(mockWeatherClient)

	// Criando a requisição de teste
	req, err := http.NewRequest("GET", "/weather?cep=12345678", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		web.WeatherHandler(context.Background(), addressUsecase, weatherUsecase, w, r)
	})

	// Executa o manipulador
	handler.ServeHTTP(rr, req)

	// Verifica o status de resposta
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Verifica o corpo da resposta
	expected := entity.WeatherData{
		TempC: 25.0,
		TempF: 77.0,
		TempK: 298.15,
	}

	var responseBody entity.WeatherData
	if err := json.NewDecoder(rr.Body).Decode(&responseBody); err != nil {
		t.Fatalf("could not decode response: %v", err)
	}

	if !reflect.DeepEqual(responseBody, expected) {
		t.Errorf("handler returned unexpected body: got %+v want %+v", responseBody, expected)
	}
}

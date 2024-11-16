package main

import (
	"log"
	"os"

	"github.com/jeancarlosdanese/go-temp-service/internal/infra/api"
	"github.com/jeancarlosdanese/go-temp-service/internal/infra/web"
	"github.com/jeancarlosdanese/go-temp-service/internal/usecase"
	"github.com/joho/godotenv"
)

func main() {
	// Carregar variáveis de ambiente do arquivo .env
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found or unable to load .env")
	}

	// Recuperar a chave de API do ambiente
	apiKey := os.Getenv("WEATHER_API_KEY")
	if apiKey == "" {
		log.Fatal("WEATHER_API_KEY not set in environment")
	}

	// Inicializando os clientes de API

	viaCepClient := &api.ViaCepClient{}
	brasilApiClient := &api.BrasilApiClient{}
	weatherClient := api.NewWeatherApiClient(apiKey)

	// Inicializando as camadas de usecase com os clientes de API
	addressUsecase := usecase.NewAddressUsecase(viaCepClient, brasilApiClient)
	weatherUsecase := usecase.NewWeatherUsecase(weatherClient)

	// Configura o roteador
	mux := web.NewRouter(addressUsecase, weatherUsecase)

	// Inicia o servidor
	web.StartServer(mux)
}

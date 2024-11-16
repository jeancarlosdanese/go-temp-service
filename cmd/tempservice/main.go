package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

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

	// Configurando o manipulador HTTP com a usecase
	http.HandleFunc("/weather", func(w http.ResponseWriter, r *http.Request) {
		// Criando um contexto com timeout para a requisição
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		// Chama o manipulador passando o contexto e as usecases
		web.WeatherHandler(ctx, addressUsecase, weatherUsecase, w, r)
	})

	// Serve a pasta `public` na rota raiz
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	// Iniciando o servidor
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %v\n", err)
	}
}

// internal/infra/web/router.go

package web

import (
	"context"
	"net/http"
	"time"

	"github.com/jeancarlosdanese/go-temp-service/internal/usecase"
)

// NewRouter configura e retorna as rotas do servidor
func NewRouter(addressUsecase *usecase.AddressUsecase, weatherUsecase *usecase.WeatherUsecase) http.Handler {
	mux := http.NewServeMux()

	// Configura a rota /weather
	mux.HandleFunc("/weather", func(w http.ResponseWriter, r *http.Request) {
		// Criando um contexto com timeout para a requisição
		ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
		defer cancel()

		// Chama o manipulador de clima passando o contexto e as usecases
		WeatherHandler(ctx, addressUsecase, weatherUsecase, w, r)
	})

	// Serve arquivos estáticos na rota raiz
	fileServer := http.FileServer(http.Dir("./public"))
	mux.Handle("/", fileServer)

	return mux
}

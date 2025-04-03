// internal/interfaces/weather_client.go

package interfaces

import (
	"context"

	"github.com/jeancarlosdanese/go-temp-service/internal/entity"
)

type WeatherClientInterface interface {
	FetchWeather(ctx context.Context, city, state string) (*entity.WeatherData, error)
}

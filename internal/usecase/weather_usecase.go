// internal/usecase/weather_usecase.go

package usecase

import (
	"context"

	"github.com/jeancarlosdanese/go-temp-service/internal/entity"
	"github.com/jeancarlosdanese/go-temp-service/internal/interfaces"
)

type WeatherUsecase struct {
	weatherClient interfaces.WeatherClientInterface
}

func NewWeatherUsecase(client interfaces.WeatherClientInterface) *WeatherUsecase {
	return &WeatherUsecase{weatherClient: client}
}

func (u *WeatherUsecase) GetWeather(ctx context.Context, city, state string) (*entity.WeatherData, error) {
	return u.weatherClient.FetchWeather(ctx, city, state)
}

// internal/infra/api/weather_api_client.go

package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"unicode"

	"github.com/jeancarlosdanese/go-temp-service/internal/entity"
	"github.com/jeancarlosdanese/go-temp-service/internal/interfaces"
)

type WeatherApiClient struct {
	apiKey  string
	BaseURL string
}

func NewWeatherApiClient(apiKey string) *WeatherApiClient {
	return &WeatherApiClient{apiKey: apiKey}
}

func (w *WeatherApiClient) FetchWeather(ctx context.Context, city, state string) (*entity.WeatherData, error) {
	location := url.QueryEscape(fmt.Sprintf("%s,%s,%s", w.removeAccents(city), state, "Brazil"))
	url := fmt.Sprintf("%s/v1/current.json?key=%s&q=%s", w.getBaseURL(), w.apiKey, location)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error fetching weather: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Error from Weather API: status code %d, URL: %s", resp.StatusCode, url)
		return nil, fmt.Errorf("error from Weather API: status code %d", resp.StatusCode)
	}

	var weatherResponse struct {
		Current struct {
			TempC float64 `json:"temp_c"`
		} `json:"current"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&weatherResponse); err != nil {
		return nil, fmt.Errorf("error decoding Weather API response: %w", err)
	}

	return &entity.WeatherData{
		TempC: weatherResponse.Current.TempC,
		TempF: weatherResponse.Current.TempC*1.8 + 32,
		TempK: weatherResponse.Current.TempC + 273.15,
	}, nil
}

func (w *WeatherApiClient) getBaseURL() string {
	if w.BaseURL != "" {
		return w.BaseURL
	}
	return "http://api.weatherapi.com"
}

// removeAccents - Remove acentos de uma string
func (w *WeatherApiClient) removeAccents(input string) string {
	var output strings.Builder
	for _, r := range input {
		if r > 127 {
			r = unicode.SimpleFold(r)
			switch r {
			case 'á', 'à', 'â', 'ã', 'ä', 'Á', 'À', 'Â', 'Ã', 'Ä':
				r = 'a'
			case 'é', 'è', 'ê', 'ë', 'É', 'È', 'Ê', 'Ë':
				r = 'e'
			case 'í', 'ì', 'î', 'ï', 'Í', 'Ì', 'Î', 'Ï':
				r = 'i'
			case 'ó', 'ò', 'ô', 'õ', 'ö', 'Ó', 'Ò', 'Ô', 'Õ', 'Ö':
				r = 'o'
			case 'ú', 'ù', 'û', 'ü', 'Ú', 'Ù', 'Û', 'Ü':
				r = 'u'
			case 'ç', 'Ç':
				r = 'c'
			default:
				continue
			}
		}
		output.WriteRune(r)
	}
	return output.String()
}

// Certificando que WeatherApiClient implementa WeatherClientInterface
var _ interfaces.WeatherClientInterface = (*WeatherApiClient)(nil)

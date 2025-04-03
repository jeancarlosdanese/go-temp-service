// internal/entity/weather_data.go

package entity

import (
	"encoding/json"
	"math"
)

type WeatherData struct {
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

// roundToTwoDecimalPlaces arredonda um float para 2 casas decimais
func roundToTwoDecimalPlaces(value float64) float64 {
	return math.Round(value*100) / 100
}

// MarshalJSON garante que os valores numéricos sejam formatados com 2 casas decimais
func (w WeatherData) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		TempC float64 `json:"temp_C"`
		TempF float64 `json:"temp_F"`
		TempK float64 `json:"temp_K"`
	}{
		TempC: roundToTwoDecimalPlaces(w.TempC),
		TempF: roundToTwoDecimalPlaces(w.TempF),
		TempK: roundToTwoDecimalPlaces(w.TempK),
	})
}

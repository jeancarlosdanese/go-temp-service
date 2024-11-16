package entity

// WeatherData - struct para armazenar os dados de temperatura em diferentes escalas
type WeatherData struct {
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

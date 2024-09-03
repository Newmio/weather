package dto

import "weather/internal/domain/entity"

type weatherResponse struct {
	City      string  `json:"city"`
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	WindSpeed float64 `json:"wind_speed"`
	WindDeg   int     `json:"wind_deg"`
}

func NewWeatherListResponse(weather []entity.List) []weatherResponse {
	var response []weatherResponse

	for _, w := range weather {
		response = append(response, NewWeatherResponse([]entity.List{w}))
	}

	return response
}

func NewWeatherResponse(weather []entity.List) weatherResponse {
	return weatherResponse{
		City:      weather[0].Name,
		Temp:      weather[0].Main.Temp,
		FeelsLike: weather[0].Main.FeelsLike,
		TempMin:   weather[0].Main.TempMin,
		TempMax:   weather[0].Main.TempMax,
		WindSpeed: weather[0].Wind.Speed,
		WindDeg:   weather[0].Wind.Deg,
	}
}

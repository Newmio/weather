package service

import (
	"weather/internal/domain/entity"
)

func (s *service) GetWeather(cityId int) (entity.Weather, error) {
	return s.r.GetWeather(cityId)
}

func (s *service) GetWeatherList(citiesId []int) (entity.Weather, error) {
	return s.r.GetWeatherList(citiesId)
}

func (s *service) GetAverage(cities []string) ([]entity.WeatherAverage, error) {
	var average []entity.WeatherAverage
	weathers := make(map[string][]entity.List)

	for _, city := range cities {

		weather, err := s.r.GetForecast(city)
		if err != nil {
			return nil, err
		}

		weathers[city] = weather.List
	}

	for city, list := range weathers {
		var temp float64
		var humidity int

		for _, item := range list {
			temp += (item.Main.TempMin + item.Main.TempMax) / 2
			humidity += item.Main.Humidity
		}

		average = append(average, entity.WeatherAverage{
			City:     city,
			Temp:     int(temp / float64(len(list))),
			Humidity: humidity / len(list),
		})
	}
	return average, nil
}

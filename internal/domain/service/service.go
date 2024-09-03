package service

import (
	"weather/internal/domain/entity"
	"weather/internal/repository"
)

//go:generate mockery --name=IService --output=./mocks --case=underscore
type IService interface {
	GetWeather(cityId int) (entity.Weather, error)
	GetWeatherList(citiesId []int) (entity.Weather, error)
	GetAverage(cities []string) ([]entity.WeatherAverage, error)
}

type service struct {
	r repository.IRepository
}

func NewService(r repository.IRepository) IService {
	return &service{r: r}
}

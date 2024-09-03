package repository

import (
	"weather/internal/domain/entity"
	repocache "weather/internal/repository/cache"
	repohttp "weather/internal/repository/http"
)

//go:generate mockery --name=IRepository --output=./mocks --case=underscore
type IRepository interface {
	GetWeather(cityId int) (entity.Weather, error)
	GetWeatherList(citiesId []int) (entity.Weather, error)
	GetForecast(city string) (entity.Weather, error)
}

type repo struct {
	http  repohttp.IHttp
	cache repocache.ICache
}

func NewRepo(http repohttp.IHttp, cache repocache.ICache) IRepository {
	return &repo{
		http:  http,
		cache: cache,
	}
}

func (r *repo) GetWeather(cityId int) (entity.Weather, error) {
	weather := r.cache.GetWeather(cityId)
	if len(weather.List) == 0 {
		webWeather, err := r.http.GetWeather(cityId)
		if err != nil {
			return weather, err
		}

		r.cache.SetWeather(cityId, webWeather)
		return webWeather, nil
	}

	return weather, nil
}

func (r *repo) GetWeatherList(citiesId []int) (entity.Weather, error) {
	var weather entity.Weather
	var newIds []int

	for _, id := range citiesId {
		w := r.cache.GetWeather(id)

		if len(w.List) == 0 {
			newIds = append(newIds, id)
		}

		weather.List = append(weather.List, w.List...)
	}

	if len(newIds) == 0 {
		return weather, nil
	}

	webWeather, err := r.http.GetWeatherList(newIds)
	if err != nil {
		return weather, err
	}

	for _, w := range webWeather.List {
		r.cache.SetWeather(w.Id, entity.Weather{List: []entity.List{w}})
		weather.List = append(weather.List, w)
	}

	return weather, nil
}

func (r *repo) GetForecast(city string) (entity.Weather, error) {
	weather := r.cache.GetForecast(city)
	if len(weather.List) == 0 {

		webWeather, err := r.http.GetForecast(city)
		if err != nil {
			return weather, err
		}

		r.cache.SetForecast(city, webWeather)
		return webWeather, nil
	}
	
	return weather, nil
}

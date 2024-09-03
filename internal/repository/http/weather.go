package repohttp

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"weather/internal/domain/entity"
)

func (r *httprepo) GetForecast(city string) (entity.Weather, error) {
	var weather entity.Weather

	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/forecast?units=metric&q=%s&appid=%s", city, r.token)

	req := entity.Request{
		Url:    url,
		Method: "GET",
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: nil,
	}

	resp, err := r.do(req)
	if err != nil {
		return entity.Weather{}, err
	}

	if resp.Status != 200 {
		return entity.Weather{}, r.getError(resp.Body)
	}

	err = json.Unmarshal(resp.Body, &weather)
	if err != nil {
		return entity.Weather{}, err
	}

	return weather, nil
}

func (r *httprepo) GetWeatherList(citiesId []int) (entity.Weather, error) {
	var weather entity.Weather
	var cities []string

	for _, id := range citiesId {
		cities = append(cities, strconv.Itoa(id))
	}

	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/group?id=%s&units=metric&appid=%s", strings.Join(cities, ","), r.token)

	req := entity.Request{
		Url:    url,
		Method: "GET",
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: nil,
	}

	resp, err := r.do(req)
	if err != nil {
		return entity.Weather{}, err
	}

	if resp.Status != 200 {
		return entity.Weather{}, r.getError(resp.Body)
	}

	err = json.Unmarshal(resp.Body, &weather)
	if err != nil {
		return entity.Weather{}, err
	}

	return weather, nil
}

func (r *httprepo) GetWeather(cityId int) (entity.Weather, error) {
	var weather entity.Weather

	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/group?id=%d&units=metric&appid=%s", cityId, r.token)

	req := entity.Request{
		Url:    url,
		Method: "GET",
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: nil,
	}

	resp, err := r.do(req)
	if err != nil {
		return entity.Weather{}, err
	}

	if resp.Status != 200 {
		return entity.Weather{}, r.getError(resp.Body)
	}

	err = json.Unmarshal(resp.Body, &weather)
	if err != nil {
		return entity.Weather{}, err
	}

	return weather, nil
}

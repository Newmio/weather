package transport

import (
	"fmt"
	"weather/internal/domain/dto"

	"github.com/labstack/echo/v4"
)

func (h *handler) average(c echo.Context) error {
	var cities []string

	if len(h.cities) == 0 {
		return c.JSON(400, h.errorResp(fmt.Errorf("empty cities map")))
	}

	for city := range h.cities {
		cities = append(cities, city)
	}

	average, err := h.s.GetAverage(cities)
	if err != nil {
		return c.JSON(500, h.errorResp(err))
	}

	return c.JSON(200, average)
}

func (h *handler) weatherByCity(c echo.Context) error {
	city := c.Param("city")

	if _, ok := h.cities[city]; !ok {
		return c.JSON(400, h.errorResp(fmt.Errorf("city not found")))
	}

	weather, err := h.s.GetWeather(h.cities[city])
	if err != nil {
		return c.JSON(500, h.errorResp(err))
	}

	if len(weather.List) == 0{
		return c.JSON(400, h.errorResp(fmt.Errorf("city not found")))
	}

	return c.JSON(200, dto.NewWeatherResponse(weather.List))
}

func (h *handler) weather(c echo.Context) error {
	var citiesId []int

	if len(h.cities) == 0 {
		return c.JSON(400, h.errorResp(fmt.Errorf("empty cities map")))
	}

	for _, id := range h.cities {
		citiesId = append(citiesId, id)
	}

	weather, err := h.s.GetWeatherList(citiesId)
	if err != nil {
		return c.JSON(500, h.errorResp(err))
	}

	if len(weather.List) != len(h.cities){
		return c.JSON(400, h.errorResp(fmt.Errorf("city not found")))
	}

	return c.JSON(200, dto.NewWeatherListResponse(weather.List))
}
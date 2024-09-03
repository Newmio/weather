package repocache

import (
	"fmt"
	"sync"
	"time"
	"weather/internal/domain/entity"
)

//go:generate mockery --name=ICache --output=./mocks --case=underscore
type ICache interface {
	GetWeather(cityId int) entity.Weather
	SetWeather(cityId int, weather entity.Weather)
	GetForecast(city string) entity.Weather
	SetForecast(city string, average entity.Weather)
}

type cache struct {
	weather  map[int]entity.Weather
	forecast map[string]entity.Weather
	mu       sync.RWMutex
}

func NewChache() ICache {
	c := &cache{
		weather:  make(map[int]entity.Weather),
		forecast: make(map[string]entity.Weather),
	}

	go c.clearWeather()
	go c.clearAverage()

	return c
}

func (c *cache) clearWeather() {
	ticker := time.NewTicker(time.Minute * 5)

	for range ticker.C {
		c.mu.Lock()
		c.weather = make(map[int]entity.Weather)
		fmt.Println("deleted weather")
		c.mu.Unlock()
	}
}

func (c *cache) clearAverage() {
	ticker := time.NewTicker(time.Hour * 24)

	for range ticker.C {
		c.mu.Lock()
		c.forecast = make(map[string]entity.Weather)
		c.mu.Unlock()
	}
}

func (c *cache) GetWeather(cityId int) entity.Weather {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if weather, ok := c.weather[cityId]; ok {
		return weather
	}

	return entity.Weather{}
}

func (c *cache) SetWeather(cityId int, weather entity.Weather) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.weather[cityId] = weather
}

func (c *cache) GetForecast(city string) entity.Weather {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if forecast, ok := c.forecast[city]; ok {
		return forecast
	}

	return entity.Weather{}
}

func (c *cache) SetForecast(city string, average entity.Weather) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.forecast[city] = average
}

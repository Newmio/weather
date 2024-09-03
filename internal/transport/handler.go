package transport

import (
	"weather/internal/domain/service"

	"github.com/labstack/echo/v4"
)

type handler struct {
	s      service.IService
	cities map[string]int
}

func NewHandler(cities map[string]int, s service.IService) *handler {
	return &handler{
		s:      s,
		cities: cities,
	}
}

func (h *handler) InitRoutes(e *echo.Echo) {
	e.GET("/weather/:city", h.weatherByCity)
	e.GET("/weather", h.weather)
	e.GET("/weather/average", h.average)
}

func (h *handler) errorResp(err error) map[string]interface{} {
	return map[string]interface{}{"error": err.Error()}
}

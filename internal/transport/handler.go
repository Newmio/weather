package transport

import (
	"path/filepath"
	"strings"
	"text/template"
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
	e.GET("/", func(c echo.Context) error { return c.File("template/index.html") })
	e.GET("/weather/:city", h.weatherByCity)
	e.GET("/weather", h.weather)
	e.GET("/weather/average", h.average)
}

func (h *handler) errorResp(err error) map[string]interface{} {
	return map[string]interface{}{"error": err.Error()}
}

func renderHtml(directory string, data interface{}) (string, error) {
	buffer := new(strings.Builder)

	name := filepath.Base(directory)

	tmpl, err := template.New(name).ParseFiles(directory)
	if err != nil {
		return "", err
	}

	if err := tmpl.ExecuteTemplate(buffer, name, data); err != nil {
		return "", err
	}

	return buffer.String(), nil
}

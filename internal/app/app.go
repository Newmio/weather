package app

import (
	"fmt"
	"net/http"
	"time"
	"weather/internal/domain/service"
	"weather/internal/repository"
	repocache "weather/internal/repository/cache"
	repohttp "weather/internal/repository/http"
	"weather/internal/transport"

	"github.com/fatih/color"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Run() error {

	cfg, err := getCities()
	if err != nil {
		return err
	}

	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:        10,
			IdleConnTimeout:     30 * time.Second,
			TLSHandshakeTimeout: 10 * time.Second,
		},
		Timeout: 10 * time.Second,
	}

	httpRepo := repohttp.NewHttpRepo(client, cfg.Token)
	cacheRepo := repocache.NewChache()
	repo := repository.NewRepo(httpRepo, cacheRepo)
	service := service.NewService(repo)
	handler := transport.NewHandler(cfg.Cities, service)

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
	}))

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("Cache-Control", "no-cache, no-store")
			return next(c)
		}
	})

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           " [ ${time_custom} ]  ${latency_human}  ${status}   ${method}   ${uri}\n\n",
		CustomTimeFormat: "2006/01/02 15:04:05",
		Output:           color.Output,
	}))

	e.Static("/template", "template")

	handler.InitRoutes(e)

	printRoutes(e)
	e.Logger.Fatal(e.Start(":8088"))

	return nil
}

func printRoutes(e *echo.Echo) {
	color.New(color.BgHiBlack, color.Bold).Println("                                                    ")

	for _, value := range e.Routes() {
		var customColor color.Attribute

		switch value.Method {
		case "GET":
			customColor = color.FgGreen

		case "POST":
			customColor = color.FgHiYellow

		case "PUT":
			customColor = color.FgBlue

		case "PATCH":
			customColor = color.FgMagenta

		case "DELETE":
			customColor = color.FgHiRed

		default:
			continue
		}

		color.New(customColor, color.Bold).Print(fmt.Sprintf("\n\t%s : %s", value.Path, value.Method))
	}

	fmt.Println()
	color.New(color.BgHiBlack, color.Bold).Println("                                                    ")
}

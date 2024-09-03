package app

import (
	"net/http"
	"time"
	"weather/internal/domain/service"
	"weather/internal/repository"
	repocache "weather/internal/repository/cache"
	repohttp "weather/internal/repository/http"
	"weather/internal/transport"

	"github.com/labstack/echo/v4"
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
	handler.InitRoutes(e)
	e.Logger.Fatal(e.Start(":8088"))

	return nil
}

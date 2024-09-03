package repohttp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"weather/internal/domain/entity"
)

//go:generate mockery --name=IHttp --output=./mocks --case=underscore

type IHttp interface{
	GetWeather(cityId int) (entity.Weather, error)
	GetWeatherList(citiesId []int) (entity.Weather, error)
	GetForecast(city string) (entity.Weather, error)
}

type httprepo struct {
	client *http.Client
	token  string
}

func NewHttpRepo(client *http.Client, token string) IHttp {
	return &httprepo{
		client: client,
		token:  token,
	}
}

func (r *httprepo) getError(body []byte)error{
	var data map[string]interface{}

	err := json.Unmarshal(body, &data)
	if err != nil {
		return err
	}

	if _, ok := data["message"]; ok {
		return fmt.Errorf(data["message"].(string))
	}

	return fmt.Errorf("unknown error")
}

func (r *httprepo) do(req entity.Request) (entity.Response, error) {
	var body []byte

	if req.Body != nil {
		b, err := json.Marshal(req)
		if err != nil {
			return entity.Response{}, err
		}
		body = b
	}

	request, err := http.NewRequest(req.Method, req.Url, bytes.NewBuffer(body))
	if err != nil {
		return entity.Response{}, err
	}

	for key, value := range req.Headers {
		request.Header.Add(key, value)
	}

	response, err := r.client.Do(request)
	if err != nil {
		return entity.Response{}, err
	}
	defer response.Body.Close()

	body, err = io.ReadAll(response.Body)
	if err != nil {
		return entity.Response{}, err
	}

	return entity.Response{
		Body:    body,
		Headers: response.Header,
		Status:  response.StatusCode,
	}, nil
}

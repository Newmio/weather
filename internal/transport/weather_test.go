package transport

import (
	"fmt"
	"net/http/httptest"
	"testing"
	"weather/internal/domain/entity"
	"weather/internal/domain/service/mocks"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandler_average(t *testing.T) {
	type mockBehavior func(s *mocks.IService)

	testTable := []struct {
		name         string
		cities       map[string]int
		mockBehavior mockBehavior
		expectedCode int
		expectedBody string
	}{
		{
			name: "OK",
			cities: map[string]int{
				"kyiv": 703448, "vilnius": 593116,
				"riga": 456173, "tallinn": 588409,
				"sofia": 727011, "minsk": 625144,
			},
			mockBehavior: func(s *mocks.IService) {
				weather := []entity.WeatherAverage{
					{
						City: "riga",
						Temp: 20,
						Humidity: 55,
					},
					{
						City: "tallinn",
						Temp: 18,
						Humidity: 66,
					},
					{
						City: "sofia",
						Temp: 22,
						Humidity: 51,
					},
					{
						City: "minsk",
						Temp: 19,
						Humidity: 41,
					},
					{
						City: "kyiv",
						Temp: 23,
						Humidity: 24,
					},
					{
						City: "vilnius",
						Temp: 19,
						Humidity: 48,
					},
				}

				s.On("GetAverage",  mock.MatchedBy(func(arg []string) bool { return true })).Return(weather, nil)
			},
			expectedCode: 200,
			expectedBody: `[{"city":"riga","temp":20,"humidity":55},{"city":"tallinn","temp":18,"humidity":66},{"city":"sofia","temp":22,"humidity":51},{"city":"minsk","temp":19,"humidity":41},{"city":"kyiv","temp":23,"humidity":24},{"city":"vilnius","temp":19,"humidity":48}]`,
		},
		{
			name:         "bad request empty cities map",
			cities:       map[string]int{},
			mockBehavior: func(s *mocks.IService) {},
			expectedCode: 400,
			expectedBody: `{"error":"empty cities map"}`,
		},
		{
			name:         "bad request empty cities map",
			cities: map[string]int{
				"kyiv": 703448, "vilnius": 593116,
				"riga": 456173, "tallinn": 588409,
				"sofia": 727011, "minsk": 625144,
			},
			mockBehavior: func(s *mocks.IService) {
				s.On("GetAverage",  mock.MatchedBy(func(arg []string) bool { return true })).Return([]entity.WeatherAverage{}, 
					fmt.Errorf("Invalid API key. Please see https://openweathermap.org/faq#error401 for more info."))
			},
			expectedCode: 500,
			expectedBody: `{"error":"Invalid API key. Please see https://openweathermap.org/faq#error401 for more info."}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {

			mockService := new(mocks.IService)
			testCase.mockBehavior(mockService)

			h := handler{
				s:      mockService,
				cities: testCase.cities,
			}

			e := echo.New()
			h.InitRoutes(e)

			req := httptest.NewRequest("GET", "/weather/average", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			err := h.average(c)
			assert.NoError(t, err)
			assert.Equal(t, testCase.expectedCode, rec.Code)
			assert.JSONEq(t, testCase.expectedBody, rec.Body.String())

			mockService.AssertExpectations(t)
		})
	}
}

func TestHandler_weather(t *testing.T) {
	type mockBehavior func(s *mocks.IService)

	testTable := []struct {
		name         string
		cities       map[string]int
		mockBehavior mockBehavior
		expectedCode int
		expectedBody string
	}{
		{
			name: "OK",
			cities: map[string]int{
				"kyiv": 703448, "vilnius": 593116,
				"riga": 456173, "tallinn": 588409,
				"sofia": 727011, "minsk": 625144,
			},
			mockBehavior: func(s *mocks.IService) {
				weather := entity.Weather{
					List: []entity.List{
						{
							Main: entity.Main{
								Temp:      20.85,
								FeelsLike: 20.77,
								TempMin:   20.19,
								TempMax:   22.19,
								Pressure:  1025,
								Humidity:  68,
							},
							Visibility: 10000,
							Wind: entity.Wind{
								Speed: 3.09,
								Deg:   330,
							},
							Name: "Tallinn",
						},
						{
							Main: entity.Main{
								Temp:      29.83,
								FeelsLike: 28.3,
								TempMin:   28.91,
								TempMax:   29.83,
								Pressure:  1015,
								Humidity:  25,
							},
							Visibility: 10000,
							Wind: entity.Wind{
								Speed: 3.09,
								Deg:   60,
							},
							Name: "Sofia",
						},
						{
							Main: entity.Main{
								Temp:      26.86,
								FeelsLike: 26.58,
								TempMin:   26.86,
								TempMax:   26.86,
								Pressure:  1023,
								Humidity:  36,
							},
							Visibility: 10000,
							Wind: entity.Wind{
								Speed: 4.25,
								Deg:   70,
							},
							Name: "Minsk",
						},
						{
							Main: entity.Main{
								Temp:      32.07,
								FeelsLike: 31.41,
								TempMin:   32.07,
								TempMax:   32.07,
								Pressure:  1018,
								Humidity:  34,
							},
							Visibility: 10000,
							Wind: entity.Wind{
								Speed: 1.34,
								Deg:   303,
							},
							Name: "Kyiv",
						},
						{
							Main: entity.Main{
								Temp:      28.49,
								FeelsLike: 27.57,
								TempMin:   28.49,
								TempMax:   29.29,
								Pressure:  1023,
								Humidity:  32,
							},
							Visibility: 10000,
							Wind: entity.Wind{
								Speed: 3.09,
								Deg:   130,
							},
							Name: "Vilnius",
						},
						{
							Main: entity.Main{
								Temp:      23.08,
								FeelsLike: 23.48,
								TempMin:   22.04,
								TempMax:   24.4,
								Pressure:  1023,
								Humidity:  78,
							},
							Visibility: 10000,
							Wind: entity.Wind{
								Speed: 3.09,
								Deg:   340,
							},
							Name: "Rīga",
						},
					},
				}

				s.On("GetWeatherList", mock.MatchedBy(func(arg []int) bool { return true })).Return(weather, nil)
			},
			expectedCode: 200,
			expectedBody: `[{"city":"Tallinn","temp":20.85,"feels_like":20.77,"temp_min":20.19,"temp_max":22.19,"wind_speed":3.09,"wind_deg":330},{"city":"Sofia","temp":29.83,"feels_like":28.3,"temp_min":28.91,"temp_max":29.83,"wind_speed":3.09,"wind_deg":60},{"city":"Minsk","temp":26.86,"feels_like":26.58,"temp_min":26.86,"temp_max":26.86,"wind_speed":4.25,"wind_deg":70},{"city":"Kyiv","temp":32.07,"feels_like":31.41,"temp_min":32.07,"temp_max":32.07,"wind_speed":1.34,"wind_deg":303},{"city":"Vilnius","temp":28.49,"feels_like":27.57,"temp_min":28.49,"temp_max":29.29,"wind_speed":3.09,"wind_deg":130},{"city":"Rīga","temp":23.08,"feels_like":23.48,"temp_min":22.04,"temp_max":24.4,"wind_speed":3.09,"wind_deg":340}]`,
		},
		{
			name: "len resp != len cities",
			cities: map[string]int{
				"kyiv": 703448, "vilnius": 593116,
				"riga": 456173, "tallinn": 588409,
				"sofia": 727011, "minsk": 625144,
			},
			mockBehavior: func(s *mocks.IService) {
				weather := entity.Weather{
					List: []entity.List{},
				}

				s.On("GetWeatherList", mock.MatchedBy(func(arg []int) bool { return true })).Return(weather, nil)
			},
			expectedCode: 400,
			expectedBody: `{"error":"city not found"}`,
		},
		{
			name:         "bad request map",
			cities:       map[string]int{},
			mockBehavior: func(s *mocks.IService) {},
			expectedCode: 400,
			expectedBody: `{"error":"empty cities map"}`,
		},
		{
			name:         "invalid api key",
			cities: map[string]int{
				"kyiv": 703448, "vilnius": 593116,
				"riga": 456173, "tallinn": 588409,
				"sofia": 727011, "minsk": 625144,
			},
			mockBehavior: func(s *mocks.IService) {
				s.On("GetWeatherList",  mock.MatchedBy(func(arg []int) bool { return true })).
				Return(entity.Weather{}, fmt.Errorf("Invalid API key. Please see https://openweathermap.org/faq#error401 for more info."))
			},
			expectedCode: 500,
			expectedBody: `{"error":"Invalid API key. Please see https://openweathermap.org/faq#error401 for more info."}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {

			mockService := new(mocks.IService)
			testCase.mockBehavior(mockService)

			h := handler{
				s:      mockService,
				cities: testCase.cities,
			}

			e := echo.New()
			h.InitRoutes(e)

			req := httptest.NewRequest("GET", "/weather", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			err := h.weather(c)
			assert.NoError(t, err)
			assert.Equal(t, testCase.expectedCode, rec.Code)
			assert.JSONEq(t, testCase.expectedBody, rec.Body.String())

			mockService.AssertExpectations(t)
		})
	}
}

func TestHandler_weatherByCity(t *testing.T) {
	type mockBehavior func(s *mocks.IService)

	testTable := []struct {
		name         string
		city         string
		cities       map[string]int
		mockBehavior mockBehavior
		expectedCode int
		expectedBody string
	}{
		{
			name: "OK",
			city: "kyiv",
			cities: map[string]int{
				"kyiv": 703448, "vilnius": 593116,
				"riga": 456173, "tallinn": 588409,
				"sofia": 727011, "minsk": 625144,
			},
			mockBehavior: func(s *mocks.IService) {
				weather := entity.Weather{
					List: []entity.List{
						{
							Main: entity.Main{
								Temp:      30.4,
								FeelsLike: 29.8,
								TempMin:   29.81,
								TempMax:   30.4,
								Pressure:  1019,
								Humidity:  37,
							},
							Visibility: 10000,
							Wind: entity.Wind{
								Speed: 0.89,
								Deg:   307,
							},
							Name: "Kyiv",
						},
					},
				}

				s.On("GetWeather", 703448).Return(weather, nil)
			},
			expectedCode: 200,
			expectedBody: `{"city":"Kyiv","temp":30.4,"feels_like":29.8,"temp_min":29.81,"temp_max":30.4,"wind_speed":0.89,"wind_deg":307}`,
		},
		{
			name: "bad request api",
			city: "kremenchuk",
			cities: map[string]int{
				"kyiv": 703448, "vilnius": 593116,
				"riga": 456173, "tallinn": 588409,
				"sofia": 727011, "minsk": 625144,
				"kremenchuk": 1111,
			},
			mockBehavior: func(s *mocks.IService) {
				weather := entity.Weather{
					List: []entity.List{},
				}

				s.On("GetWeather", 1111).Return(weather, nil)
			},
			expectedCode: 400,
			expectedBody: `{"error":"city not found"}`,
		},
		{
			name:         "bad request map",
			city:         "kremenchuk",
			cities:       map[string]int{"riga": 456173},
			mockBehavior: func(s *mocks.IService) {},
			expectedCode: 400,
			expectedBody: `{"error":"city not found"}`,
		},
		{
			name:         "bad param in service",
			city:         "kremenchuk",
			cities: map[string]int{
				"kyiv": 703448, "vilnius": 593116,
				"riga": 456173, "tallinn": 588409,
				"sofia": 727011, "minsk": 625144,
				"kremenchuk": 1111,
			},
			mockBehavior: func(s *mocks.IService) {
				s.On("GetWeather", 1111).Return(entity.Weather{}, fmt.Errorf("No data: 404006"))
			},
			expectedCode: 500,
			expectedBody: `{"error":"No data: 404006"}`,
		},
		{
			name:         "invalid api key",
			city:         "kyiv",
			cities: map[string]int{
				"kyiv": 703448, "vilnius": 593116,
				"riga": 456173, "tallinn": 588409,
				"sofia": 727011, "minsk": 625144,
			},
			mockBehavior: func(s *mocks.IService) {
				s.On("GetWeather",  mock.MatchedBy(func(arg int) bool { return true })).
				Return(entity.Weather{}, fmt.Errorf("Invalid API key. Please see https://openweathermap.org/faq#error401 for more info."))
			},
			expectedCode: 500,
			expectedBody: `{"error":"Invalid API key. Please see https://openweathermap.org/faq#error401 for more info."}`,
		},
		{
			name:         "empty weather list",
			city:         "kyiv",
			cities: map[string]int{
				"kyiv": 703448, "vilnius": 593116,
				"riga": 456173, "tallinn": 588409,
				"sofia": 727011, "minsk": 625144,
			},
			mockBehavior: func(s *mocks.IService) {
				s.On("GetWeather",  mock.MatchedBy(func(arg int) bool { return true })).
				Return(entity.Weather{}, nil)
			},
			expectedCode: 400,
			expectedBody: `{"error":"city not found"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {

			mockService := new(mocks.IService)
			testCase.mockBehavior(mockService)

			h := handler{
				s:      mockService,
				cities: testCase.cities,
			}

			e := echo.New()
			h.InitRoutes(e)

			req := httptest.NewRequest("GET", fmt.Sprintf("/weather/%s", testCase.city), nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetParamNames("city")
			c.SetParamValues(testCase.city)

			err := h.weatherByCity(c)
			assert.NoError(t, err)
			assert.Equal(t, testCase.expectedCode, rec.Code)
			assert.JSONEq(t, testCase.expectedBody, rec.Body.String())

			mockService.AssertExpectations(t)
		})
	}
}

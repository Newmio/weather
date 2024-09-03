package repository

import (
	"fmt"
	"testing"
	"weather/internal/domain/entity"
	cachemock "weather/internal/repository/cache/mocks"
	httpmock "weather/internal/repository/http/mocks"

	"github.com/stretchr/testify/assert"
)

var weatherKyiv = entity.Weather{
	List: []entity.List{
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
			Id:   703448,
		},
	},
}

var weatherRiga = entity.Weather{
	List: []entity.List{
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
			Id:   456173,
		},
	},
}

var weatherTallin = entity.Weather{
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
			Id:   588409,
		},
	},
}

var weatherSofia = entity.Weather{
	List: []entity.List{
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
			Id:   727011,
		},
	},
}

var weatherMinsk = entity.Weather{
	List: []entity.List{
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
			Id:   625144,
		},
	},
}

var weatherVilnius = entity.Weather{
	List: []entity.List{
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
			Id:   593116,
		},
	},
} 

var weathers = entity.Weather{
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
			Id:   588409,
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
			Id:   727011,
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
			Id:   625144,
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
			Id:   703448,
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
			Id:   593116,
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
			Id:   456173,
		},
	},
}

func TestManager_GetForecast(t *testing.T) {
	type mockBehaviourCache func(c *cachemock.ICache)
	type mockBehaviourHttp func(c *httpmock.IHttp)

	testTable := []struct {
		name               string
		city string
		mockBehaviourHttp  mockBehaviourHttp
		mockBehaviourCache mockBehaviourCache
		expected           entity.Weather
		expectedErr        error
	}{

		{
			name:   "success cache",
			city:   "tallinn",
			mockBehaviourCache: func(c *cachemock.ICache) {
				c.On("GetForecast", "tallinn").Return(weathers, nil)
			},
			mockBehaviourHttp: func(c *httpmock.IHttp) {},
			expected:          weathers,
			expectedErr:       nil,
		},
		{
			name:   "success http",
			city:   "tallinn",
			mockBehaviourCache: func(c *cachemock.ICache) {
				c.On("GetForecast", "tallinn").Return(entity.Weather{}, nil)
				c.On("SetForecast", "tallinn", weathers)
			},
			mockBehaviourHttp: func(c *httpmock.IHttp) {
				c.On("GetForecast", "tallinn").Return(weathers, nil)
			},
			expected:          weathers,
			expectedErr:       nil,
		},
		{
			name:   "invalid api key",
			city:   "tallinn",
			mockBehaviourCache: func(c *cachemock.ICache) {
				c.On("GetForecast", "tallinn").Return(entity.Weather{}, nil)
			},
			mockBehaviourHttp: func(h *httpmock.IHttp) {
				h.On("GetForecast", "tallinn").
				Return(entity.Weather{}, fmt.Errorf("Invalid API key. Please see https://openweathermap.org/faq#error401 for more info."))
			},
			expected:    entity.Weather{},
			expectedErr: fmt.Errorf("Invalid API key. Please see https://openweathermap.org/faq#error401 for more info."),
		},
		{
			name:   "not found city",
			city:   "tallinn11111",
			mockBehaviourCache: func(c *cachemock.ICache) {
				c.On("GetForecast", "tallinn11111").Return(entity.Weather{}, nil)
			},
			mockBehaviourHttp: func(h *httpmock.IHttp) {
				h.On("GetForecast",  "tallinn11111").
				Return(entity.Weather{}, fmt.Errorf("No data: 404006"))
			},
			expected:    entity.Weather{},
			expectedErr: fmt.Errorf("No data: 404006"),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {

			mockHttp := new(httpmock.IHttp)
			mockCache := new(cachemock.ICache)
			testCase.mockBehaviourHttp(mockHttp)
			testCase.mockBehaviourCache(mockCache)

			manager := repo{
				http:  mockHttp,
				cache: mockCache,
			}

			result, err := manager.GetForecast(testCase.city)

			if testCase.expectedErr != nil {
				assert.Error(t, err)
				assert.Equal(t, testCase.expectedErr.Error(), err.Error())
			} else {
				assert.NoError(t, err)
				assert.ElementsMatch(t, testCase.expected.List, result.List)
			}

			mockHttp.AssertExpectations(t)
			mockCache.AssertExpectations(t)
		})
	}
}

func TestManager_GetWeatherList(t *testing.T) {
	type mockBehaviourCache func(c *cachemock.ICache)
	type mockBehaviourHttp func(c *httpmock.IHttp)

	testTable := []struct {
		name               string
		citiesId           []int
		mockBehaviourHttp  mockBehaviourHttp
		mockBehaviourCache mockBehaviourCache
		expected           entity.Weather
		expectedErr        error
	}{
		{
			name:   "success cache",
			citiesId: []int{588409, 727011, 625144, 703448, 593116, 456173},
			mockBehaviourCache: func(c *cachemock.ICache) {
				c.On("GetWeather", 588409).Return(weatherTallin, nil)
				c.On("GetWeather", 727011).Return(weatherSofia, nil)
				c.On("GetWeather", 625144).Return(weatherMinsk, nil)
				c.On("GetWeather", 703448).Return(weatherKyiv, nil)
				c.On("GetWeather", 593116).Return(weatherVilnius, nil)
				c.On("GetWeather", 456173).Return(weatherRiga, nil)
			},
			mockBehaviourHttp: func(h *httpmock.IHttp) {},
			expected:          weathers,
			expectedErr:       nil,
		},
		{
			name:   "success http",
			citiesId: []int{588409, 727011, 625144, 703448, 593116, 456173},
			mockBehaviourCache: func(c *cachemock.ICache) {
				c.On("GetWeather", 588409).Return(entity.Weather{}, nil)
				c.On("GetWeather", 727011).Return(entity.Weather{}, nil)
				c.On("GetWeather", 625144).Return(entity.Weather{}, nil)
				c.On("GetWeather", 703448).Return(entity.Weather{}, nil)
				c.On("GetWeather", 593116).Return(entity.Weather{}, nil)
				c.On("GetWeather", 456173).Return(entity.Weather{}, nil)

				c.On("SetWeather", 588409, entity.Weather{List: weatherTallin.List})
				c.On("SetWeather", 727011, entity.Weather{List: weatherSofia.List})
				c.On("SetWeather", 625144, entity.Weather{List: weatherMinsk.List})
				c.On("SetWeather", 703448, entity.Weather{List: weatherKyiv.List})
				c.On("SetWeather", 593116, entity.Weather{List: weatherVilnius.List})
				c.On("SetWeather", 456173, entity.Weather{List: weatherRiga.List})
			},
			mockBehaviourHttp: func(h *httpmock.IHttp) {
				h.On("GetWeatherList", []int{588409, 727011, 625144, 703448, 593116, 456173}).Return(weathers, nil)
			},
			expected:    weathers,
			expectedErr: nil,
		},
		{
			name:   "invalid api key",
			citiesId: []int{588409, 727011, 625144, 703448, 593116, 456173},
			mockBehaviourCache: func(c *cachemock.ICache) {
				c.On("GetWeather", 588409).Return(entity.Weather{}, nil)
				c.On("GetWeather", 727011).Return(entity.Weather{}, nil)
				c.On("GetWeather", 625144).Return(entity.Weather{}, nil)
				c.On("GetWeather", 703448).Return(entity.Weather{}, nil)
				c.On("GetWeather", 593116).Return(entity.Weather{}, nil)
				c.On("GetWeather", 456173).Return(entity.Weather{}, nil)
			},
			mockBehaviourHttp: func(h *httpmock.IHttp) {
				h.On("GetWeatherList", []int{588409, 727011, 625144, 703448, 593116, 456173}).
				Return(entity.Weather{}, fmt.Errorf("Invalid API key. Please see https://openweathermap.org/faq#error401 for more info."))
			},
			expected:    entity.Weather{},
			expectedErr: fmt.Errorf("Invalid API key. Please see https://openweathermap.org/faq#error401 for more info."),
		},
		{
			name:   "not found city",
			citiesId: []int{588409, 727011, 625144, 703448, 593116, 456173},
			mockBehaviourCache: func(c *cachemock.ICache) {
				c.On("GetWeather", 588409).Return(entity.Weather{}, nil)
				c.On("GetWeather", 727011).Return(entity.Weather{}, nil)
				c.On("GetWeather", 625144).Return(entity.Weather{}, nil)
				c.On("GetWeather", 703448).Return(entity.Weather{}, nil)
				c.On("GetWeather", 593116).Return(entity.Weather{}, nil)
				c.On("GetWeather", 456173).Return(entity.Weather{}, nil)
			},
			mockBehaviourHttp: func(h *httpmock.IHttp) {
				h.On("GetWeatherList", []int{588409, 727011, 625144, 703448, 593116, 456173}).
				Return(entity.Weather{}, fmt.Errorf("No data: 404006"))
			},
			expected:    entity.Weather{},
			expectedErr: fmt.Errorf("No data: 404006"),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {

			mockHttp := new(httpmock.IHttp)
			mockCache := new(cachemock.ICache)
			testCase.mockBehaviourHttp(mockHttp)
			testCase.mockBehaviourCache(mockCache)

			manager := repo{
				http:  mockHttp,
				cache: mockCache,
			}

			result, err := manager.GetWeatherList(testCase.citiesId)

			if testCase.expectedErr != nil {
				assert.Error(t, err)
				assert.Equal(t, testCase.expectedErr.Error(), err.Error())
			} else {
				assert.NoError(t, err)
				assert.ElementsMatch(t, testCase.expected.List, result.List)
			}

			mockHttp.AssertExpectations(t)
			mockCache.AssertExpectations(t)
		})
	}
}

func TestManager_GetWeather(t *testing.T) {
	type mockBehaviourCache func(c *cachemock.ICache)
	type mockBehaviourHttp func(c *httpmock.IHttp)

	testTable := []struct {
		name               string
		cityId             int
		mockBehaviourHttp  mockBehaviourHttp
		mockBehaviourCache mockBehaviourCache
		expected           entity.Weather
		expectedErr        error
	}{
		{
			name:   "success cache",
			cityId: 703448,
			mockBehaviourCache: func(c *cachemock.ICache) {
				c.On("GetWeather", 703448).Return(weatherKyiv, nil)
			},
			mockBehaviourHttp: func(h *httpmock.IHttp) {},
			expected:          weatherKyiv,
			expectedErr:       nil,
		},
		{
			name:   "success http",
			cityId: 703448,
			mockBehaviourCache: func(c *cachemock.ICache) {
				c.On("GetWeather", 703448).Return(entity.Weather{}, nil)
				c.On("SetWeather", 703448, weatherKyiv)
			},
			mockBehaviourHttp: func(h *httpmock.IHttp) {
				h.On("GetWeather", 703448).Return(weatherKyiv, nil)
			},
			expected:    weatherKyiv,
			expectedErr: nil,
		},
		{
			name:   "invalid api key",
			cityId: 703448,
			mockBehaviourCache: func(c *cachemock.ICache) {
				c.On("GetWeather", 703448).Return(entity.Weather{}, nil)
			},
			mockBehaviourHttp: func(h *httpmock.IHttp) {
				h.On("GetWeather", 703448).Return(entity.Weather{}, 
					fmt.Errorf("Invalid API key. Please see https://openweathermap.org/faq#error401 for more info."))
			},
			expected:    entity.Weather{},
			expectedErr: fmt.Errorf("Invalid API key. Please see https://openweathermap.org/faq#error401 for more info."),
		},
		{
			name:   "not found city",
			cityId: 703448,
			mockBehaviourCache: func(c *cachemock.ICache) {
				c.On("GetWeather", 703448).Return(entity.Weather{}, nil)
			},
			mockBehaviourHttp: func(h *httpmock.IHttp) {
				h.On("GetWeather", 703448).Return(entity.Weather{}, fmt.Errorf("No data: 404006"))
			},
			expected:    entity.Weather{},
			expectedErr: fmt.Errorf("No data: 404006"),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {

			mockHttp := new(httpmock.IHttp)
			mockCache := new(cachemock.ICache)
			testCase.mockBehaviourHttp(mockHttp)
			testCase.mockBehaviourCache(mockCache)

			manager := repo{
				http:  mockHttp,
				cache: mockCache,
			}

			result, err := manager.GetWeather(testCase.cityId)

			if testCase.expectedErr != nil {
				assert.Error(t, err)
				assert.Equal(t, testCase.expectedErr.Error(), err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.expected, result)
			}

			mockHttp.AssertExpectations(t)
			mockCache.AssertExpectations(t)
		})
	}
}

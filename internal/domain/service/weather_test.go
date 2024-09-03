package service

import (
	"fmt"
	"testing"
	"weather/internal/domain/entity"
	"weather/internal/repository/mocks"

	"github.com/stretchr/testify/assert"
)

func TestService_GetAverage(t *testing.T) {
	type mockBehavior func(r *mocks.IRepository)

	testTable := []struct {
		name         string
		cities       []string
		mockBehavior mockBehavior
		expected     []entity.WeatherAverage
		ecpectedErr  error
	}{
		{
			name:   "success",
			cities: []string{"kyiv", "vilnius"},
			mockBehavior: func(r *mocks.IRepository) {

				kyiv := entity.Weather{
					List: []entity.List{
						{
							Main: entity.Main{
								TempMin:  10,
								TempMax:  20,
								Humidity: 30,
							},
						},
						{
							Main: entity.Main{
								TempMin:  20,
								TempMax:  30,
								Humidity: 40,
							},
						},
						{
							Main: entity.Main{
								TempMin:  30,
								TempMax:  40,
								Humidity: 50,
							},
						},
					},
				}

				vilnius := entity.Weather{
					List: []entity.List{
						{
							Main: entity.Main{
								TempMin:  40,
								TempMax:  50,
								Humidity: 60,
							},
						},
						{
							Main: entity.Main{
								TempMin:  70,
								TempMax:  80,
								Humidity: 90,
							},
						},
						{
							Main: entity.Main{
								TempMin:  100,
								TempMax:  110,
								Humidity: 100,
							},
						},
					},
				}

				r.On("GetForecast", "kyiv").Return(kyiv, nil)
				r.On("GetForecast", "vilnius").Return(vilnius, nil)
			},
			expected: []entity.WeatherAverage{
				{
					City:     "kyiv",
					Temp:     25,
					Humidity: 40,
				},
				{
					City:     "vilnius",
					Temp:     75,
					Humidity: 83,
				},
			},
			ecpectedErr: nil,
		},
		{
			name:   "not found",
			cities: []string{"dnepr"},
			mockBehavior: func(r *mocks.IRepository) {
				r.On("GetForecast", "dnepr").Return(entity.Weather{}, fmt.Errorf("No data: 404006"))
			},
			expected:    nil,
			ecpectedErr: fmt.Errorf("No data: 404006"),
		},
		{
			name:   "invalid api key",
			cities: []string{"kyiv"},
			mockBehavior: func(r *mocks.IRepository) {
				r.On("GetForecast", "kyiv").Return(entity.Weather{}, fmt.Errorf("Invalid API key. Please see https://openweathermap.org/faq#error401 for more info."))
			},
			expected:    nil,
			ecpectedErr: fmt.Errorf("Invalid API key. Please see https://openweathermap.org/faq#error401 for more info."),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {

			mockRepo := new(mocks.IRepository)
			testCase.mockBehavior(mockRepo)

			s := service{r: mockRepo}

			result, err := s.GetAverage(testCase.cities)

			if testCase.ecpectedErr != nil {
				assert.Error(t, err)
				assert.Equal(t, testCase.ecpectedErr.Error(), err.Error())
			} else {
				assert.NoError(t, err)
				assert.ElementsMatch(t, testCase.expected, result)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

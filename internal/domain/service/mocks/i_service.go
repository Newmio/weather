// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import (
	entity "weather/internal/domain/entity"

	mock "github.com/stretchr/testify/mock"
)

// IService is an autogenerated mock type for the IService type
type IService struct {
	mock.Mock
}

// GetAverage provides a mock function with given fields: cities
func (_m *IService) GetAverage(cities []string) ([]entity.WeatherAverage, error) {
	ret := _m.Called(cities)

	if len(ret) == 0 {
		panic("no return value specified for GetAverage")
	}

	var r0 []entity.WeatherAverage
	var r1 error
	if rf, ok := ret.Get(0).(func([]string) ([]entity.WeatherAverage, error)); ok {
		return rf(cities)
	}
	if rf, ok := ret.Get(0).(func([]string) []entity.WeatherAverage); ok {
		r0 = rf(cities)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.WeatherAverage)
		}
	}

	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(cities)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetWeather provides a mock function with given fields: cityId
func (_m *IService) GetWeather(cityId int) (entity.Weather, error) {
	ret := _m.Called(cityId)

	if len(ret) == 0 {
		panic("no return value specified for GetWeather")
	}

	var r0 entity.Weather
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (entity.Weather, error)); ok {
		return rf(cityId)
	}
	if rf, ok := ret.Get(0).(func(int) entity.Weather); ok {
		r0 = rf(cityId)
	} else {
		r0 = ret.Get(0).(entity.Weather)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(cityId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetWeatherList provides a mock function with given fields: citiesId
func (_m *IService) GetWeatherList(citiesId []int) (entity.Weather, error) {
	ret := _m.Called(citiesId)

	if len(ret) == 0 {
		panic("no return value specified for GetWeatherList")
	}

	var r0 entity.Weather
	var r1 error
	if rf, ok := ret.Get(0).(func([]int) (entity.Weather, error)); ok {
		return rf(citiesId)
	}
	if rf, ok := ret.Get(0).(func([]int) entity.Weather); ok {
		r0 = rf(citiesId)
	} else {
		r0 = ret.Get(0).(entity.Weather)
	}

	if rf, ok := ret.Get(1).(func([]int) error); ok {
		r1 = rf(citiesId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIService creates a new instance of IService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIService(t interface {
	mock.TestingT
	Cleanup(func())
}) *IService {
	mock := &IService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

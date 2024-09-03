package app

import (
	"github.com/spf13/viper"
)

type Config struct {
	Cities map[string]int
	Token  string
}

func getCities() (Config, error) {
	v := viper.New()
	v.AddConfigPath("configs")
	v.SetConfigName("config")

	if err := v.ReadInConfig(); err != nil {
		return Config{}, err
	}

	cities := make(map[string]int)

	for key, value := range v.GetStringMap("cities") {
		cities[key] = value.(int)
	}

	return Config{
		Cities: cities,
		Token:  v.GetString("token"),
	}, nil
}

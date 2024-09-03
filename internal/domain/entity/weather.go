package entity

type WeatherAverage struct {
	City     string `json:"city"`
	Temp     int    `json:"temp"`
	Humidity int    `json:"humidity"`
}

type Weather struct {
	List []List `json:"list"`
}

type List struct {
	Main       Main   `json:"main"`
	Visibility int    `json:"visibility"`
	Wind       Wind   `json:"wind"`
	Name       string `json:"name"`
	Id         int    `json:"id"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
}

package main

type WeatherData struct {
	City        string `json:"city"`
	Temperature string `json:"temperature"`
	Condition   string `json:"condition"`
	Humidity    string `json:"humidity"`
	WindSpeed   string `json:"windSpeed"`
	Sunrise     string `json:"sunrise"`
	Sunset      string `json:"sunset"`
	Icon        string `json:"icon"`
}

type ExternalWeatherResponse struct {
	ResolvedAddress   string `json:"resolvedAddress"`
	CurrentConditions struct {
		Temp       float64 `json:"temp"`
		Conditions string  `json:"conditions"`
		Humidity   float64 `json:"humidity"`
		Windspeed  float64 `json:"windspeed"`
		Sunrise    string  `json:"sunrise"`
		Sunset     string  `json:"sunset"`
		Icon       string  `json:"icon"`
	} `json:"currentConditions"`
}
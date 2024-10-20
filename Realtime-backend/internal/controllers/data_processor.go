package controllers

type ProcessedData struct {
	City      string
	Main      string
	Temp      float64
	FeelsLike float64
	Timestamp int64
}

func Process(apiResp WeatherAPIResponse, unit string) *ProcessedData {
	temp, feelsLike := apiResp.Main.Temp, apiResp.Main.FeelsLike
	if unit == "Celsius" {
		temp = kelvinToCelsius(temp)
		feelsLike = kelvinToCelsius(feelsLike)
	} else if unit == "Fahrenheit" {
		temp = kelvinToFahrenheit(temp)
		feelsLike = kelvinToFahrenheit(feelsLike)
	}
	return &ProcessedData{
		City:      apiResp.Name,
		Main:      apiResp.Weather[0].Main,
		Temp:      temp,
		FeelsLike: feelsLike,
		Timestamp: apiResp.Dt,
	}
}

func kelvinToCelsius(k float64) float64 {
	return k - 273.15
}

func kelvinToFahrenheit(k float64) float64 {
	return (k-273.15)*9/5 + 32
}

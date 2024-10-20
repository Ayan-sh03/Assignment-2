package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"realtime-weather-agg/internal/config"
	"time"
)

type WeatherAPIResponse struct {
	Name    string `json:"name"`
	Weather []struct {
		Main string `json:"main"`
	} `json:"weather"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
	} `json:"main"`
	Dt int64 `json:"dt"`
}

func FetchWeatherData(cfg *config.Config) {
	for _, city := range cfg.Cities {
		go func(city string) {
			url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, cfg.APIKey)
			resp, err := http.Get(url)
			if err != nil {
				log.Printf("Error fetching data for %s: %v\n", city, err)
				return
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Printf("Error reading response for %s: %v\n", city, err)
				return
			}

			var apiResp WeatherAPIResponse
			if err := json.Unmarshal(body, &apiResp); err != nil {
				log.Printf("Error unmarshalling response for %s: %v\n", city, err)
				return
			}

			processedData := Process(apiResp, cfg.TemperatureUnit)
			log.Printf("Fetched data for %s: %+v\n", city, processedData)
			AddData(processedData)
			CheckAlerts(processedData, cfg)
		}(city)
	}
}

func StartFetching(ctx context.Context, cfg *config.Config) {

	ticker := time.NewTicker(time.Duration(30 * time.Second))
	defer ticker.Stop()

	log.Println("Fetching weather data...")

	FetchWeatherData(cfg) // Initial fetch

	for {
		select {
		case <-ticker.C:
			FetchWeatherData(cfg)
		}
	}
}

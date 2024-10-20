package controllers

import (
	"log"
	"realtime-weather-agg/internal/models"
	"realtime-weather-agg/internal/utils"
	"sync"
)

type CityData struct {
	Temps      []float64
	Conditions map[string]int
}

var (
	dataStore = make(map[string]map[string]*CityData) // city -> date -> data
	dataMutex     = &sync.Mutex{}
)

func AddData(data *ProcessedData) {
	dataMutex.Lock()
	defer dataMutex.Unlock()

	date := utils.FormatDate(data.Timestamp)

	if _, exists := dataStore[data.City]; !exists {
		dataStore[data.City] = make(map[string]*CityData)
	}

	if _, exists := dataStore[data.City][date]; !exists {
		dataStore[data.City][date] = &CityData{
			Temps:      []float64{},
			Conditions: make(map[string]int),
		}
	}

	cityDateData := dataStore[data.City][date]
	cityDateData.Temps = append(cityDateData.Temps, data.Temp)
	cityDateData.Conditions[data.Main]++
}

func SummarizeDaily() {
	mutex.Lock()
	defer mutex.Unlock()

	today := utils.GetTodayDate()
	for city, dates := range dataStore {
		if cityData, exists := dates[today]; exists {
			temps := cityData.Temps
			if len(temps) == 0 {
				continue
			}

			avgTemp := utils.CalculateAverage(temps)
			maxTemp := utils.CalculateMax(temps)
			minTemp := utils.CalculateMin(temps)
			dominantCondition := utils.GetDominantCondition(cityData.Conditions)

			summary := &models.WeatherSummary{
				City:               city,
				Date:               today,
				AverageTemperature: avgTemp,
				MaxTemperature:     maxTemp,
				MinTemperature:     minTemp,
				DominantCondition:  dominantCondition,
			}

			if err := summary.Save(); err != nil {
				log.Printf("Error saving summary for %s on %s: %v\n", city, today, err)
			} else {
				log.Printf("Saved summary for %s on %s\n", city, today)
			}
		}
	}

	// Reset dataStore for the next day
	dataStore = make(map[string]map[string]*CityData)
}

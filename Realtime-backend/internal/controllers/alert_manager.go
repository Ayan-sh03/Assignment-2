package controllers

import (
	"fmt"
	"log"
	"realtime-weather-agg/internal/config"
	"realtime-weather-agg/internal/services"
	"sync"
)

type AlertState struct {
	TempCount int
}

var (
	stateStore = make(map[string]*AlertState) // city -> state
	mutex      = &sync.Mutex{}
)

func CheckAlerts(data *ProcessedData, cfg *config.Config) {
	mutex.Lock()
	defer mutex.Unlock()

	// Temperature Alerts

	log.Printf("Checking alerts for %s with temp %v againts %v \n", data.City, data.Temp, cfg.AlertThresholds.MaxTemperature)

	if float64(data.Temp) > float64(cfg.AlertThresholds.MaxTemperature) {

		if _, exists := stateStore[data.City]; !exists {
			stateStore[data.City] = &AlertState{TempCount: 1}
		} else {
			stateStore[data.City].TempCount++
		}

		if stateStore[data.City].TempCount >= cfg.AlertThresholds.Consecutive {
			message := fmt.Sprintf("Temperature in %s has exceeded %.2f°C for %d consecutive updates.", data.City, cfg.AlertThresholds.MaxTemperature, cfg.AlertThresholds.Consecutive)
			log.Println("ALERT:", message)
			triggerAlert(message, cfg)
			// Reset counter after alert
			stateStore[data.City].TempCount = 0
		}
	} else {
		if _, exists := stateStore[data.City]; exists {
			stateStore[data.City].TempCount = 0
		}
	}

	// Condition Alerts
	for _, condition := range cfg.AlertThresholds.Conditions {
		if data.Main == condition {
			message := fmt.Sprintf("Weather condition in %s is %s.", data.City, data.Main)
			triggerAlert(message, cfg)
		}
	}
}

func triggerAlert(message string, cfg *config.Config) {
	log.Println("ALERT:", message)
	err := services.SendEmail(cfg.EmailConfig, "Weather Alert", message)
	if err != nil {
		log.Printf("Failed to send alert email: %v\n", err)
	}
}

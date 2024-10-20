package models

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"realtime-weather-agg/internal/db"

	"github.com/lib/pq"
)

type WeatherConfig struct {
	Cities                     []string    `json:"cities"`
	ThresholdTemparature       json.Number `json:"threshold_temperature"`
	Email                      string      `json:"email"`
	ConsecutiveAlertThreshhold json.Number `json:"consecutive_alert_threshold"`
}

func (wc *WeatherConfig) Save() error {
	query := `
	INSERT INTO weather_configs (cities, threshold_temperature, email, consecutive_alert_threshold)
	VALUES ($1, $2, $3, $4)

`
	_, err := db.DB.Exec(context.Background(), query,
		pq.Array(wc.Cities),
		wc.ThresholdTemparature,
		wc.Email,
		wc.ConsecutiveAlertThreshhold)

	if err != nil {
		log.Println("Error Inserting into DB ", err)
		return err
	}

	log.Printf("Saving configuration: %+v\n", wc)
	return nil
}

func GetConfig() (*WeatherConfig, error) {
	query := `
        SELECT cities, threshold_temperature, email, consecutive_alert_threshold
        FROM weather_configs
    `

	row := db.DB.QueryRow(context.Background(), query)

	var wc WeatherConfig
	// Temporary variables to hold the numeric values
	var threshold float64
	var alertThreshold int

	// Scan into temporary variables for numeric fields
	err := row.Scan(&wc.Cities, &threshold, &wc.Email, &alertThreshold)
	if err != nil {
		log.Println("Error scanning row", err)
		return nil, err
	}

	// Convert the numeric values to json.Number
	wc.ThresholdTemparature = json.Number(fmt.Sprintf("%v", threshold))
	wc.ConsecutiveAlertThreshhold = json.Number(fmt.Sprintf("%v", alertThreshold))

	return &wc, nil
}

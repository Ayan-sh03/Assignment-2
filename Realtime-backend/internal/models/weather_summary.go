package models

import (
	"context"
	"log"
	"realtime-weather-agg/internal/db"
	"time"
)

type WeatherSummary struct {
	City               string  `json:"city"`
	Date               string  `json:"date"`
	AverageTemperature float64 `json:"average_temperature"`
	MaxTemperature     float64 `json:"max_temperature"`
	MinTemperature     float64 `json:"min_temperature"`
	DominantCondition  string  `json:"dominant_condition"`
}

type AlertCount struct {
	City  string `json:"city"`
	Count int    `json:"count"`
}

func (ws *WeatherSummary) Save() error {
	query := `
        INSERT INTO weather_summaries (city, date, data)
        VALUES ($1, $2, $3)
        ON CONFLICT (city, date)
        DO UPDATE SET data = EXCLUDED.data
    `
	data := map[string]interface{}{
		"average_temperature": ws.AverageTemperature,
		"max_temperature":     ws.MaxTemperature,
		"min_temperature":     ws.MinTemperature,
		"dominant_condition":  ws.DominantCondition,
	}

	_, err := db.DB.Exec(context.Background(), query, ws.City, ws.Date, data)
	if err != nil {
		log.Printf("Error saving weather summary: %v\n", err)
		return err
	}
	return nil
}

func GetSummary(city, date string) ([]*WeatherSummary, error) {
	query := `
    SELECT date, data
    FROM weather_summaries
    WHERE city = $1 AND date >= $2::date - INTERVAL '6 days' AND date <= $2::date
    ORDER BY date DESC
`

	rows, err := db.DB.Query(context.Background(), query, city, date)
	if err != nil {
		log.Printf("Error fetching summary: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var summaries []*WeatherSummary

	for rows.Next() {
		var ws WeatherSummary
		var data map[string]interface{}
		var summaryDate time.Time // To store the date returned from the query

		// Scanning the date and the data (JSON) from the row
		err := rows.Scan(&summaryDate, &data)
		if err != nil {
			log.Printf("Error scanning row: %v\n", err)
			return nil, err
		}

		// Parsing the data into the WeatherSummary struct fields
		ws.City = city
		ws.Date = summaryDate.Format("2006-01-02") // Format date to a standard string (yyyy-mm-dd)

		// Safely parse and assign data fields, with type assertions to avoid runtime panics
		if avgTemp, ok := data["average_temperature"].(float64); ok {
			ws.AverageTemperature = avgTemp
		} else {
			log.Printf("Warning: Missing or invalid average_temperature for city: %s, date: %s", city, summaryDate)
		}

		if maxTemp, ok := data["max_temperature"].(float64); ok {
			ws.MaxTemperature = maxTemp
		} else {
			log.Printf("Warning: Missing or invalid max_temperature for city: %s, date: %s", city, summaryDate)
		}

		if minTemp, ok := data["min_temperature"].(float64); ok {
			ws.MinTemperature = minTemp
		} else {
			log.Printf("Warning: Missing or invalid min_temperature for city: %s, date: %s", city, summaryDate)
		}

		if condition, ok := data["dominant_condition"].(string); ok {
			ws.DominantCondition = condition
		} else {
			log.Printf("Warning: Missing or invalid dominant_condition for city: %s, date: %s", city, summaryDate)
		}

		// Append the parsed WeatherSummary to the result list
		summaries = append(summaries, &ws)
	}

	if len(summaries) == 0 {
		return nil, nil
	}

	return summaries, nil
}

func GetAlertCount() ([]AlertCount, error) {

	//get alert count from alerts table for all cities
	query := `
				SELECT city_name, SUM(alert_count) as alert_count
				FROM alerts
	  		GROUP BY city_name
				ORDER BY city_name
	`

	rows, err := db.DB.Query(context.Background(), query)
	if err != nil {
		log.Printf("Error fetching alert count: %v\n", err)
		return nil, err
	}

	var alerts []AlertCount
	defer rows.Close()
	for rows.Next() {
		var alert AlertCount
		err := rows.Scan(&alert.City, &alert.Count)
		if err != nil {
			log.Printf("Error scanning row: %v\n", err)
			return nil, err
		}
		alerts = append(alerts, alert)
	}

	return alerts, nil

}

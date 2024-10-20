package config

import (
	"log"
	"os"
	"realtime-weather-agg/internal/models"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	APIKey          string
	MailTrapAPIKey  string
	FetchInterval   int // in seconds
	Cities          []string
	TemperatureUnit string
	AlertThresholds AlertThresholds
	EmailConfig     EmailData
	DatabaseURL     string
}

type AlertThresholds struct {
	MaxTemperature float64
	Consecutive    int
	Conditions     []string
}

type EmailAddress struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type EmailData struct {
	From     EmailAddress   `json:"from"`
	To       []EmailAddress `json:"to"`
	Subject  string         `json:"subject"`
	TextBody string         `json:"text"`
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}

	//

	fetchInterval, err := strconv.Atoi(getEnv("FETCH_INTERVAL", "300"))
	if err != nil {
		fetchInterval = 300
	}

	config := &Config{
		APIKey:          getEnv("OPENWEATHERMAP_API_KEY", ""),
		MailTrapAPIKey:  getEnv("MAILTRAP_API_KEY", ""),
		FetchInterval:   fetchInterval,
		TemperatureUnit: getEnv("TEMPERATURE_UNIT", "Celsius"),
		AlertThresholds: AlertThresholds{
			Conditions: getEnvAsSlice("ALERT_CONDITIONS", "Rain,Snow"),
		},
		EmailConfig: EmailData{
			From: EmailAddress{
				Email: getEnv("EMAIL_FROM", ""),
				Name:  getEnv("EMAIL_FROM_NAME", ""),
			},
		},
		DatabaseURL: getEnv("DATABASE_URL", ""),
	}

	log.Printf("Config: %+v\n", config)

	return config
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		log.Printf("Loaded %s from environment variable\n", key)
		return value
	}
	log.Printf("Using fallback value for %s: %s\n", key, fallback)
	return fallback
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultVal
}

func getEnvAsFloat(name string, defaultVal float64) float64 {
	valueStr := getEnv(name, "")
	if value, err := strconv.ParseFloat(valueStr, 64); err == nil {
		return value
	}
	return defaultVal
}

func getEnvAsSlice(name string, defaultVal string) []string {
	valueStr := getEnv(name, defaultVal)
	return splitAndTrim(valueStr, ",")
}

func splitAndTrim(s string, sep string) []string {
	var result []string
	for _, part := range split(s, sep) {
		trimmed := trim(part)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
}

func split(s string, sep string) []string {
	return strings.Split(s, sep)
}

func trim(s string) string {
	return strings.TrimSpace(s)
}

func (cfg *Config) SetWeatherConfig() error {
	dbConfig, err := models.GetConfig()
	if err != nil {
		log.Println("Error fetching config from db:", err)
		return err
	}

	if dbConfig != nil {

		thresholdTemparature, err := dbConfig.ThresholdTemparature.Float64()
		if err != nil {
			log.Println("Error fetching config from db:", err)
			return err
		}

		consecutiveAlert, err := dbConfig.ConsecutiveAlertThreshhold.Int64()
		if err != nil {
			log.Println("Error fetching config from db:", err)
			return err
		}

		cfg.AlertThresholds.MaxTemperature = thresholdTemparature

		cfg.Cities = dbConfig.Cities
		cfg.AlertThresholds.Consecutive = int(consecutiveAlert)
	}

	return nil

}

package services

import (
	"fmt"
	"log"
	"realtime-weather-agg/internal/config"

	"github.com/go-resty/resty/v2"
)

func SendEmail(emailCfg config.EmailData, subject, body string) error {
	client := resty.New()

	// Define the email structure

	emailCfg.Subject = subject
	if subject == "" {
		emailCfg.Subject = "Weather Alert, One or More weather threshold crossed"
	}

	emailCfg.TextBody = body

	// Send the request
	resp, err := client.R().
		SetHeader("Authorization", "Bearer "+config.LoadConfig().MailTrapAPIKey).
		SetHeader("Content-Type", "application/json").
		SetBody(emailCfg).
		Post("https://send.api.mailtrap.io/api/send")

	// Handle response
	if err != nil {
		log.Fatalf("Error sending email: %v", err)
	}

	// Check the status code
	if resp.StatusCode() == 200 {
		fmt.Println("Email sent successfully!")
	} else {
		fmt.Printf("Failed to send email. Status: %s\n", resp.Status())
		fmt.Println("Response: ", resp)
	}
	return err
}

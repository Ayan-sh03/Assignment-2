package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"realtime-weather-agg/internal/config"
	"realtime-weather-agg/internal/controllers"
	"realtime-weather-agg/internal/db"
	"realtime-weather-agg/internal/visualization"
	"syscall"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load Configuration
	cfg := config.LoadConfig()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Initialize Database
	db.InitDB(cfg.DatabaseURL)
	defer db.CloseDB()

	//set configuration from DB
	cfg.SetWeatherConfig()

	// Initialize Gin Router
	router := gin.Default()
	router.Use(cors.Default())
	visualization.SetupRoutes(router)

	// Start Data Fetching in a Goroutine
	go controllers.StartFetching(ctx, cfg)

	// Setup Scheduler for Daily Summaries
	// c := cron.New()
	// c.AddFunc("@midnight", func() {
	// 	aggregator.SummarizeDaily()
	// })
	// c.Start()
	// defer c.Stop()

	// Start Server in a Goroutine
	go func() {
		if err := router.Run(":8080"); err != nil {
			log.Fatalf("Failed to run server: %v", err)
		}
	}()

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")
}

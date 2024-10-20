package visualization

import (
	"net/http"
	"realtime-weather-agg/internal/models"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/summary/:city/:date", getSummary)
	router.GET("/alerts", GetAlerts)
	router.POST("/config", createConfig)
	router.GET("/config", getConfig)

	//TODO :add setting config route

}

func GetAlerts(c *gin.Context) {
	alerts, err := models.GetAlertCount()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch alerts"})
		return
	}
	c.JSON(http.StatusOK, alerts)

}

func getSummary(c *gin.Context) {
	city := c.Param("city")
	date := c.Param("date")

	summary, err := models.GetSummary(city, date)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Summary not found"})
		return
	}

	c.JSON(http.StatusOK, summary)
}

func createConfig(c *gin.Context) {
	var wc models.WeatherConfig

	if err := c.BindJSON(&wc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse JSON: " + err.Error()})
		return
	}

	// Save the configuration
	if err := wc.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save configuration"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Configuration saved successfully"})
}

func getConfig(x *gin.Context) {
	cfg, err := models.GetConfig()
	if err != nil {
		x.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch configuration"})
		return
	}
	x.JSON(http.StatusOK, cfg)
}

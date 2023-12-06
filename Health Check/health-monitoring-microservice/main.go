package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello from the health monitoring microservice!"})
	})

	go Monitor()

	err := router.Run(":8080")
	if err != nil {
		fmt.Printf("Error starting the server: %v\n", err)
	}
}

func Monitor() {
	ticker := time.NewTicker(5 * time.Second)

	for {
		select {
		case <-ticker.C:
			// Perform health checks by querying the Payment Service.
			if !isPaymentServiceHealthy() {
				// Trigger an alert if the service is unhealthy.
				fmt.Println("ALERT: Payment Service is unhealthy! Check logs for details.")
			} else {
				fmt.Println("ALERT: Payment Service is healthy!")
			}
		}
	}
}

func isPaymentServiceHealthy() bool {
	resp, err := http.Get("http://payment:8080/payment-api/v1/health")
	if err != nil || resp.StatusCode != http.StatusOK {
		return false
	}
	return true
}

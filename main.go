package main

import (
	"fetch-rewards-receipt-processor-challenge/config"
	"fetch-rewards-receipt-processor-challenge/routes"
	"log"
	"github.com/gin-gonic/gin"
)


func main () {
	config.LoadConfig()
	ginMode := config.GetEnv("GIN_MODE")
	if ginMode != "" {
			gin.SetMode(ginMode)
	}
	router := gin.Default()
	routes.SetupRoutes(router)
	port := config.GetEnv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Printf("Starting server on port %s in %s mode", port, ginMode)
    router.Run(": " + port)
}
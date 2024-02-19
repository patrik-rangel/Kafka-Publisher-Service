package main

import (
	"github.com/HunCoding/golang-architecture/hexagonal-news-api/configuration/logger"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/patrik-rangel/Kafka-Publisher-Service/blob/main/kafka-publiser-service-api/adapter/input/routes"
)

func main() {
	godotenv.Load()
	logger.Info("About to start application")
	router := gin.Default()

	routes.InitRoutes(router)

	if err := router.Run(":8080"); err != nil {
		logger.Error("Error trying to init application on port 8080", err)
		return
	}
}

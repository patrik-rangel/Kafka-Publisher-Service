package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/patrik-rangel/Kafka-Publisher-Service/blob/main/kafka-publiser-service-api/adapter/input/controller"
	"github.com/patrik-rangel/Kafka-Publisher-Service/blob/main/kafka-publiser-service-api/adapter/output/purchase_kafka"
	"github.com/patrik-rangel/Kafka-Publisher-Service/blob/main/kafka-publiser-service-api/application/service"
	"github.com/patrik-rangel/Kafka-Publisher-Service/blob/main/kafka-publiser-service-api/configuration/env"
)

func InitRoutes(r *gin.Engine) {
	addr := env.GetAddressKafka()
	topic := env.GetTopicKafka()
	purchaseOutput := purchase_kafka.NewPurchaseProducer(addr, topic)
	purchaseService := service.NewPurchaseService(purchaseOutput)
	purchaseController := controller.NewPurchaseController(purchaseService)

	r.POST("/purchase", purchaseController.PostPurchase)
}

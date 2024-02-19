package controller

import (
	"net/http"

	"github.com/HunCoding/golang-architecture/hexagonal-news-api/configuration/logger"
	"github.com/HunCoding/golang-architecture/hexagonal-news-api/configuration/validation"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/patrik-rangel/Kafka-Publisher-Service/blob/main/kafka-publiser-service-api/adapter/input/model/request"
	"github.com/patrik-rangel/Kafka-Publisher-Service/blob/main/kafka-publiser-service-api/application/domain"
	"github.com/patrik-rangel/Kafka-Publisher-Service/blob/main/kafka-publiser-service-api/application/port/input"
)

type purchaseController struct {
	purchaseUseCase input.PurchaseUseCase
}

func NewPurchaseController (purchaseUseCase input.PurchaseUseCase,) *purchaseController {
	return &purchaseController{purchaseUseCase}
}

func (pc *purchaseController) PostPurchase (c *gin.Context) {
	logger.Info("Init PostPurchase Controller API")
	purchaseRequest := request.PurchaseRequest{}

	if err := c.ShouldBindJSON(&purchaseRequest); err != nil {
		logger.Error("Error trying to validate fields from request", err)
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	purchaseDomainCopy := domain.PurchaseDomain{}
	copier.Copy(&purchaseDomainCopy, &purchaseRequest)

	purchaseResponseDoamin, err := pc.purchaseUseCase.PostPurchaseService(purchaseDomainCopy)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, purchaseResponseDoamin)
}
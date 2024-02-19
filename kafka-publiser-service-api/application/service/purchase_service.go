package service

import (
	"fmt"

	"github.com/HunCoding/golang-architecture/hexagonal-news-api/configuration/logger"
	"github.com/patrik-rangel/Kafka-Publisher-Service/blob/main/kafka-publiser-service-api/application/domain"
	"github.com/patrik-rangel/Kafka-Publisher-Service/blob/main/kafka-publiser-service-api/application/port/output"
)

type purchaseService struct {
	purchasePort output.PurchasePort
}

func NewPurchaseService(purchasePort output.PurchasePort) *purchaseService {
	return &purchaseService{purchasePort}
}

func (ps *purchaseService) PostPurchaseService(purchaseDomain domain.PurchaseDomain,
) (bool, error) {
	logger.Info(
		fmt.Sprintf("init PostPurchaseService function"),
	)

	err := ps.purchasePort.PostPurchase(purchaseDomain)
	if err != nil {
		return false, err
	}

	return true, nil
}

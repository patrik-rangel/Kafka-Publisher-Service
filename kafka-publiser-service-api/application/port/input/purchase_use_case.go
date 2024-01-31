package input

import (
	"github.com/patrik-rangel/Kafka-Publisher-Service/kafka-publisher-service-api/application/domain"
	"github.com/HunCoding/golang-architecture/hexagonal-news-api/configuration/rest_err"
)

type PurchaseUseCase interface {
	PostPurchase(domain.PurchaseDomain) (*domain.PurchaseDomain, *rest_err.RestErr)
}
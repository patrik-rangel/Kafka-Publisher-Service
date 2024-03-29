package input

import (
	"github.com/patrik-rangel/Kafka-Publisher-Service/blob/main/kafka-publiser-service-api/application/domain"
)

type PurchaseUseCase interface {
	PostPurchaseService(domain.PurchaseDomain) (bool, error)
}

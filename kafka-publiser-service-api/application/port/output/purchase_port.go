package output

import (
	"github.com/patrik-rangel/Kafka-Publisher-Service/blob/main/kafka-publiser-service-api/application/domain"
)

type PurchasePort interface {
	PostPurchase(domain.PurchaseDomain) error
}

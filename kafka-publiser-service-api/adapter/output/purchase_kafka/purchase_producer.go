package purchase_kafka

import (
	"context"
	"log"

	"github.com/patrik-rangel/Kafka-Publisher-Service/blob/main/kafka-publiser-service-api/application/domain"

	"github.com/segmentio/kafka-go"
)

type purchaseProducer struct {
	address string
	topic   string
}

func NewPurchaseProducer(address string, topic string) *purchaseProducer {
	return &purchaseProducer{
		address: address,
		topic:   topic,
	}
}

func (pp *purchaseProducer) PostPurchase(domain.PurchaseDomain) error {
	w := &kafka.Writer{
		Addr:  kafka.TCP(pp.address),
		Topic: pp.topic,
	}

	err := w.WriteMessages(
		context.Background(),
		kafka.Message{Value: []byte("Hello kafka!")},
	)

	if err != nil {
		log.Fatal("Failed to send message", err)
		return err
	}

	w.Close()

	return nil
}

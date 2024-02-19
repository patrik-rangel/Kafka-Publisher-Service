package env

import "os"

func GetAddressKafka() string {
	return os.Getenv("ADDR_KAFKA")
}

func GetTopicKafka() string {
	return os.Getenv("TOPIC")
}

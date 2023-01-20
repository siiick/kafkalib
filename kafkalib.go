package kafkalib

import (
	"fmt"

	"github.com/Shopify/sarama"
)

func init() {
	fmt.Println("kafkalib init")
}

func NewConfig() *sarama.Config {
	config := sarama.NewConfig()
	config.Version = sarama.V3_2_1_0
	return config
}

func NewKafkaLib(config *sarama.Config) (sarama.Consumer, error) {
	return sarama.NewConsumer([]string{"localhost:9092"}, config)
}

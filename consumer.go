package kafkalib

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/Shopify/sarama"
)

var consumer sarama.Consumer

func init() {
	// This will get called on main initialization
	fmt.Println("This will get called on main initialization. (consumer)")

	conf := sarama.NewConfig()
	conf.Version = sarama.V3_2_1_0

	var err error
	consumer, err = sarama.NewConsumer([]string{"localhost:9092"}, conf)
	if err != nil {
		panic(err)
	}
}

func Hello(topic string, partition int32) {
	fmt.Println("This will not get called on main initialization.(consumer)")

	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	partitionConsumer, err := consumer.ConsumePartition(topic, partition, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	// Trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	consumed := 0

ConsumerLoop:

	for {
		select {
		case msg := <-partitionConsumer.Messages():
			log.Printf("Consumed message offset %d\n%s\n", msg.Offset, msg.Value)
			consumed++
		case <-signals:
			break ConsumerLoop
		}
	}
}

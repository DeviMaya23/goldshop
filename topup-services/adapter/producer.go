package adapter

import (
	"devimaya/goldshop/topupservices/internal/config"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var kafkaProducer *kafka.Producer

func init() {

	var err error

	conf := config.GetConfig()
	bootstrapServer := conf.KafkaHost + ":" + conf.KafkaPort
	kafkaProducer, err = kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": bootstrapServer})
	// kafkaProducer, err = kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	if err != nil {
		fmt.Printf("Failed to create producer: %s", err)
		panic(err)
	}
	go func() {
		for e := range kafkaProducer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Failed to deliver message: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Produced event to topic %s: key = %-10s value = %s\n",
						*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))
				}
			}
		}
	}()
}

func Produce(topic string, key string, data []byte) error {

	err := kafkaProducer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            []byte(key),
		Value:          data,
	}, nil)

	return err

}

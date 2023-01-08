package main

import (
	"devimaya/goldshop/inputhargastorage/adapter/db/model"
	"devimaya/goldshop/inputhargastorage/adapter/db/repository"
	"devimaya/goldshop/inputhargastorage/internal/config"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {

	conf := config.GetConfig()
	bootstrapServer := conf.KafkaHost + ":" + conf.KafkaPort

	kafkaConsumer, err := kafka.NewConsumer(&kafka.ConfigMap{"bootstrap.servers": bootstrapServer, "group.id": conf.KafkaGroupId})

	if err != nil {
		fmt.Printf("Failed to create Consumer: %s", err)
		panic(err)
	}

	topic := "input-harga"
	_ = kafkaConsumer.SubscribeTopics([]string{topic}, nil)
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	run := true
	for run {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			ev, err := kafkaConsumer.ReadMessage(100 * time.Millisecond)
			if err != nil {
				continue
			}

			var request *model.Request
			json.Unmarshal(ev.Value, &request)
			repository.SetHarga(request)
			fmt.Printf("Consumed event from topic %s: key = %-10s value = %s\n",
				*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))

		}
	}

	kafkaConsumer.Close()

}

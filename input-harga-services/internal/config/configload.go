package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	ServiceHost string
	ServicePort string
	KafkaHost   string
	KafkaPort   string
}

var cfg config

func init() {
	cfg = config{}
	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg.ServiceHost = os.Getenv("SERVICEHOST")
	cfg.ServicePort = os.Getenv("SERVICEPORT")
	cfg.KafkaHost = os.Getenv("KAFKAHOST")
	cfg.KafkaPort = os.Getenv("KAFKAPORT")

}
func GetConfig() *config {
	return &cfg
}

package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	ServiceHost        string
	ServicePort        string
	KafkaHost          string
	KafkaPort          string
	CheckSaldoEndpoint string
	CheckHargaEndpoint string
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
	cfg.CheckSaldoEndpoint = os.Getenv("CHECKSALDOENDPOINT")
	cfg.CheckHargaEndpoint = os.Getenv("CHECKHARGAENDPOINT")

}
func GetConfig() *config {
	return &cfg
}

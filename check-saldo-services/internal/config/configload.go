package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	ServiceHost string
	ServicePort string
	DBHost      string
	DBPort      string
	DBUsername  string
	DBPassword  string
	DBName      string
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
	cfg.DBHost = os.Getenv("DBHOST")
	cfg.DBPort = os.Getenv("DBPORT")
	cfg.DBUsername = os.Getenv("DBUSER")
	cfg.DBPassword = os.Getenv("DBPASSWORD")
	cfg.DBName = os.Getenv("DBNAME")

}
func GetConfig() *config {
	return &cfg
}

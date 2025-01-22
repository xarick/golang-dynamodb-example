package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Application struct {
	GinMode string
	RunPort string

	DynamoDBEndPoint string
	DynamoDBRegion   string
	ProxyURL         string
}

func LoadConfig() Application {
	cfg := Application{}

	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env file not found")
	}

	cfg.GinMode = os.Getenv("GIN_MODE")
	cfg.RunPort = os.Getenv("RUN_PORT")

	cfg.DynamoDBEndPoint = os.Getenv("DYNAMODB_ENDPOINT")
	cfg.DynamoDBRegion = os.Getenv("DYNAMODB_REGION")
	cfg.ProxyURL = os.Getenv("PROXY_URL")

	return cfg
}

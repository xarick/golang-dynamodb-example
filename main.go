package main

import (
	"log"

	"github.com/xarick/golang-dynamodb-example/config"
	"github.com/xarick/golang-dynamodb-example/db"
	"github.com/xarick/golang-dynamodb-example/routes"
)

func main() {
	cfg := config.LoadConfig()

	db.InitDynamoDB(cfg)

	r := routes.SetupRoutes()

	if err := r.Run(cfg.RunPort); err != nil {
		log.Fatalf("Serverda xatolik: %v", err)
	}
}

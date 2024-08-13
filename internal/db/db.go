package db

import (
	"log"
	"os"

	"github.com/anandu86130/Microservice_gRPC_product/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn, ok := os.LookupEnv("DSN")

	if !ok {
		log.Fatal("error loading database env")
	}

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database %v", err.Error())
	}

	DB.AutoMigrate(model.Product{})
	return DB
}

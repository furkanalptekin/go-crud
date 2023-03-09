package database

import (
	"fmt"
	"log"
	"os"

	"github.com/furkanalptekin/go-crud/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Instance *gorm.DB

func Connect() {
	connectionString := fmt.Sprintf("host=db user=%s password=%s dbname= %s port=5432 sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"))

	var err error
	Instance, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to the db \n", err)
		os.Exit(2)
	}

	Instance.Logger = logger.Default.LogMode(logger.Info)

	Instance.AutoMigrate(&models.Post{})
}

package database

import (
	"log"
	"os"

	"github.com/SurajSmidms/GOLANG_MICROSERVICE_PROJECT/payment-service/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := os.Getenv("PAYMENT_DB_URL")
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to Payment database: %v", err)
	}
	DB.AutoMigrate(&models.Payment{})
}
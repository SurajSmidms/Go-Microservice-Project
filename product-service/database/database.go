package database

import (
	"log"
	"os"

	"github.com/SurajSmidms/GOLANG_MICROSERVICE_PROJECT/product-service/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := os.Getenv("PRODUCT_DB_URL")
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to Product database: %v", err)
	}
	DB.AutoMigrate(&models.Product{})
}
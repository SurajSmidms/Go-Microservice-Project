package main

import (
	"log"
	"net/http"

	"github.com/SurajSmidms/GOLANG_MICROSERVICE_PROJECT/payment-service/database"
	"github.com/SurajSmidms/GOLANG_MICROSERVICE_PROJECT/payment-service/routes"
)

func main() {
	database.ConnectDB()

	r := routes.SetupPaymentRouter()

	log.Println("Payment Service running on port 8084...")
	if err := http.ListenAndServe(":8084", r); err != nil {
		log.Fatalf("Payment service failed: %v", err)
	}
}
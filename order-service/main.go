package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/SurajSmidms/GOLANG_MICROSERVICE_PROJECT/order-service/database"
	"github.com/SurajSmidms/GOLANG_MICROSERVICE_PROJECT/order-service/handlers"
)

func main() {
	database.ConnectDB()

	r := chi.NewRouter()
	r.Get("/api/orders", handlers.GetOrders)
	r.Post("/api/orders", handlers.CreateOrder)
	r.Get("/api/orders/{id}", handlers.GetOrderByID)

	log.Println("Order Service running on port 8083...")
	if err := http.ListenAndServe(":8083", r); err != nil {
		log.Fatalf("Order service failed: %v", err)
	}
}
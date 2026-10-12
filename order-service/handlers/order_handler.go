package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/SurajSmidms/GOLANG_MICROSERVICE_PROJECT/order-service/client"
	"github.com/SurajSmidms/GOLANG_MICROSERVICE_PROJECT/order-service/database"
	"github.com/SurajSmidms/GOLANG_MICROSERVICE_PROJECT/order-service/models"
)

func GetOrders(w http.ResponseWriter, r *http.Request) {
	var orders []models.Order
	database.DB.Find(&orders)
	json.NewEncoder(w).Encode(orders)
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Fetch product info & stock via gRPC from Product Service
	productInfo, err := client.FetchProductDetails(int32(order.ProductID))
	if err != nil {
		http.Error(w, "Failed to communicate with product service via gRPC", http.StatusInternalServerError)
		return
	}

	if productInfo.Stock < int32(order.Quantity) {
		http.Error(w, "Insufficient product stock", http.StatusBadRequest)
		return
	}

	order.TotalPrice = float64(productInfo.Price) * float64(order.Quantity)
	order.Status = "PENDING"

	database.DB.Create(&order)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}

func GetOrderByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var order models.Order
	if result := database.DB.First(&order, id); result.Error != nil {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(order)
}
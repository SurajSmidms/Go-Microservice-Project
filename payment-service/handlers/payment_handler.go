package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/SurajSmidms/GOLANG_MICROSERVICE_PROJECT/payment-service/database"
	"github.com/SurajSmidms/GOLANG_MICROSERVICE_PROJECT/payment-service/models"
)

func ProcessPayment(w http.ResponseWriter, r *http.Request) {
	var payment models.Payment
	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	payment.Status = "SUCCESS"
	database.DB.Create(&payment)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(payment)
}

func GetPaymentStatus(w http.ResponseWriter, r *http.Request) {
	orderID := chi.URLParam(r, "order_id")
	var payment models.Payment
	if result := database.DB.Where("order_id = ?", orderID).First(&payment); result.Error != nil {
		http.Error(w, "Payment record not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(payment)
}
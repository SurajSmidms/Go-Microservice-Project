package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/SurajSmidms/GOLANG_MICROSERVICE_PROJECT/payment-service/handlers"
	"github.com/SurajSmidms/GOLANG_MICROSERVICE_PROJECT/payment-service/middleware"
)

func SetupPaymentRouter() http.Handler {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Use(middleware.AuthMiddleware)
		r.Post("/api/payments", handlers.ProcessPayment)
		r.Get("/api/payments/{order_id}", handlers.GetPaymentStatus)
	})

	return r
}
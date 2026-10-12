package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/SurajSmidms/GOLANG_MICROSERVICE_PROJECT/auth-service/database"
	"github.com/SurajSmidms/GOLANG_MICROSERVICE_PROJECT/auth-service/handlers"
)

func main() {
	database.ConnectDB()

	r := chi.NewRouter()
	r.Post("/api/auth/register", handlers.Register)
	r.Post("/api/auth/login", handlers.Login)

	log.Println("Auth Service running on port 8081...")
	if err := http.ListenAndServe(":8081", r); err != nil {
		log.Fatalf("Auth service failed: %v", err)
	}
}
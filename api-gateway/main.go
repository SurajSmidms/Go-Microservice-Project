package main

import (
	"log"
	"net/http"
	"os"

	"github.com/SurajSmidms/GOLANG_MICROSERVICE_PROJECT/api-gateway/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := routes.SetupRouter()

	log.Printf("API Gateway running on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("API Gateway failed to start: %v", err)
	}
}
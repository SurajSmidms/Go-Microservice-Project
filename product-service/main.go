package main

import (
	"log"
	"net"
	"net/http"

	"github.com/SurajSmidms/GOLANG_MICROSERVICE_PROJECT/product-service/database"
	"github.com/SurajSmidms/GOLANG_MICROSERVICE_PROJECT/product-service/handlers"
	"github.com/SurajSmidms/GOLANG_MICROSERVICE_PROJECT/product-service/server"
	pb "github.com/SurajSmidms/GOLANG_MICROSERVICE_PROJECT/proto/product"
	"github.com/go-chi/chi/v5"
	"google.golang.org/grpc"
)

func main() {
	database.ConnectDB()

	// 1. Start gRPC Server for Internal Service-to-Service Comms
	go func() {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("Failed to listen on gRPC port: %v", err)
		}
		grpcServer := grpc.NewServer()
		pb.RegisterProductServiceServer(grpcServer, &server.ProductGRPCServer{})
		log.Println("Product gRPC Server running on port 50051...")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("gRPC server failed: %v", err)
		}
	}()

	// 2. Start Chi HTTP Server for REST APIs
	r := chi.NewRouter()
	r.Get("/api/products", handlers.GetProducts)
	r.Post("/api/products", handlers.CreateProduct)
	
	r.Route("/api/products/{id}", func(r chi.Router) {
		r.Get("/", handlers.GetProductByID)
		r.Put("/", handlers.UpdateProduct)
		r.Delete("/", handlers.DeleteProduct)
	})

	log.Println("Product REST Service running on port 8082...")
	if err := http.ListenAndServe(":8082", r); err != nil {
		log.Fatalf("HTTP server failed: %v", err)
	}
}
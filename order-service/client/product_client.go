package client

import (
	"context"
	"os"
	"time"

	pb "github.com/SurajSmidms/GOLANG_MICROSERVICE_PROJECT/proto/product"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func FetchProductDetails(productID int32) (*pb.GetProductResponse, error) {
	grpcHost := os.Getenv("PRODUCT_GRPC_HOST")
	if grpcHost == "" {
		grpcHost = "localhost:50051"
	}

	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := pb.NewProductServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	return client.GetProduct(ctx, &pb.GetProductRequest{Id: productID})
}
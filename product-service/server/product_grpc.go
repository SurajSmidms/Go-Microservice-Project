package server

import (
	"context"

	"github.com/SurajSmidms/GOLANG_MICROSERVICE_PROJECT/product-service/database"
	"github.com/SurajSmidms/GOLANG_MICROSERVICE_PROJECT/product-service/models"
	pb "github.com/SurajSmidms/GOLANG_MICROSERVICE_PROJECT/proto/product"
)

type ProductGRPCServer struct {
	pb.UnimplementedProductServiceServer
}

func (s *ProductGRPCServer) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	var product models.Product
	if result := database.DB.First(&product, req.Id); result.Error != nil {
		return nil, result.Error
	}

	return &pb.GetProductResponse{
		Id:    int32(product.ID),
		Name:  product.Name,
		Price: float32(product.Price),
		Stock: int32(product.Stock),
	}, nil
}
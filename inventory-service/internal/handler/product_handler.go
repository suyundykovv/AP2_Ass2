package handler

import (
	"inventory-service/internal/domain"

	pb "github.com/suyundykovv/protos/gen/go/inventory/v1"
)

func ProtoToDomainProduct(req *pb.CreateProductRequest) *domain.Product {
	return &domain.Product{
		Name:       req.GetName(),
		CategoryID: req.GetCategoryId(),
		Price:      float64(req.GetPrice()),
		Stock:      int(req.GetStock()),
	}
}
func DomainToProtoProduct(product *domain.Product) *pb.Product {
	return &pb.Product{
		Id:         product.ID,
		Name:       product.Name,
		CategoryId: product.CategoryID,
		Price:      float32(product.Price),
		Stock:      int64(product.Stock),
	}
}

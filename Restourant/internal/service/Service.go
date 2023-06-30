package service

import (
	"FinalTaskFromMediaSoft/Restourant/internal/database"
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/restaurant"
	"golang.org/x/net/context"
	"time"
)

type Service struct {
	//retailpb.UnimplementedProductServiceServer
	restaurant.UnimplementedProductServiceServer

	//retailpb.Un
	rep database.Rep
}

func New(rep database.Rep) *Service {
	return &Service{
		rep: rep,
	}
}

func (s *Service) CreateProduct(ctx context.Context, request *restaurant.CreateProductRequest) (*restaurant.CreateProductResponse, error) {
	model := database.Product{
		Name:        request.Name,
		Description: request.Description,
		ProductType: request.Type.String(),
		Weight:      request.Weight,
		Price:       request.Price,
		CreatedAt:   time.Now(),
	}
	s.rep.CreateProduct(ctx, &model)
	return &restaurant.CreateProductResponse{}, nil
}

func (s *Service) GetProductList(ctx context.Context, request *restaurant.GetProductListRequest) (*restaurant.GetProductListResponse, error) {
	var products []*database.Product
	products, _ = s.rep.GetProductList(ctx)
	result := []*restaurant.Product{}
	for _, p := range products {
		result = append(result, &restaurant.Product{
			Uuid:        p.Id.String(),
			Name:        p.Name,
			Description: p.Description,
			Type:        getProductTypeFromString(p.ProductType),
			Weight:      p.Weight,
			Price:       p.Price,
		})
	}
	return &restaurant.GetProductListResponse{Result: result}, nil

}

func getProductTypeFromString(value string) restaurant.ProductType {
	switch value {
	case "Salad":
		return 1
	case "Garnish":
		return 2
	case "Meat":
		return 3
	case "Soup":
		return 4
	case "Drink":
		return 5
	case "Dessert":
		return 6
	default:
		return 0
	}
}

//func (s *Service) GetProductTypeList(ctx context.Context, request)

//func (s *Service) CreateMenu(ctx context.Context, request *retailpb.)

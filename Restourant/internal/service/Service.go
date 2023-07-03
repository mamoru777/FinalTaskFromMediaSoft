package service

import (
	"FinalTaskFromMediaSoft/Restourant/internal/database"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/restaurant"
	"golang.org/x/net/context"
	"log"
	"time"
)

type Service struct {
	//retailpb.UnimplementedProductServiceServer
	restaurant.UnimplementedProductServiceServer
	restaurant.UnimplementedMenuServiceServer
	//retailpb.Un
	rep database.Rep
}

func New(rep database.Rep) *Service {
	return &Service{
		rep: rep,
	}
}

func (s *Service) CreateMenu(ctx context.Context, request *restaurant.CreateMenuRequest) (*restaurant.CreateMenuResponse, error) {
	//var timenow timestamp.Timestamp
	/*timenow, error := ptypes.TimestampProto(time.Now())
	if error != nil {
		log.Fatal(error)
	}*/
	var currentTime time.Time = time.Now()
	nextDay := currentTime.AddDate(0, 0, 1)
	modeldb, _ := s.rep.GetMenu(ctx, nextDay)
	if modeldb != nil {
		return &restaurant.CreateMenuResponse{}, nil
	}
	var productsproto []string
	productsproto = append(productsproto, request.Desserts...)
	productsproto = append(productsproto, request.Drinks...)
	productsproto = append(productsproto, request.Meats...)
	productsproto = append(productsproto, request.Garnishes...)
	productsproto = append(productsproto, request.Salads...)
	productsproto = append(productsproto, request.Soups...)
	fmt.Println("var productsproto []string: ", productsproto)
	var products []database.Product
	for _, p := range productsproto {
		products = append(products, s.rep.GetProduct(ctx, p))
	}
	fmt.Println("var products []database.Product: ", products)
	model := database.Menu{
		OnDate:          request.OnDate.AsTime(),
		OpeningRecordAt: request.OpeningRecordAt.AsTime(),
		ClosingRecordAt: request.ClosingRecordAt.AsTime(),
		CreatedAt:       time.Now(),
		Products:        products,
	}
	fmt.Println("Model.Menu: ", model)
	s.rep.CreateMenu(ctx, &model)

	return &restaurant.CreateMenuResponse{}, nil
}

func (s *Service) GetMenu(ctx context.Context, request *restaurant.GetMenuRequest) (*restaurant.GetMenuResponse, error) {
	var currentTime time.Time = time.Now()
	nextDay := currentTime.AddDate(0, 0, 1)
	nextDayProto, error := ptypes.TimestampProto(nextDay)
	if error != nil {
		log.Fatal(error)
	}
	model, err := s.rep.GetMenu(ctx, nextDay)
	if err != nil {
		log.Fatal("Запись не найдена", err)
	}
	fmt.Println(model)
	OpRecAtProto, error := ptypes.TimestampProto(model.OpeningRecordAt)
	if error != nil {
		log.Fatal(error)
	}
	ClRecAtProto, error := ptypes.TimestampProto(model.ClosingRecordAt)
	if error != nil {
		log.Fatal(error)
	}
	CreatedAtProto, error := ptypes.TimestampProto(model.CreatedAt)
	if error != nil {
		log.Fatal(error)
	}
	var SaladsProto []*restaurant.Product    //make([]*restaurant.Product, 0)
	var GarnishesProto []*restaurant.Product //:= make([]*restaurant.Product, 0)
	var MeatsProto []*restaurant.Product     //:= make([]*restaurant.Product, 0)
	var SoupsProto []*restaurant.Product     //:= make([]*restaurant.Product, 0)
	var DrinksProto []*restaurant.Product    //:= make([]*restaurant.Product, 0)
	var DessertsProto []*restaurant.Product  //:= make([]*restaurant.Product, 0)
	/*var ProductsProto []*restaurant.Product
	for _, p := range model.Products {
		ProductsProto = append(ProductsProto, &restaurant.Product{
			Uuid:        p.Id.String(),
			Name:        p.Name,
			Description: p.Description,
			Weight:      p.Weight,
			Price:       p.Price,
			CreatedAt:   CreatedAtProto,
			Type:        3,
		})
	}*/

	for _, p := range model.Products {
		if p.ProductType == "PRODUCT_TYPE_SALAD" {
			CreatedAtProtoS, error := ptypes.TimestampProto(p.CreatedAt)
			if error != nil {
				log.Fatal(error)
			}
			SaladsProto = append(SaladsProto, &restaurant.Product{
				Uuid:        p.Id.String(),
				Name:        p.Name,
				Description: p.Description,
				Weight:      p.Weight,
				Price:       p.Price,
				CreatedAt:   CreatedAtProtoS,
				Type:        1,
			})

		}
		if p.ProductType == "PRODUCT_TYPE_GARNISH" {
			CreatedAtProtoG, error := ptypes.TimestampProto(p.CreatedAt)
			if error != nil {
				log.Fatal(error)
			}
			GarnishesProto = append(GarnishesProto, &restaurant.Product{
				Uuid:        p.Id.String(),
				Name:        p.Name,
				Description: p.Description,
				Weight:      p.Weight,
				Price:       p.Price,
				CreatedAt:   CreatedAtProtoG,
				Type:        2,
			})
		}
		if p.ProductType == "PRODUCT_TYPE_MEAT" {
			CreatedAtProtoM, error := ptypes.TimestampProto(p.CreatedAt)
			if error != nil {
				log.Fatal(error)
			}
			MeatsProto = append(MeatsProto, &restaurant.Product{
				Uuid:        p.Id.String(),
				Name:        p.Name,
				Description: p.Description,
				Weight:      p.Weight,
				Price:       p.Price,
				CreatedAt:   CreatedAtProtoM,
				Type:        3,
			})
		}
		if p.ProductType == "PRODUCT_TYPE_SOUP" {
			CreatedAtProtoSp, error := ptypes.TimestampProto(p.CreatedAt)
			if error != nil {
				log.Fatal(error)
			}
			SoupsProto = append(SoupsProto, &restaurant.Product{
				Uuid:        p.Id.String(),
				Name:        p.Name,
				Description: p.Description,
				Weight:      p.Weight,
				Price:       p.Price,
				CreatedAt:   CreatedAtProtoSp,
				Type:        4,
			})
		}
		if p.ProductType == "PRODUCT_TYPE_DRINK" {
			CreatedAtProtoDr, error := ptypes.TimestampProto(p.CreatedAt)
			if error != nil {
				log.Fatal(error)
			}
			DrinksProto = append(DrinksProto, &restaurant.Product{
				Uuid:        p.Id.String(),
				Name:        p.Name,
				Description: p.Description,
				Weight:      p.Weight,
				Price:       p.Price,
				CreatedAt:   CreatedAtProtoDr,
				Type:        5,
			})
		}
		if p.ProductType == "PRODUCT_TYPE_DESSERT" {
			CreatedAtProtoDs, error := ptypes.TimestampProto(p.CreatedAt)
			if error != nil {
				log.Fatal(error)
			}
			DessertsProto = append(DessertsProto, &restaurant.Product{
				Uuid:        p.Id.String(),
				Name:        p.Name,
				Description: p.Description,
				Weight:      p.Weight,
				Price:       p.Price,
				CreatedAt:   CreatedAtProtoDs,
				Type:        6,
			})
		}
	}
	menu := &restaurant.Menu{
		Uuid:            model.Id.String(),
		OnDate:          nextDayProto,
		OpeningRecordAt: OpRecAtProto,
		ClosingRecordAt: ClRecAtProto,
		CreatedAt:       CreatedAtProto,
		Salads:          SaladsProto,
		Garnishes:       GarnishesProto,
		Meats:           MeatsProto,
		Soups:           SoupsProto,
		Drinks:          DrinksProto,
		Desserts:        DessertsProto,
	}
	fmt.Print(menu)
	return &restaurant.GetMenuResponse{Menu: menu}, nil
}

func (s *Service) CreateProduct(ctx context.Context, request *restaurant.CreateProductRequest) (*restaurant.CreateProductResponse, error) {
	/*timenow, error := ptypes.TimestampProto(time.Now())
	if error != nil {
		log.Fatal(error)
	}*/
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

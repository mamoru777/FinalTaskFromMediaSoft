package service

import (
	"FinalTaskFromMediaSoft/Restourant/internal/database"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	customer2 "gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/customer"
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/restaurant"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

type Service struct {
	restaurant.UnimplementedProductServiceServer
	restaurant.UnimplementedMenuServiceServer
	restaurant.UnimplementedOrderServiceServer
	rep database.Rep
}

func New(rep database.Rep) *Service {
	return &Service{
		rep: rep,
	}
}

func (s *Service) CreateMenu(ctx context.Context, request *restaurant.CreateMenuRequest) (*restaurant.CreateMenuResponse, error) {
	loc, err := time.LoadLocation("UTC")
	if err != nil {
		log.Println("Не удалось задать нужный часовой пояс")
	}
	time.Local = loc
	//var timenow timestamp.Timestamp
	/*timenow, error := ptypes.TimestampProto(time.Now())
	if error != nil {
		log.Fatal(error)
	}*/
	//var currentTime time.Time = time.Now()
	//nextDay := currentTime.AddDate(0, 0, 1)
	//modeldb, _ := s.rep.GetMenu(ctx, nextDay)
	//if modeldb != nil {
	//	return &restaurant.CreateMenuResponse{}, nil
	//}
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
	loc, err := time.LoadLocation("UTC")
	if err != nil {
		log.Println("Не удалось задать нужный часовой пояс")
	}
	time.Local = loc
	//var currentTime time.Time = time.Now()
	//nextDay := currentTime.AddDate(0, 0, 1)
	//nextDayProto, error := ptypes.TimestampProto(nextDay)
	//if error != nil {
	//	log.Fatal(error)
	//}
	model, err := s.rep.GetMenu(ctx, request.OnDate.AsTime())
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
		OnDate:          request.OnDate, //nextDayProto,
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

func (s *Service) GetUpToDateOrderList(ctx context.Context, request *restaurant.GetUpToDateOrderListRequest) (*restaurant.GetUpToDateOrderListResponse, error) {
	orderList, err := s.rep.GetOrderList(ctx)
	if err != nil {
		return nil, err
	}
	var orderListProto []*restaurant.Order
	for _, o := range orderList {
		orderListProto = append(orderListProto, &restaurant.Order{
			ProductId:   o.ProductId.String(),
			ProductName: o.ProductName,
			Count:       o.Count,
		})
	}

	conn, err := grpc.Dial("localhost:13998", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("error connect to grpc server err:", err)
	}
	customer := customer2.NewUserServiceClient(conn)
	office := customer2.NewOfficeServiceClient(conn)
	var orderByOfficeListProto []*restaurant.OrdersByOffice

	res, _ := office.GetOfficeList(context.Background(), &customer2.GetOfficeListRequest{})
	for _, of := range res.Result {
		rescus, _ := customer.GetUserList(context.Background(), &customer2.GetUserListRequest{OfficeUuid: of.Uuid})
		orders := []*database.Order{}
		var ordersProto []*restaurant.Order
		for _, c := range rescus.Result {
			UuidProto, err := uuid.Parse(c.Uuid)
			if err != nil {
				log.Fatal("Не удалось преобразовать строку", err)
			}
			orders, _ = s.rep.GetOrdersByCustomer(context.Background(), UuidProto)
			for _, o := range orders {
				ordersProto = append(ordersProto, &restaurant.Order{
					ProductId:   o.ProductId.String(),
					ProductName: o.ProductName,
					Count:       o.Count,
				})
			}
		}
		orderByOfficeListProto = append(orderByOfficeListProto, &restaurant.OrdersByOffice{
			OfficeUuid:    of.Uuid,
			OfficeName:    of.Name,
			OfficeAddress: of.Address,
			Result:        ordersProto,
		})
	}
	conn.Close()
	return &restaurant.GetUpToDateOrderListResponse{TotalOrders: orderListProto, TotalOrdersByCompany: orderByOfficeListProto}, nil
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

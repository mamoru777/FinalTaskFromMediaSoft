package service

import (
	"FinalTaskFromMediaSoft/Statistics/internal/database"
	"FinalTaskFromMediaSoft/pkg/contracts/restaurant"
	"FinalTaskFromMediaSoft/pkg/contracts/statistics"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"sort"
)

type Service struct {
	/*restaurant.UnimplementedProductServiceServer
	restaurant.UnimplementedMenuServiceServer
	restaurant.UnimplementedOrderServiceServer*/
	statistics.UnsafeStatisticsServiceServer
	rep database.Rep
}

func New(rep database.Rep) *Service {
	return &Service{
		rep: rep,
	}
}

func (s *Service) GetAmountOfProfit(ctx context.Context, request *statistics.GetAmountOfProfitRequest) (*statistics.GetAmountOfProfitResponse, error) {
	conn, err := grpc.Dial("localhost:13999", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("error connect to grpc server err:", err)
	}
	product := restaurant.NewProductServiceClient(conn)
	res, err := product.GetProductList(context.Background(), &restaurant.GetProductListRequest{})
	if err != nil {
		return nil, err
	}
	productList := res.Result
	var profit float64
	orders, err := s.rep.GetOrderList(ctx)
	if err != nil {
		log.Fatal("Не удалось получить список заказов", err)
	}
	for _, o := range orders {
		if o.CreatedAt.After(request.StartDate.AsTime()) && o.CreatedAt.Before(request.EndDate.AsTime()) {
			for _, p := range productList {
				if p.Uuid == o.ProductId.String() {
					countFloat := float64(o.Count)
					profit += countFloat * p.Price
				}
			}
		}
	}
	conn.Close()
	return &statistics.GetAmountOfProfitResponse{Profit: profit}, nil
}

func (s *Service) TopProducts(ctx context.Context, request *statistics.TopProductsRequest) (*statistics.TopProductsResponse, error) {
	conn, err := grpc.Dial("localhost:13999", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("error connect to grpc server err:", err)
	}
	product := restaurant.NewProductServiceClient(conn)
	res, err := product.GetProductList(context.Background(), &restaurant.GetProductListRequest{})
	if err != nil {
		return nil, err
	}
	productList := res.Result
	orders, err := s.rep.GetOrderList(ctx)
	if err != nil {
		log.Fatal("Не удалось получить список заказов", err)
	}
	var productListRes []*statistics.Product
	var isExit bool = false
	var productName string
	var productType string

	for _, o := range orders {
		if o.CreatedAt.After(request.StartDate.AsTime()) && o.CreatedAt.Before(request.EndDate.AsTime()) {
			for _, plr := range productListRes {
				if plr.Uuid == o.ProductId.String() {
					plr.Count += o.Count
					isExit = true
				}
			}
			if isExit == false {
				for _, pl := range productList {
					if o.ProductId.String() == pl.Uuid {
						productName = pl.Name
						productType = pl.Type.String()
					}
				}
				if request.ProductType == nil || getProductTypeFromString(productType) == *request.ProductType {
					productListRes = append(productListRes, &statistics.Product{
						Uuid:        o.ProductId.String(),
						Name:        productName,
						Count:       o.Count,
						ProductType: getProductTypeFromString(productType),
					})
				}
			}
			isExit = false
		}
	}
	compare := func(i, j int) bool {
		return productListRes[i].Count > productListRes[j].Count
	}
	sort.Slice(productListRes, compare)
	conn.Close()
	return &statistics.TopProductsResponse{Result: productListRes}, nil
}

func getProductTypeFromString(value string) statistics.StatisticsProductType {
	switch value {
	case "PRODUCT_TYPE_SALAD":
		return 1
	case "PRODUCT_TYPE_GARNISH":
		return 2
	case "PRODUCT_TYPE_MEAT":
		return 3
	case "PRODUCT_TYPE_SOUP":
		return 4
	case "PRODUCT_TYPE_DRINK":
		return 5
	case "PRODUCT_TYPE_DESSERT":
		return 6
	default:
		return 0
	}
}

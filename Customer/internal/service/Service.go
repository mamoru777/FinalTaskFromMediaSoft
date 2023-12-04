package service

import (
	"FinalTaskFromMediaSoft/Customer/internal/database"
	"FinalTaskFromMediaSoft/pkg/contracts/customer"
	"FinalTaskFromMediaSoft/pkg/contracts/restaurant"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

type Service struct {
	//restaurant.UnimplementedProductServiceServer
	//restaurant.UnimplementedMenuServiceServer
	customer.UnimplementedOfficeServiceServer
	customer.UnimplementedUserServiceServer
	customer.UnimplementedOrderServiceServer
	rep database.Rep
	//product restaurant.NewProductServiceClient()
}

func New(rep database.Rep) *Service {
	return &Service{
		rep: rep,
	}
}

func (s *Service) CreateUser(ctx context.Context, request *customer.CreateUserRequest) (*customer.CreateUserResponse, error) {
	OfficeUuid, err := uuid.Parse(request.OfficeUuid)
	if err != nil {
		log.Fatal("Не удалось преобразовать строку", err)
	}
	var officeName string
	var offices []*database.Office
	offices, _ = s.rep.GetOfficeList(ctx)
	for _, o := range offices {
		if OfficeUuid == o.Id {
			officeName = o.Name
		}
	}
	model := database.Customer{
		Name:       request.Name,
		OfficeId:   OfficeUuid,
		OfficeName: officeName,
		CreatedAt:  time.Now(),
	}
	s.rep.CreateCustomer(ctx, &model)
	return &customer.CreateUserResponse{}, nil
}

func (s *Service) GetUserList(ctx context.Context, request *customer.GetUserListRequest) (*customer.GetUserListResponse, error) {
	var customers []*database.Customer
	OfficeUuidDb, err := uuid.Parse(request.OfficeUuid)
	if err != nil {
		log.Fatal("Не удалось преобразовать строку", err)
	}
	customers, _ = s.rep.GetCustomerList(ctx, OfficeUuidDb)
	result := []*customer.User{}
	for _, c := range customers {
		CreatedAtProto, error := ptypes.TimestampProto(c.CreatedAt)
		if error != nil {
			log.Fatal(error)
		}
		result = append(result, &customer.User{
			Uuid:       c.Id.String(),
			Name:       c.Name,
			OfficeUuid: c.OfficeId.String(),
			OfficeName: c.OfficeName,
			CreatedAt:  CreatedAtProto,
		})
	}
	return &customer.GetUserListResponse{Result: result}, nil
}

func (s *Service) CreateOffice(ctx context.Context, request *customer.CreateOfficeRequest) (*customer.CreateOfficeResponse, error) {
	currentTime := time.Now()
	model := database.Office{
		Name:      request.Name,
		Address:   request.Address,
		CreatedAt: currentTime,
	}
	s.rep.CreateOffice(ctx, &model)
	return &customer.CreateOfficeResponse{}, nil
}

func (s *Service) GetOfficeList(ctx context.Context, request *customer.GetOfficeListRequest) (*customer.GetOfficeListResponse, error) {

	var offices []*database.Office
	offices, _ = s.rep.GetOfficeList(ctx)
	result := []*customer.Office{}
	for _, p := range offices {
		CreatedAtProto, error := ptypes.TimestampProto(p.CreatedAt)
		if error != nil {
			log.Fatal(error)
		}
		result = append(result, &customer.Office{
			Uuid:      p.Id.String(),
			Name:      p.Name,
			Address:   p.Address,
			CreatedAt: CreatedAtProto,
		})
	}
	return &customer.GetOfficeListResponse{Result: result}, nil
}

func (s *Service) GetActualMenu(ctx context.Context, request *customer.GetActualMenuRequest) (*customer.GetActualMenuResponse, error) {
	conn, err := grpc.Dial("localhost:13999", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("error connect to grpc server err:", err)
	}
	//product := restaurant.NewProductServiceClient(conn)
	menu := restaurant.NewMenuServiceClient(conn)
	var currentTime time.Time = time.Now()
	nextDay := currentTime.AddDate(0, 0, 1)
	nextDayProto, error := ptypes.TimestampProto(nextDay)
	if error != nil {
		log.Fatal(error)
	}
	res, _ := menu.GetMenu(context.Background(), &restaurant.GetMenuRequest{OnDate: nextDayProto})
	menuRes := res.Menu
	var Salads []*customer.Product
	for _, s := range menuRes.Salads {
		Salads = append(Salads, &customer.Product{
			Uuid:        s.Uuid,
			Name:        s.Name,
			Description: s.Description,
			Type:        1,
			Weight:      s.Weight,
			Price:       s.Price,
			CreatedAt:   s.CreatedAt,
		})
	}
	var Garnishes []*customer.Product
	for _, g := range menuRes.Garnishes {
		Garnishes = append(Garnishes, &customer.Product{
			Uuid:        g.Uuid,
			Name:        g.Name,
			Description: g.Description,
			Type:        2,
			Weight:      g.Weight,
			Price:       g.Price,
			CreatedAt:   g.CreatedAt,
		})
	}
	var Meats []*customer.Product
	for _, m := range menuRes.Meats {
		Meats = append(Meats, &customer.Product{
			Uuid:        m.Uuid,
			Name:        m.Name,
			Description: m.Description,
			Type:        3,
			Weight:      m.Weight,
			Price:       m.Price,
			CreatedAt:   m.CreatedAt,
		})
	}
	var Soups []*customer.Product
	for _, sp := range menuRes.Soups {
		Soups = append(Soups, &customer.Product{
			Uuid:        sp.Uuid,
			Name:        sp.Name,
			Description: sp.Description,
			Type:        4,
			Weight:      sp.Weight,
			Price:       sp.Price,
			CreatedAt:   sp.CreatedAt,
		})
	}
	var Drinks []*customer.Product
	for _, dr := range menuRes.Drinks {
		Drinks = append(Drinks, &customer.Product{
			Uuid:        dr.Uuid,
			Name:        dr.Name,
			Description: dr.Description,
			Type:        5,
			Weight:      dr.Weight,
			Price:       dr.Price,
			CreatedAt:   dr.CreatedAt,
		})
	}
	var Desserts []*customer.Product
	for _, ds := range menuRes.Desserts {
		Desserts = append(Desserts, &customer.Product{
			Uuid:        ds.Uuid,
			Name:        ds.Name,
			Description: ds.Description,
			Type:        6,
			Weight:      ds.Weight,
			Price:       ds.Price,
			CreatedAt:   ds.CreatedAt,
		})
	}
	return &customer.GetActualMenuResponse{Salads: Salads, Garnishes: Garnishes, Meats: Meats, Soups: Soups, Drinks: Drinks, Desserts: Desserts}, nil
}

func (s *Service) CreateOrder(ctx context.Context, requset *customer.CreateOrderRequest) (*customer.CreateOrderResponse, error) {
	conn, err := grpc.Dial("localhost:13999", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("error connect to grpc server err:", err)
	}
	//product := restaurant.NewProductServiceClient(conn)
	menu := restaurant.NewMenuServiceClient(conn)
	currentTime := time.Now()
	nextDay := currentTime.AddDate(0, 0, 1)
	nextDayProto, error := ptypes.TimestampProto(nextDay)
	if error != nil {
		log.Fatal(error)
	}
	res, _ := menu.GetMenu(context.Background(), &restaurant.GetMenuRequest{OnDate: nextDayProto})
	menuRes := res.Menu
	timeOpen := menuRes.OpeningRecordAt.AsTime()
	timeClose := menuRes.ClosingRecordAt.AsTime()
	if currentTime.Before(timeOpen) || currentTime.After(timeClose) {
		err := error
		return nil, err
	}
	configkafka := sarama.NewConfig()
	configkafka.Producer.Return.Successes = true

	client, err := sarama.NewClient([]string{"localhost:9092"}, configkafka)
	if err != nil {
		log.Fatal("Error creating client: ", err)
	}

	configkafka.Producer.Partitioner = sarama.NewManualPartitioner
	configkafka.Producer.RequiredAcks = sarama.WaitForAll

	producer, err := sarama.NewSyncProducerFromClient(client)
	if err != nil {
		log.Fatal("Failed to create producer", err)
	}

	var Salads []string
	var Garnishes []string
	var Meats []string
	var Soups []string
	var Drinks []string
	var Desserts []string
	for _, p := range requset.Salads {
		Salads = append(Salads, p.String())
	}
	for _, p := range requset.Meats {
		Meats = append(Meats, p.String())
	}
	for _, p := range requset.Garnishes {
		Garnishes = append(Garnishes, p.String())
	}
	for _, p := range requset.Soups {
		Soups = append(Soups, p.String())
	}
	for _, p := range requset.Drinks {
		Drinks = append(Drinks, p.String())
	}
	for _, p := range requset.Desserts {
		Desserts = append(Desserts, p.String())
	}
	for _, str := range Salads {
		msg := &sarama.ProducerMessage{
			Topic: "Order",
			Value: sarama.StringEncoder(str + " customer_uuid:" + requset.UserUuid + " created_at:" + currentTime.String()),
		}
		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Сообщение отправлено в partiton", partition, offset)
	}
	for _, str := range Garnishes {
		msg := &sarama.ProducerMessage{
			Topic: "Order",
			Value: sarama.StringEncoder(str + " customer_uuid:" + requset.UserUuid + " created_at:" + currentTime.String()),
		}
		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Сообщение отправлено в partiton", partition, offset)
	}
	for _, str := range Meats {
		msg := &sarama.ProducerMessage{
			Topic: "Order",
			Value: sarama.StringEncoder(str + " customer_uuid:" + requset.UserUuid + " created_at:" + currentTime.String()),
		}
		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Сообщение отправлено в partiton", partition, offset)
	}
	for _, str := range Soups {
		msg := &sarama.ProducerMessage{
			Topic: "Order",
			Value: sarama.StringEncoder(str + " customer_uuid:" + requset.UserUuid + " created_at:" + currentTime.String()),
		}
		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Сообщение отправлено в partiton", partition, offset)
	}
	for _, str := range Drinks {
		msg := &sarama.ProducerMessage{
			Topic: "Order",
			Value: sarama.StringEncoder(str + " customer_uuid:" + requset.UserUuid + " created_at:" + currentTime.String()),
		}
		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Сообщение отправлено в partiton", partition, offset)
	}
	for _, str := range Desserts {
		msg := &sarama.ProducerMessage{
			Topic: "Order",
			Value: sarama.StringEncoder(str + " customer_uuid:" + requset.UserUuid + " created_at:" + currentTime.String()),
		}
		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Сообщение отправлено в partiton", partition, offset)
	}
	err = client.Close()
	if err != nil {
		fmt.Println("Ошибка при закрытии client:", err)
	}
	err = producer.Close()
	if err != nil {
		fmt.Println("Ошибка при закрытии producer:", err)
	}
	return &customer.CreateOrderResponse{}, nil
}

/*func ProductList(products []*restaurant.Product) []*restaurant.Product {

}*/

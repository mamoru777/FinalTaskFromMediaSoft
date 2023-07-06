package service

import (
	"FinalTaskFromMediaSoft/Customer/internal/database"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/customer"
	"golang.org/x/net/context"
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

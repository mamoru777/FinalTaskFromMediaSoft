package main

import (
	"fmt"
	customer2 "gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/customer"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:13998", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("error connect to grpc server err:", err)
	}
	customer := customer2.NewUserServiceClient(conn)
	office := customer2.NewOfficeServiceClient(conn)
	//order := customer2.NewOrderServiceClient(conn)
Loop:
	for {
		fmt.Println("Приветствую, это программа для заказа обеда из ресторана в офисы, вам выведен список доступных команд, введите цифру без пробелов и других символов для команды, которую хотите выполнить:\n", "1: Создать офис\n", "2: Посмотреть список офисов\n", "3: Создать пользователя\n", "4: Посмотреть список пользователей по офисам\n", "5: Посмотреть актуальное меню\n", "6: Создать заказ\n", "7: Выход")
		var usrcase string
		fmt.Scanln(&usrcase)
		switch usrcase {
		case "1":
			var name string
			var address string
			fmt.Println("Введите имя офиса: ")
			fmt.Scanln(&name)
			fmt.Println("Введите адрес офиса: ")
			fmt.Scanln(&address)
			err := CreateOffice(office, customer2.CreateOfficeRequest{
				Name:    name,
				Address: address,
			})
			if err != nil {
				log.Fatal("Не удалось создать офис")
			}
		case "2":
			fmt.Println("Список существующих офисов:")
			names, err := GetOfficeList(office, customer2.GetOfficeListRequest{})
			if err != nil {
				log.Fatal("Failed to get office list: %v", err)
			}
			for _, n := range names {
				fmt.Println(n)
			}
			fmt.Println("\n")
		case "3":
			var name string
			var officeName string
			var officeId string
			var isExit bool = false
			fmt.Println("Введите имя пользователя:")
			fmt.Scanln(&name)
			fmt.Println("Перед вами список существующих офисов, впишите название того, к которому вы относитесь")
			names, err := GetOfficeList(office, customer2.GetOfficeListRequest{})
			if err != nil {
				log.Fatal("Failed to get office list: %v", err)
			}
			for _, n := range names {
				fmt.Println(n)
			}
			for isExit == false {
				fmt.Scanln(&officeName)
				offices, err2 := GetOfficeListModels(office, customer2.GetOfficeListRequest{})
				if err2 != nil {
					log.Fatal("Не удалось получить список офисов", err2)
				}
				fmt.Println("Введенное имя:", officeName)
				fmt.Println(len(offices))
				for _, o := range offices {
					fmt.Println("Имя из бд:", o.Name)

					if officeName == o.Name {
						officeId = o.Uuid
					}
				}
				if len(officeId) == 0 {
					fmt.Println("Вы неправильно ввели имя офиса, пожалуйста, попробуйте еще раз")
				} else {
					isExit = true
				}
			}
			fmt.Println(officeId)
			err3 := CreateCustomer(customer, customer2.CreateUserRequest{
				Name:       name,
				OfficeUuid: officeId,
			})
			if err3 != nil {
				log.Fatal("Не удалось создать пользователя")
			} else {
				fmt.Println("Пользователь создан успешно")
			}
		case "4":
			var officeId string
			var officeName string
			var isExit bool = false
			fmt.Println("Перед вами список существующих офисов, впишите название того, работников которого вы хотите посмотреть")
			names, err := GetOfficeList(office, customer2.GetOfficeListRequest{})
			if err != nil {
				log.Fatal("Failed to get office list: %v", err)
			}
			for _, n := range names {
				fmt.Println(n)
			}
			for isExit == false {
				fmt.Scanln(&officeName)
				offices, err2 := GetOfficeListModels(office, customer2.GetOfficeListRequest{})
				if err2 != nil {
					log.Fatal("Не удалось получить список офисов", err2)
				}
				for _, o := range offices {
					if officeName == o.Name {
						officeId = o.Uuid
					}
				}
				if len(officeId) == 0 {
					fmt.Println("Вы неправильно ввели имя офиса, пожалуйста, попробуйте еще раз")
				} else {
					isExit = true
				}
			}
			customers, err := GetCustomerList(customer, customer2.GetUserListRequest{OfficeUuid: officeId})
			if err != nil {
				log.Fatal("Не получилось загрузить пользователей", err)
			} else {
				fmt.Println("Имена сотрудников", officeName, ":")
			}
			for _, c := range customers {
				fmt.Println(c.Name)
			}
		case "7":
			break Loop
		default:
			fmt.Println("Вы неправильно ввели цифру, пожалуйста выберите нужный вам пункт и введите цифру без пробелов и других знаков")
		}
	}
}

func CreateOffice(office customer2.OfficeServiceClient, model customer2.CreateOfficeRequest) error {
	_, err := office.CreateOffice(context.Background(), &model)
	if err != nil {
		return err
	}
	fmt.Println("Офис успешно создан: ", model)
	return nil
}

func GetOfficeList(office customer2.OfficeServiceClient, req customer2.GetOfficeListRequest) ([]string, error) {
	var names []string
	res, _ := office.GetOfficeList(context.Background(), &req)
	for _, p := range res.Result {
		names = append(names, p.Name)
	}
	return names, nil
}

func GetOfficeListModels(office customer2.OfficeServiceClient, req customer2.GetOfficeListRequest) ([]*customer2.Office, error) {
	offices := []*customer2.Office{}
	res, err := office.GetOfficeList(context.Background(), &req)
	if err != nil {
		return nil, err
	}
	offices = res.Result
	return offices, nil
}

func CreateCustomer(customer customer2.UserServiceClient, model customer2.CreateUserRequest) error {
	_, err := customer.CreateUser(context.Background(), &model)
	if err != nil {
		return err
	}
	fmt.Println("Пользователь успешно создан: ", model)
	return nil
}

func GetCustomerList(customer customer2.UserServiceClient, req customer2.GetUserListRequest) ([]*customer2.User, error) {
	customers := []*customer2.User{}
	res, err := customer.GetUserList(context.Background(), &req)
	if err != nil {
		return nil, err
	}
	customers = res.Result
	return customers, nil
}

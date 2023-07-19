package main

import (
	"fmt"
	customer2 "gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/customer"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:13998", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("error connect to grpc server err:", err)
	}
	customer := customer2.NewUserServiceClient(conn)
	office := customer2.NewOfficeServiceClient(conn)
	order := customer2.NewOrderServiceClient(conn)
Loop:
	for {
		fmt.Print("Приветствую, это программа для заказа обеда из ресторана в офисы, вам выведен список доступных команд, введите цифру без пробелов и других символов для команды, которую хотите выполнить:\n", "1: Создать офис\n", "2: Посмотреть список офисов\n", "3: Создать пользователя\n", "4: Посмотреть список пользователей по офисам\n", "5: Посмотреть актуальное меню\n", "6: Создать заказ\n", "7: Выход\n")
		var usrcase string
		fmt.Scan(&usrcase)
		switch usrcase {
		case "1":
			var name string
			var address string
			fmt.Print("Введите имя офиса: \n")
			fmt.Scan(&name)
			fmt.Print("Введите адрес офиса: \n")
			fmt.Scan(&address)
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
			fmt.Print("Введите имя пользователя:\n")
			fmt.Scan(&name)
			fmt.Print("Перед вами список существующих офисов, впишите название того, к которому вы относитесь\n")
			names, err := GetOfficeList(office, customer2.GetOfficeListRequest{})
			if err != nil {
				log.Fatal("Failed to get office list: %v", err)
			}
			for _, n := range names {
				fmt.Println(n)
			}
			for isExit == false {
				fmt.Scan(&officeName)
				offices, err2 := GetOfficeListModels(office, customer2.GetOfficeListRequest{})
				if err2 != nil {
					log.Fatal("Не удалось получить список офисов", err2)
				}
				fmt.Print("Введенное имя:", officeName, "\n")
				fmt.Print(len(offices), "\n")
				for _, o := range offices {
					fmt.Print("Имя из бд:", o.Name, "\n")

					if officeName == o.Name {
						officeId = o.Uuid
					}
				}
				if len(officeId) == 0 {
					fmt.Print("Вы неправильно ввели имя офиса, пожалуйста, попробуйте еще раз\n")
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
				fmt.Scan(&officeName)
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
					fmt.Print("Вы неправильно ввели имя офиса, пожалуйста, попробуйте еще раз\n")
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
		case "5":
			var Salads []*customer2.Product
			var Garnishes []*customer2.Product
			var Meats []*customer2.Product
			var Soups []*customer2.Product
			var Drinks []*customer2.Product
			var Desserts []*customer2.Product
			currentTime := time.Now()
			nextDay := currentTime.AddDate(0, 0, 1)
			fmt.Println("Вот актуальное меню на ", nextDay, " :")
			actualmenu, err := GetActualMenu(order, customer2.GetActualMenuRequest{})
			if err != nil {
				log.Fatal("Не удалось загрузить меню", err)
			}
			for _, p := range actualmenu {
				if p.Type == 1 {
					Salads = append(Salads, p)
				}
				if p.Type == 2 {
					Garnishes = append(Garnishes, p)
				}
				if p.Type == 3 {
					Meats = append(Meats, p)
				}
				if p.Type == 4 {
					Soups = append(Soups, p)
				}
				if p.Type == 5 {
					Drinks = append(Drinks, p)
				}
				if p.Type == 6 {
					Desserts = append(Desserts, p)
				}
			}
			fmt.Println("Салаты:")
			for _, s := range Salads {
				fmt.Println(s.Name, s.Description, s.Weight, s.Price)
			}
			fmt.Println("Гарниры:")
			for _, g := range Garnishes {
				fmt.Println(g.Name, g.Description, g.Weight, g.Price)
			}
			fmt.Println("Мясо:")
			for _, m := range Meats {
				fmt.Println(m.Name, m.Description, m.Weight, m.Price)
			}
			fmt.Println("Супы:")
			for _, sp := range Soups {
				fmt.Println(sp.Name, sp.Description, sp.Weight, sp.Price)
			}
			fmt.Println("Напитки:")
			for _, dr := range Drinks {
				fmt.Println(dr.Name, dr.Description, dr.Weight, dr.Price)
			}
			fmt.Println("Дессерты:")
			for _, ds := range Desserts {
				fmt.Println(ds.Name, ds.Description, ds.Weight, ds.Price)
			}
		case "6":
			var Exit bool = false
			var productname string
			var count int32
			var Salads []*customer2.OrderItem
			var Garnishes []*customer2.OrderItem
			var Meats []*customer2.OrderItem
			var Soups []*customer2.OrderItem
			var Drinks []*customer2.OrderItem
			var Desserts []*customer2.OrderItem
			var Name string
			var CustomerId string
			var officeName string
			var officeId string

			fmt.Print("Перед вами список существующих офисов, впишите название того, работников которого вы хотите посмотреть\n")
			names, err := GetOfficeList(office, customer2.GetOfficeListRequest{})
			if err != nil {
				log.Fatal("Failed to get office list: %v", err)
			}
			for _, n := range names {
				fmt.Print(n + "\n")
			}
			for Exit == false {
				fmt.Scan(&officeName)
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
					fmt.Print("Вы неправильно ввели имя офиса, пожалуйста, попробуйте еще раз\n")
				} else {
					Exit = true
				}
			}
			Exit = false

			customers, err := GetCustomerList(customer, customer2.GetUserListRequest{OfficeUuid: officeId})
			if err != nil {
				log.Fatal("Не получилось загрузить пользователей", err)
			} else {
				fmt.Print("Имена сотрудников", officeName, ":", "\n")
			}
			for _, c := range customers {
				fmt.Println(c.Name)
			}

			fmt.Print("Введите имя пользователя, от лица которого хотите сделать заказ\n")
			fmt.Scan(&Name)
			for _, c := range customers {
				if Name == c.Name {
					CustomerId = c.Uuid
				}
			}
			if CustomerId == "" {
				CustomerId = "Unknown"
			}
			actualmenu, err := GetActualMenu(order, customer2.GetActualMenuRequest{})
			if err != nil {
				log.Fatal("Не удалось загрузить меню", err)
			}
			fmt.Print("Введите название салата и его количество и нажмите кнопку 'Enter' (он должны быть в списке продуктов), который вы хотите добавить в меню, после можете вписать название еще одного салата\n", "Напишите 'выход' без ковычек, если добавили нужные салаты\n")

			for Exit != true {
				fmt.Print(" Введите название салата\n")
				fmt.Scan(&productname)
				switch productname {
				case "выход":
					Exit = true
				default:
					fmt.Print(" Введите количество\n")
					fmt.Scan(&count)
					for _, p := range actualmenu {
						if p.Type == 1 && p.Name == productname {
							item := customer2.OrderItem{
								Count:       count,
								ProductUuid: p.Uuid,
							}
							Salads = append(Salads, &item)
						}
					}
				}
			}
			Exit = false

			fmt.Print("Введите название гарнира и его количество и нажмите кнопку 'Enter' (он должны быть в списке продуктов), который вы хотите добавить в меню, после можете вписать название еще одного гарнира\n", "Напишите 'выход' без ковычек, если добавили нужные гарниры\n")
			for Exit != true {
				fmt.Print(" Введите название гарнира\n")
				fmt.Scan(&productname)

				switch productname {
				case "выход":
					Exit = true
				default:
					fmt.Print(" Введите количество\n")
					fmt.Scan(&count)
					for _, p := range actualmenu {
						if p.Type == 2 && p.Name == productname {
							item := customer2.OrderItem{
								Count:       count,
								ProductUuid: p.Uuid,
							}
							Garnishes = append(Garnishes, &item)
						}
					}
				}
			}
			Exit = false

			fmt.Print("Введите название мяса и его количество нажмите кнопку 'Enter' (оно должны быть в списке продуктов), которое вы хотите добавить в меню, после можете вписать название еще одного мяса\n", "Напишите 'выход' без ковычек, если добавили нужное мясо\n")
			for Exit != true {
				fmt.Print(" Введите название мяса\n")
				fmt.Scan(&productname)

				switch productname {
				case "выход":
					Exit = true
				default:
					fmt.Print(" Введите количество\n")
					fmt.Scan(&count)
					for _, p := range actualmenu {
						if p.Type == 3 && p.Name == productname {
							item := customer2.OrderItem{
								Count:       count,
								ProductUuid: p.Uuid,
							}
							Meats = append(Meats, &item)
						}
					}
				}
			}
			Exit = false

			fmt.Print("Введите название супа и его количество нажмите кнопку 'Enter' (он должны быть в списке продуктов), который вы хотите добавить в меню, после можете вписать название еще одного супа\n", "Напишите 'выход' без ковычек, если добавили нужные супы\n")
			for Exit != true {
				fmt.Print(" Введите название супа\n")
				fmt.Scan(&productname)

				switch productname {
				case "выход":
					Exit = true
				default:
					fmt.Print(" Введите количество\n")
					fmt.Scan(&count)
					for _, p := range actualmenu {
						if p.Type == 4 && p.Name == productname {
							item := customer2.OrderItem{
								Count:       count,
								ProductUuid: p.Uuid,
							}
							Soups = append(Soups, &item)
						}
					}
				}
			}
			Exit = false

			fmt.Print("Введите название напитка и его количество и нажмите кнопку 'Enter' (он должны быть в списке продуктов), который вы хотите добавить в меню, после можете вписать название еще одного напитка\n", "Напишите 'выход' без ковычек, если добавили нужные напитки\n")
			for Exit != true {
				fmt.Print(" Введите название напитка\n")
				fmt.Scan(&productname)

				switch productname {
				case "выход":
					Exit = true
				default:
					fmt.Print(" Введите количество\n")
					fmt.Scan(&count)
					for _, p := range actualmenu {
						if p.Type == 5 && p.Name == productname {
							item := customer2.OrderItem{
								Count:       count,
								ProductUuid: p.Uuid,
							}
							Drinks = append(Drinks, &item)
						}
					}
				}
			}
			Exit = false
			fmt.Print("Введите название дессерта и его количество и нажмите кнопку 'Enter' (он должны быть в списке продуктов), который вы хотите добавить в меню, после можете вписать название еще одного дессерта\n", "Напишите 'выход' без ковычек, если добавили нужные дессерты\n")
			for Exit != true {
				fmt.Print(" Введите название дессерта\n")
				fmt.Scan(&productname)
				switch productname {
				case "выход":
					Exit = true
				default:
					fmt.Print(" Введите количество\n")
					fmt.Scan(&count)
					for _, p := range actualmenu {
						if p.Type == 6 && p.Name == productname {
							item := customer2.OrderItem{
								Count:       count,
								ProductUuid: p.Uuid,
							}
							Desserts = append(Desserts, &item)
						}
					}
				}
			}
			Exit = false
			err2 := CreateOrder(order, customer2.CreateOrderRequest{
				UserUuid:  CustomerId,
				Salads:    Salads,
				Garnishes: Garnishes,
				Meats:     Meats,
				Soups:     Soups,
				Drinks:    Drinks,
				Desserts:  Desserts,
			})
			if err2 != nil {
				log.Fatal("Не удалось создать заказ", err2)
			}
		case "7":
			conn.Close()
			break Loop
		default:
			fmt.Print("Вы неправильно ввели цифру, пожалуйста выберите нужный вам пункт и введите цифру без пробелов и других знаков\n")
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

func GetActualMenu(order customer2.OrderServiceClient, req customer2.GetActualMenuRequest) ([]*customer2.Product, error) {
	menu := []*customer2.Product{}
	res, err := order.GetActualMenu(context.Background(), &req)
	if err != nil {
		return nil, err
	}
	menu = append(menu, res.Salads...)
	menu = append(menu, res.Garnishes...)
	menu = append(menu, res.Meats...)
	menu = append(menu, res.Soups...)
	menu = append(menu, res.Drinks...)
	menu = append(menu, res.Desserts...)
	return menu, nil
}

func CreateOrder(order customer2.OrderServiceClient, model customer2.CreateOrderRequest) error {
	_, err := order.CreateOrder(context.Background(), &model)
	if err != nil {
		return err
	}
	fmt.Println("Заказ успешно создан")
	return nil
}

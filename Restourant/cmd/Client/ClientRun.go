package main

import (
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/restaurant"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:13999", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("error connect to grpc server err:", err)
	}
	product := restaurant.NewProductServiceClient(conn)
	menu := restaurant.NewMenuServiceClient(conn)
	//client := userapi.NewUserServiceClient(conn)
Loop:
	for {
		fmt.Println("Приветствую, это программа ресторана, вам выведен список доступных команд, введите цифру без пробелов и других символов для команды, которую хотите выполнить:\n", "1: Создать продукт\n", "2: Посмотреть список продуктов\n", "3: Создать меню\n", "4: Посмотреть меню\n", "5: Посмотреть список заказов\n", "6: Выход")
		var usrcase string
		fmt.Scanln(&usrcase)
		switch usrcase {
		case "1":
			var name string
			var description string
			var producttype restaurant.ProductType
			var producttypecase string
			var weight int32
			var price float64
			for {
				fmt.Println("Введите название продукта")
				_, err := fmt.Scanln(&name)
				if err != nil {
					fmt.Println("Введите название продукта в текстовом формате")
				} else {
					break
				}
			}
			for {
				fmt.Println("Введите описание продукта")
				_, err := fmt.Scanln(&description)
				if err != nil {
					fmt.Println("Введите описание продукта в текстовом формате")
				} else {
					break
				}
			}

			fmt.Println("Есть следующий список типов продуктов, выберите один из них, введя нужную цифру:\n", "(Если введете неправильно, то тип автоматически станет UNSPECIFIED)\n", "0: Unspecified\n", "1: Salad\n", "2: Garnish\n", "3: Meat\n", "4: Soup\n", "5: Drink\n", "6: Dessert")
			fmt.Scanln(&producttypecase)
			switch producttypecase {
			case "0":
				producttype = 0
			case "1":
				producttype = 1
			case "2":
				producttype = 2
			case "3":
				producttype = 3
			case "4":
				producttype = 4
			case "5":
				producttype = 5
			case "6":
				producttype = 6
			default:
				producttype = 0
			}
			for {
				fmt.Println("Введите вес продукта")
				_, err := fmt.Scanln(&weight)
				if err != nil {
					fmt.Println("Введите вес продукта в целочисленном числовом формате")
				} else {
					break
				}
			}
			for {
				fmt.Println("Введите цену продукта")
				_, err := fmt.Scanln(&price)
				if err != nil {
					fmt.Println("Введите цену продукта в числовом формате")
				} else {
					break
				}
			}

			err := CreateProduct(product, restaurant.CreateProductRequest{
				Name:        name,
				Description: description,
				Type:        producttype,
				Weight:      weight,
				Price:       price,
			})
			if err != nil {
				log.Fatal(err)
			} else {
				fmt.Println("Продукт создан успешно")
			}
		case "2":
			fmt.Println("Список существующих продуктов:\n")
			names, err := GetProductList(product, restaurant.GetProductListRequest{})
			if err != nil {
				log.Fatal("Failed to get product list: %v", err)
			}
			for _, n := range names {
				fmt.Println(n)
			}
			fmt.Println("\n")
		case "3":
			currentTime := time.Now()
			year, month, day := currentTime.Date()

			nextDay := currentTime.AddDate(0, 0, 1)
			nextDayProto, error := ptypes.TimestampProto(nextDay)
			if error != nil {
				log.Fatal(error)
			}
			var productname string
			var salats []string
			var garnishes []string
			var meats []string
			var soups []string
			var drinks []string
			var desserts []string
			var Exit bool = false
			var hoursOp int
			var minutesOp int
			var hoursCl int
			var minutesCl int
			fmt.Println("Введите название салата и нажмите кнопку 'Enter' (он должны быть в списке продуктов), который вы хотите добавить в меню, после можете вписать название еще одного салата\n", "Напишите 'выход' без ковычек, если добавили нужные салаты\n")

			for Exit != true {
				fmt.Println(" Введите название салата")
				fmt.Scanln(&productname)
				switch productname {
				case "выход":
					Exit = true
				default:
					salats = append(salats, productname)
				}
			}
			Exit = false

			fmt.Println("Введите название гарнира и нажмите кнопку 'Enter' (он должны быть в списке продуктов), который вы хотите добавить в меню, после можете вписать название еще одного гарнира\n", "Напишите 'выход' без ковычек, если добавили нужные гарниры\n")
			for Exit != true {
				fmt.Println(" Введите название гарнира")
				fmt.Scanln(&productname)
				switch productname {
				case "выход":
					Exit = true
				default:
					garnishes = append(garnishes, productname)
				}
			}
			Exit = false

			fmt.Println("Введите название мяса и нажмите кнопку 'Enter' (оно должны быть в списке продуктов), которое вы хотите добавить в меню, после можете вписать название еще одного мяса\n", "Напишите 'выход' без ковычек, если добавили нужное мясо\n")
			for Exit != true {
				fmt.Println(" Введите название мяса")
				fmt.Scanln(&productname)
				switch productname {
				case "выход":
					Exit = true
				default:
					meats = append(meats, productname)
				}
			}
			Exit = false

			fmt.Println("Введите название супа и нажмите кнопку 'Enter' (он должны быть в списке продуктов), который вы хотите добавить в меню, после можете вписать название еще одного супа\n", "Напишите 'выход' без ковычек, если добавили нужные супы\n")
			for Exit != true {
				fmt.Println(" Введите название супа")
				fmt.Scanln(&productname)
				switch productname {
				case "выход":
					Exit = true
				default:
					soups = append(soups, productname)
				}
			}
			Exit = false

			fmt.Println("Введите название напитка и нажмите кнопку 'Enter' (он должны быть в списке продуктов), который вы хотите добавить в меню, после можете вписать название еще одного напитка\n", "Напишите 'выход' без ковычек, если добавили нужные напитки\n")
			for Exit != true {
				fmt.Println(" Введите название напитка")
				fmt.Scanln(&productname)
				switch productname {
				case "выход":
					Exit = true
				default:
					drinks = append(drinks, productname)
				}
			}
			Exit = false

			fmt.Println("Введите название дессерта и нажмите кнопку 'Enter' (он должны быть в списке продуктов), который вы хотите добавить в меню, после можете вписать название еще одного дессерта\n", "Напишите 'выход' без ковычек, если добавили нужные дессерты\n")
			for Exit != true {
				fmt.Println(" Введите название дессерта")
				fmt.Scanln(&productname)
				switch productname {
				case "выход":
					Exit = true
				default:
					desserts = append(desserts, productname)
				}
			}
			Exit = false

			for {
				fmt.Println("Введите часы открытия приема заказов")
				_, err := fmt.Scanln(&hoursOp)
				if err != nil {
					fmt.Println("Введите часы в числовом формате")
				} else {
					break
				}
			}

			for {
				fmt.Println("Введите минуты открытия приема заказов")
				_, err := fmt.Scanln(&minutesOp)
				if err != nil {
					fmt.Println("Введите минуты в числовом формате")
				} else {
					break
				}
			}

			for {
				fmt.Println("Введите часы закрытия приема заказов")
				_, err := fmt.Scanln(&hoursCl)
				if err != nil {
					fmt.Println("Введите часы в числовом формате")
				} else {
					break
				}
			}

			for {
				fmt.Println("Введите минуты закрытия приема заказов")
				_, err := fmt.Scanln(&minutesCl)
				if err != nil {
					fmt.Println("Введите минуты в числовом формате")
				} else {
					break
				}
			}

			dateOpen := time.Date(year, month, day, hoursOp, minutesOp, 0, 0, currentTime.Location())
			dateOpenProto, error := ptypes.TimestampProto(dateOpen)
			if error != nil {
				log.Fatal(error)
			}

			dateClose := time.Date(year, month, day, hoursCl, minutesCl, 0, 0, currentTime.Location())
			dateCloseProto, error := ptypes.TimestampProto(dateClose)
			if error != nil {
				log.Fatal(error)
			}
			err := CreateMenu(menu, restaurant.CreateMenuRequest{
				OnDate:          nextDayProto,
				OpeningRecordAt: dateOpenProto,
				ClosingRecordAt: dateCloseProto,
				Salads:          salats,
				Garnishes:       garnishes,
				Meats:           meats,
				Soups:           soups,
				Drinks:          drinks,
				Desserts:        desserts,
			})
			if err != nil {
				log.Fatal(err)
			} else {
				fmt.Println("Меню создано успешно")
			}
		case "4":
			currentTime := time.Now()
			nextDay := currentTime.AddDate(0, 0, 1)
			nextDayYMD := nextDay.Format("2006-01-02")
			fmt.Println("Меню на ", nextDayYMD, ":\n")
			model, err := GetMenu(menu, restaurant.GetMenuRequest{})
			if err != nil {
				log.Fatal("Faild to load Menu", err)
			}
			fmt.Println("Id меню:", model.Uuid, "\n")
			fmt.Println("время открытия записи:", model.OpeningRecordAt.AsTime().Format("15:01"), "\n", "время закрытия записи:", model.ClosingRecordAt.AsTime().Format("15:01"), "\n")
			fmt.Println("Салаты:\n")
			for _, s := range model.Salads {
				fmt.Println(s.Name, "", s.Description, "", s.Weight, "", s.Price, "\n")
			}
			fmt.Println("Гарниры:\n")
			for _, g := range model.Garnishes {
				fmt.Println(g.Name, "", g.Description, "", g.Weight, "", g.Price, "\n")
			}
			fmt.Println("Мясо:\n")
			for _, m := range model.Meats {
				fmt.Println(m.Name, "", m.Description, "", m.Weight, "", m.Price, "\n")
			}
			fmt.Println("Супы:\n")
			for _, sp := range model.Soups {
				fmt.Println(sp.Name, "", sp.Description, "", sp.Weight, "", sp.Price, "\n")
			}
			fmt.Println("Напитки:\n")
			for _, dr := range model.Drinks {
				fmt.Println(dr.Name, "", dr.Description, "", dr.Weight, "", dr.Price, "\n")
			}
			fmt.Println("Дессерты:\n")
			for _, ds := range model.Desserts {
				fmt.Println(ds.Name, "", ds.Description, "", ds.Weight, "", ds.Price, "\n")
			}
		case "5":
		case "6":
			conn.Close()
			break Loop
		default:
			fmt.Println("Вы неправильно ввели цифру, пожалуйста выберите нужный вам пункт и введите цифру без пробелов и других знаков")
		}

	}

}

func CreateProduct(product restaurant.ProductServiceClient, model restaurant.CreateProductRequest) error {
	if _, err := product.CreateProduct(context.Background(), &model); err != nil {
		return err
	}
	log.Println("Product created: ", model)
	return nil
}

func GetProductList(product restaurant.ProductServiceClient, req restaurant.GetProductListRequest) ([]string, error) {
	/*if err, res  := product.GetProductList(context.Background(), &req); err != nil {
		log.Fatal("Failed to get product list: %v", err)
	}*/
	var names []string
	res, _ := product.GetProductList(context.Background(), &req)
	for _, p := range res.Result {
		names = append(names, p.Name)
	}
	return names, nil
}

func CreateMenu(menu restaurant.MenuServiceClient, model restaurant.CreateMenuRequest) error {
	if _, err := menu.CreateMenu(context.Background(), &model); err != nil {
		return err
	}
	log.Println("Menu created: ", model)
	return nil
}

func GetMenu(menu restaurant.MenuServiceClient, req restaurant.GetMenuRequest) (*restaurant.Menu, error) {
	res, _ := menu.GetMenu(context.Background(), &req)
	menuRes := res.Menu
	return menuRes, nil
}

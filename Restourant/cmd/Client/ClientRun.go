package main

import (
	"fmt"
	"github.com/golang/protobuf/ptypes/timestamp"
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/restaurant"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:13999", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("error connect to grpc server err:", err)
	}
	product := restaurant.NewProductServiceClient(conn)
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
			//var inputSucess bool = false
			var onDateString string
			var onDate timestamp.Timestamp
			fmt.Println("Введите на какую дату вы создаете меню ")
			_, err := fmt.Scanln(&onDateString)
			if err != nil {
				fmt.Println("Введите дату создания меню в нужном формате: ")
			} else {
				fmt.Println(onDate)
			}
			//onDate, err = time.Parse("2006-01-02", onDateString)
			/*for !inputSucess {
				fmt.Println("Введите на какую дату вы создаете меню ")
				_, err := fmt.Scanln(&onDate)
				if err != nil {
					fmt.Println("Введите дату создания меню в нужном формате: ")
				} else {
					inputSucess = true
				}

			}*/
		case "4":
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

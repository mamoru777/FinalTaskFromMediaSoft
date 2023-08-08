package main

import (
	"fmt"
	"github.com/golang/protobuf/ptypes"
	statistics2 "gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/statistics"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:13997", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("error connect to grpc server err:", err)
	}
	statistics := statistics2.NewStatisticsServiceClient(conn)
Loop:
	for {
		fmt.Print("Приветствую, это программа статистики работы ресторана, вам выведен список доступных команд, введите цифру без пробелов и других символов для команды, которую хотите выполнить:\n", "1: Посмотреть выручку\n", "2: Посмотреть топ продуктов\n", "3: Выход\n")
		var usrcase string
		fmt.Scan(&usrcase)
		switch usrcase {
		case "1":
			var dateStart string
			var dateEnd string
			var date1 time.Time
			var date2 time.Time
			for {
				fmt.Print("Введите начальную дату, с которой вы хотите посмотреть прибыль ресторана в формате YYYY-MM-DD\n")
				//_, err := fmt.Scan(&dateStart)
				fmt.Scan(&dateStart)
				date1, err = time.Parse("2006-01-02", dateStart)
				if err != nil {
					fmt.Print("Введите дату в правильном формате YYYY-MM-DD\n")
				} else {
					break
				}
			}
			for {
				fmt.Print("Введите конечную дату, с которой вы хотите посмотреть прибыль ресторана в формате YYYY-MM-DD\n")
				fmt.Scan(&dateEnd)
				date2, err = time.Parse("2006-01-02", dateEnd)
				if err != nil {
					fmt.Print("Введите дату в правильном формате YYYY-MM-DD\n")
				} else {
					break
				}
			}
			dateStartProto, error := ptypes.TimestampProto(date1)
			if error != nil {
				log.Print(error)
			}
			dateEndProto, error := ptypes.TimestampProto(date2)
			if error != nil {
				log.Print(error)
			}
			profit, err := GetAmountOfProfit(statistics, statistics2.GetAmountOfProfitRequest{
				StartDate: dateStartProto,
				EndDate:   dateEndProto,
			})
			if err != nil {
				log.Print("Не удалось получить информацию о выручке", err)
			}
			fmt.Print("Колличество выручки за период с ", dateStart, " по ", dateEnd, ": ", profit, "\n")
		case "2":
			var dateStart string
			var dateEnd string
			var date1 time.Time
			var date2 time.Time
			for {
				fmt.Print("Введите начальную дату, с которой вы хотите посмотреть топ продуктов ресторана в формате YYYY-MM-DD\n")
				//_, err := fmt.Scan(&dateStart)
				fmt.Scan(&dateStart)
				date1, err = time.Parse("2006-01-02", dateStart)
				if err != nil {
					fmt.Print("Введите дату в правильном формате YYYY-MM-DD\n")
				} else {
					break
				}
			}
			for {
				fmt.Print("Введите конечную дату, с которой вы хотите посмотреть топ продуктов ресторана в формате YYYY-MM-DD\n")
				fmt.Scan(&dateEnd)
				date2, err = time.Parse("2006-01-02", dateEnd)
				if err != nil {
					fmt.Print("Введите дату в правильном формате YYYY-MM-DD\n")
				} else {
					break
				}
			}
			dateStartProto, error := ptypes.TimestampProto(date1)
			if error != nil {
				log.Print(error)
			}
			dateEndProto, error := ptypes.TimestampProto(date2)
			if error != nil {
				log.Print(error)
			}
			var topProduct []*statistics2.Product
			fmt.Print("Перед вами список существующих видов продуктов, введите нужную цифру для просмотра топ продуктов по данному виду продуктов,\n или же нажмите 7, чтобы посмотреть топ всех продуктов, независимо от типа\n", "0: Неопределенный\n", "1: Салаты\n", "2: Гарниры\n", "3: Мясо\n", "4: Супы\n", "5: Напитки\n", "6: Дессерты\n", "7: Все продукты\n")
			var productTypeUsr string
			var productType statistics2.StatisticsProductType
			fmt.Scan(&productTypeUsr)
			switch productTypeUsr {
			case "0":
				productType = 0
				topProduct, err = TopProducts(statistics, statistics2.TopProductsRequest{
					ProductType: &productType,
					StartDate:   dateStartProto,
					EndDate:     dateEndProto,
				})
				if err != nil {
					log.Print("Не удалось получить информацию о топ продуктах")
				}
			case "1":
				productType = 1
				topProduct, err = TopProducts(statistics, statistics2.TopProductsRequest{
					ProductType: &productType,
					StartDate:   dateStartProto,
					EndDate:     dateEndProto,
				})
				if err != nil {
					log.Print("Не удалось получить информацию о топ продуктах")
				}
			case "2":
				productType = 2
				topProduct, err = TopProducts(statistics, statistics2.TopProductsRequest{
					ProductType: &productType,
					StartDate:   dateStartProto,
					EndDate:     dateEndProto,
				})
				if err != nil {
					log.Print("Не удалось получить информацию о топ продуктах")
				}
			case "3":
				productType = 3
				topProduct, err = TopProducts(statistics, statistics2.TopProductsRequest{
					ProductType: &productType,
					StartDate:   dateStartProto,
					EndDate:     dateEndProto,
				})
				if err != nil {
					log.Print("Не удалось получить информацию о топ продуктах")
				}
			case "4":
				productType = 4
				topProduct, err = TopProducts(statistics, statistics2.TopProductsRequest{
					ProductType: &productType,
					StartDate:   dateStartProto,
					EndDate:     dateEndProto,
				})
				if err != nil {
					log.Print("Не удалось получить информацию о топ продуктах")
				}
			case "5":
				productType = 5
				topProduct, err = TopProducts(statistics, statistics2.TopProductsRequest{
					ProductType: &productType,
					StartDate:   dateStartProto,
					EndDate:     dateEndProto,
				})
				if err != nil {
					log.Print("Не удалось получить информацию о топ продуктах")
				}
			case "6":
				productType = 6
				topProduct, err = TopProducts(statistics, statistics2.TopProductsRequest{
					ProductType: &productType,
					StartDate:   dateStartProto,
					EndDate:     dateEndProto,
				})
				if err != nil {
					log.Print("Не удалось получить информацию о топ продуктах")
				}
			case "7":
				topProduct, err = TopProducts(statistics, statistics2.TopProductsRequest{
					StartDate: dateStartProto,
					EndDate:   dateEndProto,
				})
				if err != nil {
					log.Print("Не удалось получить информацию о топ продуктах")
				}
			default:
				fmt.Print("Вы неправильно ввели цифру, пожалуйста выберите нужный вам пункт и введите цифру без пробелов и других знаков\n")
			}
			i := 1
			for _, p := range topProduct {
				fmt.Print(i, " Место: ", p.Name, "\n", p.Count, "\n", p.ProductType, "\n\n")
				i++
			}
		case "3":
			conn.Close()
			break Loop
		default:
			fmt.Print("Вы неправильно ввели цифру, пожалуйста выберите нужный вам пункт и введите цифру без пробелов и других знаков\n")
		}
	}

}

func GetAmountOfProfit(statistics statistics2.StatisticsServiceClient, request statistics2.GetAmountOfProfitRequest) (float64, error) {
	req, err := statistics.GetAmountOfProfit(context.Background(), &request)
	if err != nil {
		return 0, err
	}
	return req.Profit, nil
}

func TopProducts(statistics statistics2.StatisticsServiceClient, request statistics2.TopProductsRequest) ([]*statistics2.Product, error) {
	req, err := statistics.TopProducts(context.Background(), &request)
	if err != nil {
		return nil, err
	}
	return req.Result, nil
}

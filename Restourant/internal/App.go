package internal

import (
	"FinalTaskFromMediaSoft/Restourant/internal/configdb"
	"FinalTaskFromMediaSoft/Restourant/internal/configserv"
	"FinalTaskFromMediaSoft/Restourant/internal/database"
	"FinalTaskFromMediaSoft/Restourant/internal/service"
	"FinalTaskFromMediaSoft/pkg/contracts/restaurant"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/google/uuid"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"regexp"
	"strconv"
	"syscall"
	"time"
)

func Run(cfg configdb.ConfigDb, config configserv.ConfigServ) error {
	/*go func() {
		Loop
		for {

		}
	}()*/
	configkafka := sarama.NewConfig()
	configkafka.Consumer.Return.Errors = true

	client, err := sarama.NewClient([]string{"localhost:9092"}, configkafka)
	if err != nil {
		log.Fatal("Error creating client: ", err)
	}
	defer client.Close()

	consumer, err := sarama.NewConsumerFromClient(client)
	if err != nil {
		log.Fatal("Failed to create consumer", err)
	}
	topic := "Order"
	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatal("Error closing consumer: ", err)
		}
	}()

	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		log.Fatal("Failed to create partition consumer", err)
	}

	db, err := database.InitDb(cfg)
	if err != nil {
		log.Fatal("Cannot to init the database", err)
	}
	rep := database.New(db)
	serv := service.New(rep)
	s := grpc.NewServer()
	restaurant.RegisterProductServiceServer(s, serv)
	restaurant.RegisterMenuServiceServer(s, serv)
	restaurant.RegisterOrderServiceServer(s, serv)

	l, err := net.Listen("tcp", config.GRPCAddr)
	if err != nil {
		return fmt.Errorf("failed to listen tcp %s, %v", config.GRPCAddr, err)
	}

	go func() {
		log.Printf("starting listening grpc server at %s", "13999") //config.GRPCAddr)
		if err := s.Serve(l); err != nil {
			log.Fatalf("error service grpc server %v", err)
		}
	}()
	countRegex := regexp.MustCompile(`count:(\d+)`)
	productUUIDRegex := regexp.MustCompile(`product_uuid:"([^"]+)"`)
	customerUUIDRegex := regexp.MustCompile(`customer_uuid:([^" ]+)`)
	createdAtRegex := regexp.MustCompile(`created_at:([^" ]+)`)
	go func() {
		for {
			select {
			case msg := <-partitionConsumer.Messages():

				CountMatch := countRegex.FindStringSubmatch(string(msg.Value))
				ProductUuidMatch := productUUIDRegex.FindStringSubmatch(string(msg.Value))
				CustomerUuidMatch := customerUUIDRegex.FindStringSubmatch(string(msg.Value))
				strCreatedAtMatch := createdAtRegex.FindStringSubmatch(string(msg.Value))
				var strCount string
				var strProductUuid string
				var strCustomerUuid string
				var strCreatedAt string

				if len(CountMatch) > 0 {
					strCount = CountMatch[1]
					fmt.Print(strCount)
				} else {
					fmt.Print("не удалось найти значение count")
				}

				if len(ProductUuidMatch) > 0 {
					strProductUuid = ProductUuidMatch[1]
					fmt.Print(strProductUuid)
				} else {
					fmt.Print("Не удалось найти значение strProductUuid")
				}

				if len(CustomerUuidMatch) > 0 {
					strCustomerUuid = CustomerUuidMatch[1]
					fmt.Print(strCustomerUuid)
				} else {
					fmt.Print("не удалось найти значение strCustomerUuid")
				}

				if len(strCreatedAtMatch) > 0 {
					strCreatedAt = strCreatedAtMatch[1]
					fmt.Print(strCreatedAt)
				} else {
					fmt.Print("Не удалось найти значение strCreatedAt")
				}

				count, err := strconv.Atoi(strCount)
				if err != nil {
					log.Fatal("не удалось конвертировать count ", err)
				}
				count64 := int64(count)
				productUuid, err := uuid.Parse(strProductUuid)
				if err != nil {
					log.Fatal("не удалось конвертировать productUuid ", err)
				}
				customerUUID, err := uuid.Parse(strCustomerUuid)
				if err != nil {
					log.Fatal("не удалось конвертировать customerUUID ", err)
				}
				createdAt, err := time.Parse("2006-01-02", strCreatedAt)
				if err != nil {
					log.Fatal("не удалось конвертировать createdAt ", err)
				}
				order := database.Order{
					ProductId:   productUuid,
					Count:       count64,
					CreatedAt:   createdAt,
					CustomerId:  customerUUID,
					ProductName: rep.GetProductById(context.Background(), productUuid),
				}
				rep.CreateOrder(context.Background(), &order)
				fmt.Printf("Received message: Topic: %s, Partition: %d, Offset: %d, Key: %s, Value: %s\n",
					msg.Topic, msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
			case err := <-partitionConsumer.Errors():
				log.Println("Error:", err)
			}
		}
	}()

	gracefulShutDown(s)

	return nil
}

func gracefulShutDown(s *grpc.Server) {
	const waitTime = 5 * time.Second

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(ch)

	sig := <-ch
	errorMessage := fmt.Sprintf("%s %v - %s", "Received shutdown signal:", sig, "Graceful shutdown done")
	log.Println(errorMessage)
	s.GracefulStop()

}

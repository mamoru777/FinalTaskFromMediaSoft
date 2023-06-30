package internal

import (
	"FinalTaskFromMediaSoft/Restourant/internal/configdb"
	"FinalTaskFromMediaSoft/Restourant/internal/configserv"
	"FinalTaskFromMediaSoft/Restourant/internal/database"
	"FinalTaskFromMediaSoft/Restourant/internal/service"
	"fmt"
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/restaurant"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(cfg configdb.ConfigDb, config configserv.ConfigServ) error {
	db, err := database.InitDb(cfg)
	if err != nil {
		log.Fatal("Cannot to init the database", err)
	}
	serv := service.New(database.New(db))
	s := grpc.NewServer()
	restaurant.RegisterProductServiceServer(s, serv)

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

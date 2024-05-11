package app

import (
	"fmt"
	"log"
	"net"

	"github.com/RizqiSugiarto/GaleryService/config"
	handler "github.com/RizqiSugiarto/GaleryService/internal/handler/grpc"
	"github.com/RizqiSugiarto/GaleryService/internal/usecase"
	"github.com/RizqiSugiarto/GaleryService/internal/usecase/repository"
	"github.com/RizqiSugiarto/GaleryService/pkg/maria"
	"google.golang.org/grpc"
)

func Run(cfg *config.Config)  {
	ml, err := maria.New(cfg.UrlDb, maria.MaxOpenConns(10))

	if err != nil {
		log.Fatal(fmt.Errorf("app - run - maria.New"))
	}
	defer ml.Close()

	RunMigrate(ml.Db)

	lis, err := net.Listen("tcp", cfg.GrpcServer.Addr)

	if err != nil {
		log.Fatalf("app - run - net.listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	galleryService := usecase.New(repository.New(ml))
	handler.NewGrpcService(grpcServer, galleryService)

	log.Println("Starting gRPC server on", cfg.GrpcServer.Addr)

	if err :=  grpcServer.Serve(lis); err != nil {
		log.Fatalf("app - run - grpcServer.Serve")
	}
}
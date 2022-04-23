package main

import (
	"fmt"
	"log"
	"net"

	"github.com/ab3llo/go-order-svc/pkg/client"
	"github.com/ab3llo/go-order-svc/pkg/config"
	"github.com/ab3llo/go-order-svc/pkg/db"
	"github.com/ab3llo/go-order-svc/pkg/pb"
	"github.com/ab3llo/go-order-svc/pkg/services"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed to load config", err)
	}

	database := db.Connect(&cfg)

	lis, err := net.Listen("tcp", cfg.Port)

	if err != nil {
		log.Fatalln("Failed to listen to server:", err)
	}

	productSvc := client.NewProductServiceClient(cfg.ProductSvcUrl)
	s := services.Server{
		DbConnection: database,
		ProductSvc:   productSvc,
	}

	fmt.Println("Order svc listening on port: ", cfg.Port)

	grpcServer := grpc.NewServer()

	pb.RegisterOrderServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to start grpc Server", err)
	}
}

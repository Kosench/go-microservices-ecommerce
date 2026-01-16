package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	apiv1 "github.com/Kosench/go-microservices-ecommerce/inventory/internal/api/inventory/v1"
	"github.com/Kosench/go-microservices-ecommerce/inventory/internal/repository/part"
	invService "github.com/Kosench/go-microservices-ecommerce/inventory/internal/service/part"
	inventoryv1 "github.com/Kosench/go-microservices-ecommerce/shared/pkg/proto/inventory/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const grpcPort = 50051

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Printf("failed to listen: %v\n", err)
		return
	}
	defer func() {
		if cerr := lis.Close(); cerr != nil {
			log.Printf("failed to close listener :%v\n", cerr)
		}
	}()

	s := grpc.NewServer()

	repo := part.NewMemoryInventoryRepository()
	srv := invService.NewInventoryService(repo)
	api := apiv1.NewAPI(srv)

	inventoryv1.RegisterInventoryServiceServer(s, api)

	reflection.Register(s)

	go func() {
		log.Printf("gRPC server listening on %d\n", grpcPort)
		err = s.Serve(lis)
		if err != nil {
			log.Printf("failed to serve: %v\n", err)
			return
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("🛑 Shutting down gRPC server...")
	s.GracefulStop()
	log.Println("✅ Server stopped")
}

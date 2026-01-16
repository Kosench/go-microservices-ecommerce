package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	apiV1 "github.com/Kosench/go-microservices-ecommerce/payment/internal/api/payment/v1"
	"github.com/Kosench/go-microservices-ecommerce/payment/internal/service/payment"
	paymentV1 "github.com/Kosench/go-microservices-ecommerce/shared/pkg/proto/payment/v1"
	"google.golang.org/grpc"
)

const grpcPort = 50052

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

	srv := payment.NewPayService()
	api := apiV1.NewAPI(srv)

	paymentV1.RegisterPaymentServiceServer(s, api)

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

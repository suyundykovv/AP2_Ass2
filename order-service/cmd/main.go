package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"order-service/config"
	api "order-service/internal/api"
	"order-service/internal/nats"
	"order-service/internal/repository"
	"order-service/internal/service"

	pb "github.com/suyundykovv/protos/gen/go/order/v1"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func main() {
	// Load configuration
	// Initialize database connection
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()
	nc, err := nats.Connect("nats://localhost:4222", 5, 5)
	if err != nil {
		log.Fatalf("Failed to connect to NATS: %v", err)
	}
	defer nc.Close()
	natsPub := nats.NewPublisher(nc)

	// Initialize the repository, service, and gRPC server
	orderRepo := repository.NewSQLOrderRepository(db)
	orderService := service.NewOrderService(orderRepo, natsPub)
	orderServer := api.NewOrderServer(orderService)
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(api.LoggingInterceptor),
	)

	pb.RegisterOrderServiceServer(grpcServer, orderServer)

	// Start gRPC server
	lis, err := net.Listen("tcp", ":"+"8082")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Graceful shutdown
	go func() {
		log.Printf("Order service running on port %s", "8082")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	grpcServer.GracefulStop()
	log.Println("Server exited properly")
}

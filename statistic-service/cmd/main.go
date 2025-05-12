package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"statistic-service/config"
	"statistic-service/internal/api"
	"statistic-service/internal/nats"
	"statistic-service/internal/repository"
	"statistic-service/internal/service"
	"syscall"
	"time"

	"github.com/nats-io/nats.go"
	pb "github.com/suyundykovv/protos/gen/go/statistics/v1"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.Load()

	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize NATS connection
	nc, err := nats.Connect(cfg.NATS.URL)
	if err != nil {
		log.Fatalf("Failed to connect to NATS: %v", err)
	}
	defer nc.Close()

	// Initialize repository and service
	statsRepo := repository.NewSQLStatisticRepository(db)
	statsService := service.NewStatisticService(statsRepo)

	// Initialize NATS subscriber
	natsSubscriber := nats.NewSubscriber(statsService, nc)
	if err := natsSubscriber.Subscribe(); err != nil {
		log.Fatalf("Failed to subscribe to NATS: %v", err)
	}

	// Initialize gRPC server
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(api.LoggingInterceptor),
	)
	statsServer := api.NewStatisticsServer(statsService)
	pb.RegisterStatisticsServiceServer(grpcServer, statsServer)

	// Start gRPC server
	lis, err := net.Listen("tcp", ":"+cfg.GRPC.Port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	go func() {
		log.Printf("Statistics service running on port %s", cfg.GRPC.Port)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	grpcServer.GracefulStop()
	log.Println("Server exited properly")
}

package main

import (
	"context"
	"inventory-service/config"
	api "inventory-service/internal/api"
	"inventory-service/internal/nats"
	"inventory-service/internal/repository"
	"inventory-service/internal/service"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	pb "github.com/suyundykovv/protos/gen/go/inventory/v1"
	"google.golang.org/grpc"
)

func main() {
	// Initialize database connection
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()
	log.Println("Successfully connected to the database!")

	// Initialize NATS connection
	nc, err := nats.Connect("nats://nats:4222", 5, 5)
	if err != nil {
		log.Fatalf("Failed to connect to NATS: %v", err)
	}
	defer nc.Close()
	log.Println("Successfully connected to NATS!")

	// Initialize services
	natsPub := nats.NewPublisher(nc)
	inventoryRepo := repository.NewSQLProductRepository(db)
	productService := service.NewProductService(inventoryRepo, natsPub)

	// Initialize context for cache operations
	ctx := context.Background()

	// Type assertion to access cache methods

	// Warm up cache
	if err := productService.WarmCache(ctx); err != nil {
		log.Printf("Warning: Failed to warm cache: %v", err)
	}

	// Start cache refresh
	go productService.StartCacheRefresh(ctx)
	log.Println("Started cache refresh routine")

	// Initialize gRPC server
	productServer := api.NewInventoryServer(productService)
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(api.LoggingInterceptor),
	)
	pb.RegisterInventoryServiceServer(grpcServer, productServer)

	// Start gRPC server
	lis, err := net.Listen("tcp", ":"+"8081")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	go func() {
		log.Printf("Inventory service running on port %s", "8081")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	// Graceful shutdown handling
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
	grpcServer.GracefulStop()
	log.Println("Server exited properly")
}

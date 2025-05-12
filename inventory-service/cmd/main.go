package main

import (
	"context"
	"inventory-service/config"
	api "inventory-service/internal/api"
	"inventory-service/internal/repository"
	"inventory-service/internal/service"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	pb "github.com/suyundykovv/protos/gen/go/inventory/v1"

	"google.golang.org/grpc"
)

func main() {
	db, err := config.InitDB()
	if err == nil {
		log.Println("Successfully connected to the database!")
	}
	defer db.Close()
	inventoryRepo := repository.NewSQLProductRepository(db)
	productService := service.NewProductService(inventoryRepo)
	productServer := api.NewInventoryServer(productService)

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(api.LoggingInterceptor),
	)
	pb.RegisterInventoryServiceServer(grpcServer, productServer)

	lis, err := net.Listen("tcp", ":"+"8081")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	go func() {
		log.Printf("User service running on port %s", "8074")
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

package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"order-service/config"
	"order-service/internal/handler"
	"order-service/internal/repository"
	pb "order-service/proto/order"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func main() {
	// Initialize database connection
	var db *sql.DB
	var err error

	for i := 0; i < 5; i++ {
		db, err = config.InitDB()
		if err == nil {
			log.Println("Successfully connected to the database!")
			break
		}
		log.Printf("Error connecting to the database (attempt %d): %v\n", i+1, err)
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to connect to the database after 5 attempts: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	// Initialize repository
	orderRepo := repository.NewOrderRepository(db)

	// Create gRPC server
	grpcServer := grpc.NewServer()
	orderServer := handler.NewOrderServer(orderRepo)
	pb.RegisterOrderServiceServer(grpcServer, orderServer)

	// Get the service port from environment variables
	port := os.Getenv("ORDER_SERVICE_PORT")
	if port == "" {
		port = "50052" // Default gRPC port if not set
	}

	// Start listening
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("Order Service gRPC server running on port %s", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

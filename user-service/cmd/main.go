package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"user-service/config"
	api "user-service/internal/api"
	"user-service/internal/repository"
	"user-service/internal/service"

	pb "github.com/suyundykovv/protos/gen/go/user/v1"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func main() {
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	userRepo := repository.NewSQLOrderRepository(db)
	userService := service.NewOrderService(userRepo)
	userServer := api.NewUserServer(userService)

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(api.LoggingInterceptor),
	)

	pb.RegisterUserServiceServer(grpcServer, userServer)

	lis, err := net.Listen("tcp", ":"+"8082")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	go func() {
		log.Printf("User service running on port %s", "8084")
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

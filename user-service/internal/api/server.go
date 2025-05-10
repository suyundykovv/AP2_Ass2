package grpc

import (
	"context"
	"log"
	"time"

	"user-service/internal/service"

	pb "github.com/suyundykovv/protos/gen/go/user/v1"
	"google.golang.org/grpc"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
	service service.UserService
}

func NewUserServer(svc service.UserService) *UserServer {
	return &UserServer{service: svc}
}

func LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	start := time.Now()
	resp, err = handler(ctx, req)
	log.Printf("Method: %s, Duration: %v, Error: %v", info.FullMethod, time.Since(start), err)
	return resp, err
}

func (s *UserServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.User, error) {
	order, err := s.service.CreateUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (s *UserServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
	order, err := s.service.GetUser(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return order, nil
}
func (s *UserServer) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.User, error) {
	return s.service.UpdateUser(ctx, req.Id)
}

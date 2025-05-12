package client

import (
	"context"
	"log"
	"time"

	pb "github.com/suyundykovv/protos/gen/go/user/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserClient struct {
	client pb.UserServiceClient
}

func NewUserClient(address string) *UserClient {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect to order service: %v", err)
	}
	return &UserClient{
		client: pb.NewUserServiceClient(conn),
	}
}

func (oc *UserClient) CreateUser(request *pb.CreateUserRequest) (*pb.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	order, err := oc.client.CreateUser(ctx, request)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (oc *UserClient) GetUser(request *pb.GetUserRequest) (*pb.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user, err := oc.client.GetUser(ctx, request)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (oc *UserClient) UpdateUser(request *pb.UpdateUserRequest) (*pb.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	order, err := oc.client.UpdateUser(ctx, request)
	if err != nil {
		return nil, err
	}
	return order, nil
}

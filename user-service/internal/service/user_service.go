package service

import (
	"context"
	"fmt"
	"user-service/internal/handler"
	"user-service/internal/repository"

	pb "github.com/suyundykovv/protos/gen/go/user/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService interface {
	CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.User, error)
	GetUser(ctx context.Context, id string) (*pb.User, error)
	UpdateUser(ctx context.Context, id string) (*pb.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewOrderService(repo repository.UserRepository) *userService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.User, error) {
	user := handler.ProtoToDomainUser(req)
	if err := s.repo.Create(ctx, user); err != nil {
		return nil, err
	}
	return handler.DomainToProtoUser(user), nil
}

func (s *userService) GetUser(ctx context.Context, id string) (*pb.User, error) {
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, fmt.Errorf("user with ID %s not found", id)
	}
	return handler.DomainToProtoUser(user), nil
}
func (s *userService) UpdateUser(ctx context.Context, id string) (*pb.User, error) {
	err := s.repo.Update(ctx, id)
	if err != nil {
		if err.Error() == fmt.Sprintf("user with ID %s not found", id) {
			return nil, statusError(codes.NotFound, err.Error())
		}
		return nil, statusError(codes.Internal, "failed to update order status")
	}

	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, statusError(codes.Internal, "failed to fetch updated order")
	}
	if user == nil {
		return nil, statusError(codes.NotFound, fmt.Sprintf("order with ID %s not found", id))
	}

	return handler.DomainToProtoUser(user), nil
}

func statusError(code codes.Code, msg string) error {
	return status.Errorf(code, "%s", msg)
}

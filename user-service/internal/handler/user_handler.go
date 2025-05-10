package handler

import (
	"time"
	"user-service/internal/domain"

	pb "github.com/suyundykovv/protos/gen/go/user/v1"
)

// Преобразование из protobuf в domain
func ProtoToDomainUser(req *pb.CreateUserRequest) *domain.User {
	return &domain.User{
		Username:  req.GetUsername(),
		Email:     req.GetEmail(), // В нашем случае это просто список строк (ID товаров)
		Password:  req.GetPassword(),
		Role:      req.GetRole(),
		CreatedAt: time.Now().Unix(),
	}
}

// Обратное преобразование из domain в protobuf
func DomainToProtoUser(user *domain.User) *pb.User {
	return &pb.User{
		Id:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Password:  user.Password,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
	}
}

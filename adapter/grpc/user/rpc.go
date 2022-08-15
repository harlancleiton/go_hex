package rpcuser

import (
	"context"

	"github.com/harlancleiton/gopersonalities/adapter/grpc/pb"
	"github.com/harlancleiton/gopersonalities/core/domain/dto"
)

func (adapter GRPCUserAdapter) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	createUserDTO, err := dto.CreateUserDTOFromProto(req)

	if err != nil {
		return nil, err
	}

	user, err := adapter.userApi.CreateUser(createUserDTO)

	if err != nil {
		return nil, err
	}

	response := &pb.CreateUserResponse{
		Id:        user.Id,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return response, nil
}

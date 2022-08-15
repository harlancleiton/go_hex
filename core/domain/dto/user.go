package dto

import (
	"encoding/json"
	"io"
	"time"

	"github.com/harlancleiton/gopersonalities/adapter/grpc/pb"
	"github.com/harlancleiton/gopersonalities/core/domain/entity"
)

type UserDTO struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type CreateUserDTO struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func CreateUserDTOFromJSON(body io.Reader) (*CreateUserDTO, error) {
	dto := CreateUserDTO{}
	err := json.NewDecoder(body).Decode(&dto)

	if err != nil {
		return nil, err
	}

	return &dto, nil
}

func CreateUserDTOFromProto(createUser *pb.CreateUserRequest) (*CreateUserDTO, error) {
	dto := CreateUserDTO{
		FirstName: createUser.FirstName,
		LastName:  createUser.LastName,
		Email:     createUser.Email,
		Password:  createUser.Password,
	}

	return &dto, nil
}

func UserDTOFromUser(user *entity.User) *UserDTO {
	return &UserDTO{
		Id:        user.Id(),
		FirstName: user.FirstName(),
		LastName:  user.LastName(),
		Email:     user.Email(),
		CreatedAt: user.CreatedAt().Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt().Format(time.RFC3339),
	}
}

package usecase

import (
	"github.com/harlancleiton/gopersonalities/core/domain/dto"
	"github.com/harlancleiton/gopersonalities/core/domain/entity"
)

type hashPassword struct {
	createUser CreateUser
}

func NewHashPassword(createUser CreateUser) CreateUser {
	return &hashPassword{
		createUser: createUser,
	}
}

func (u *hashPassword) Execute(dto *dto.CreateUserDTO) (*entity.User, error) {
	dto.Password = ""

	return u.createUser.Execute(dto)
}

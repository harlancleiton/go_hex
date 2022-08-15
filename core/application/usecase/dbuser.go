package usecase

import (
	"time"

	"github.com/harlancleiton/gopersonalities/core/domain/dto"
	"github.com/harlancleiton/gopersonalities/core/domain/entity"
	"github.com/harlancleiton/gopersonalities/core/ports/right/repository"
)

type createUserUsecase struct {
	userRepository repository.UserRepository
}

func NewCreateUser(userRepository repository.UserRepository) CreateUser {
	return &createUserUsecase{
		userRepository: userRepository,
	}
}

func (u *createUserUsecase) Execute(dto *dto.CreateUserDTO) (*entity.User, error) {
	user := entity.NewUser(
		"",
		dto.FirstName,
		dto.LastName,
		dto.Email,
		dto.Password,
		time.Time{},
		time.Time{},
	)

	return u.userRepository.Create(user)
}

package usecase

import (
	"github.com/harlancleiton/gopersonalities/core/domain/dto"
	"github.com/harlancleiton/gopersonalities/core/domain/entity"
)

type CreateUser interface {
	Execute(*dto.CreateUserDTO) (*entity.User, error)
}

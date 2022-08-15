package repository

import (
	"github.com/harlancleiton/gopersonalities/core/domain/entity"
)

type UserRepository interface {
	Create(user *entity.User) (*entity.User, error)
}

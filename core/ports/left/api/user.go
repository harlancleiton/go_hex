package apiport

import (
	"github.com/harlancleiton/gopersonalities/core/domain/dto"
)

type UserAPIPort interface {
	CreateUser(*dto.CreateUserDTO) (*dto.UserDTO, error)
}

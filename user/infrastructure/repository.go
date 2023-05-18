package infrastructure

import (
	"tutu-gin/user/domain"
	"tutu-gin/user/infrastructure/dto"
)

type Repository interface {
	GetByConfig(config *dto.ConfigDto) (domain.User, error)
	Detail() domain.User
}

package infrastructure

import (
	"tutu-gin/user/domain/entity"
	"tutu-gin/user/infrastructure/dto"
)

type Mapper interface {
	GetByConfig(config *dto.ConfigDto) (entity.User, error)
}

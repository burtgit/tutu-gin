package repository

import (
	"tutu-gin/user/adapter/mapper"
	"tutu-gin/user/domain"
	"tutu-gin/user/infrastructure"
	"tutu-gin/user/infrastructure/dto"
)

type Imp struct {
	mapper infrastructure.Mapper
	user   domain.User
}

func (i *Imp) GetByConfig(config *dto.ConfigDto) (domain.User, error) {
	var obj domain.User
	entity, err := i.mapper.GetByConfig(config)
	obj.User = entity

	return obj, err
}

func (i *Imp) Detail() domain.User {
	return i.user
}

func NewRepository(config *dto.ConfigDto) *Imp {
	obj := &Imp{
		mapper: mapper.NewDbMapper(),
	}

	if config != nil {
		obj.user, _ = obj.GetByConfig(config)
	}

	return obj
}

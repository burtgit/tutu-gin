package mapper

import (
	"tutu-gin/core/global"
	"tutu-gin/user/domain/entity"
	"tutu-gin/user/infrastructure/dto"
)

type DbMapper struct{}

func (d *DbMapper) GetByConfig(config *dto.ConfigDto) (entity.User, error) {
	var obj entity.User

	query := global.DB.Alias("o")

	if config.Id > 0 {
		query = query.Where("id = ?", config.Id)
	}

	if len(config.Token) > 0 {
		query = query.Where("token = ?", config.Token)
	}

	_, err := query.Get(&obj)

	return obj, err
}

func NewDbMapper() *DbMapper {
	return &DbMapper{}
}

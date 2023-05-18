package user

import (
	"tutu-gin/user/adapter/repository"
	"tutu-gin/user/domain"
	"tutu-gin/user/infrastructure"
	"tutu-gin/user/infrastructure/dto"
)

type User struct {
	repository infrastructure.Repository
}

func (u *User) Detail() domain.User {
	return u.repository.Detail()
}

func NewUserService(config *dto.ConfigDto) *User {
	return &User{
		repository: repository.NewRepository(config),
	}
}

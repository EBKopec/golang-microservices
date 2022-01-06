package services

import (
	"github.com/evertonkopec/golang-microservices-main/mvc/domain"
	"github.com/evertonkopec/golang-microservices-main/mvc/utils"
)


type userService struct {}

var (
	UsersService userService
)

func (u *userService)GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	user, err := domain.UserDao.GetUser(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}
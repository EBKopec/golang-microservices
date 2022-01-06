package domain

import (
	"fmt"
	"log"
	"net/http"

	"github.com/evertonkopec/golang-microservices-main/mvc/utils"
)

var (
	users = map[int64]*User{
		123:{ Id:123, FirstName: "Everton", LastName: "Kopec", Email: "evertonkopec@hotmail.com"},
	}
	UserDao usersDaoInterface
)
func init(){
	UserDao = &userDao{}
}
type usersDaoInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

type userDao struct {}

func (u *userDao)GetUser( userId int64) (*User, *utils. ApplicationError) {
	log.Println("We are accessing the database")
	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message: fmt.Sprintf("user: %v does not exists!", userId),
		StatusCode: http.StatusNotFound,
		Code: "not_found",
	}
}



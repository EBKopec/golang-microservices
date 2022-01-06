package app

import (
	"github.com/evertonkopec/golang-microservices-main/mvc/controller"
)

func mapUrls(){
	router.GET("/users/:user_id", controller.GetUser)
}
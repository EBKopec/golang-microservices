package app

import (
	"github.com/evertonkopec/golang-microservices-main/src/api/controllers/polo"
	"github.com/evertonkopec/golang-microservices-main/src/api/controllers/repositories"
)

func mapUrls(){
	router.GET("/marco", polo.Marco)
	router.POST("/repository", repositories.CreateRepo)
	router.POST("/repositories", repositories.CreateRepos)
}
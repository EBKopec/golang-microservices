package app

import (
	"github.com/evertonkopec/golang-microservices-main/oauth-api/src/api/controllers/oauth"
	"github.com/evertonkopec/golang-microservices-main/src/api/controllers/polo"
)

func mapUrls() {
	router.GET("/marco", polo.Marco)
	router.POST("/oauth/access_token", oauth.CreateAccessToken)
	router.GET("/oauth/access_token/:token_id", oauth.GetAccessToken)
}

package services

import (
	"github.com/evertonkopec/golang-microservices-main/oauth-api/src/api/domain/oauth"
	"github.com/evertonkopec/golang-microservices-main/src/api/utils/errors"
	"time"
)

type oauthService struct{}

type oauthServiceInterface interface {
	CreateAccessToken(request oauth.AccessTokenRequest) (*oauth.AccessToken, errors.ApiError)
	GetAccessToken(accessToken string) (*oauth.AccessToken, errors.ApiError)
}

var (
	OauthService oauthServiceInterface
)

func init() {
	OauthService = &oauthService{}
}

func (s *oauthService) CreateAccessToken(request oauth.AccessTokenRequest) (*oauth.AccessToken, errors.ApiError) {
	if err := request.Validate(); err != nil {
		return nil, err
	}
	user, err := oauth.GetUserByUsernameAndPassword(request.Username, request.Password)
	if err != nil {
		return nil, err
	}
	token := oauth.AccessToken{
		UserId:  user.Id,
		Expires: time.Now().UTC().Add(24 * time.Hour).Unix(),
	}
	if err := token.Save(); err != nil {
		return nil, err
	}
	return &token, nil
}

func (s *oauthService) GetAccessToken(accessToken string) (*oauth.AccessToken, errors.ApiError) {
	token, err := oauth.GetAccessTokenByToken(accessToken)
	if err != nil {
		return nil, err
	}
	if token.IsExpired(){
		return nil, errors.NewNotFoundApiError("no access token found for given information")
	}
	return token, err
}

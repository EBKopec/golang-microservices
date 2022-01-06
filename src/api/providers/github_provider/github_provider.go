package github_provider

import (
	"encoding/json"
	"fmt"
	"github.com/evertonkopec/golang-microservices-main/src/api/clients/restclient"
	"github.com/evertonkopec/golang-microservices-main/src/api/domain/github"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	headerAuthorization = "Authorization"
	headerAuthorizationFormat = "token %s"

	urlCreateRepo = "https://api.github.com/user/repos"
)

func getAuthorizationHeader(accessToken string) string {
	return fmt.Sprintf(headerAuthorizationFormat, accessToken)
}


func CreateRepo(accessToken string, request github.CreateRepoRequest) (*github.CreateRepoResponse, *github.GithubErrorResponse){
	//Authorization: token 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
	headers := http.Header{}
	headers.Set(headerAuthorization, getAuthorizationHeader(accessToken))

	response, err := restclient.Post(urlCreateRepo, request, headers)
	if err != nil {
		log.Println(fmt.Sprintf("error when trying to create new repo in github: %s", err.Error()))
		return nil, &github.GithubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message: err.Error()}
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, &github.GithubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message: "invalid response body",
		}
	}
	defer response.Body.Close()

	if response.StatusCode > 299 {
		var errResponse github.GithubErrorResponse
		if err := json.Unmarshal(bytes, &errResponse); err != nil {
			return nil, &github.GithubErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Message: "invalid json response body",
			}
		}
		errResponse.StatusCode = response.StatusCode
		return nil, &errResponse
	}

	var result github.CreateRepoResponse
	if err:= json.Unmarshal(bytes, &result); err != nil {
		log.Println(fmt.Sprintf("error when trying to unmarshal create repo successul response: %s", err.Error()))
		return nil, &github.GithubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message: "error when trying to unmarshal github create repo response",
		}
	}

	return &result, nil

}
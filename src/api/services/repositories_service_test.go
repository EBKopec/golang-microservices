package services

import (
	"github.com/evertonkopec/golang-microservices-main/src/api/clients/restclient"
	"github.com/evertonkopec/golang-microservices-main/src/api/domain/repositories"
	"github.com/evertonkopec/golang-microservices-main/src/api/utils/errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"sync"
	"testing"
)

func TestMain(m *testing.M) {
	restclient.StartMockups()
	os.Exit(m.Run())
}

func TestCreateRepoInvalidInputName(t *testing.T) {
	request := repositories.CreateRepoRequest{}

	result, err := RepositoryService.CreateRepo(request)
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status())
	assert.EqualValues(t, "invalid repository name", err.Message())
}

func TestCreateRepoErrorFromGithub(t *testing.T) {
	restclient.FlushMocks()
	restclient.AddMockups(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		BodyText: `{ "message": "Requires authentication", "documentation_url": "https://docs.github.com/rest/reference/repos#create-a-repository-for-the-authenticated-user"}`,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
		},
	})
	request := repositories.CreateRepoRequest{Name: "testings"}

	result, err := RepositoryService.CreateRepo(request)
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusUnauthorized, err.Status())
	assert.EqualValues(t, "Requires authentication", err.Message())
}

func TestCreateRepoNoError(t *testing.T) {
	restclient.FlushMocks()
	restclient.AddMockups(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		BodyText:   `{ "id": 123, "name":"testing", "owner":{"login":"EBKopec" }}`,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
		},
	})
	request := repositories.CreateRepoRequest{Name: "testings"}

	result, err := RepositoryService.CreateRepo(request)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.EqualValues(t, 123, result.Id)
	assert.EqualValues(t, "testing", result.Name)
	assert.EqualValues(t, "EBKopec", result.Owner)
}

func TestCreateRepoConcurrentInvalidRequest(t *testing.T) {
	request := repositories.CreateRepoRequest{}

	output := make(chan repositories.CreateRepositoriesResult)

	service := reposService{}
	go service.createRepoConcurrent(request, output)

	result := <-output
	assert.NotNil(t, result)
	assert.Nil(t, result.Response)
	assert.NotNil(t, result.Error)
	assert.EqualValues(t, http.StatusBadRequest, result.Error.Status())
	assert.EqualValues(t, "invalid repository name", result.Error.Message())
}

func TestCreateRepoConcurrentErrorFromGithub(t *testing.T) {
	restclient.FlushMocks()
	restclient.AddMockups(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		BodyText: `{ "message": "Requires authentication", "documentation_url": "https://docs.github.com/rest/reference/repos#create-a-repository-for-the-authenticated-user"}`,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
		},
	})
	request := repositories.CreateRepoRequest{Name: "testing"}

	output := make(chan repositories.CreateRepositoriesResult)

	service := reposService{}
	go service.createRepoConcurrent(request, output)

	result := <-output
	assert.NotNil(t, result)
	assert.Nil(t, result.Response)
	assert.NotNil(t, result.Error)
	assert.EqualValues(t, http.StatusUnauthorized, result.Error.Status())
	assert.EqualValues(t, "Requires authentication", result.Error.Message())
}

func TestCreateRepoConcurrentNoError(t *testing.T) {
	restclient.FlushMocks()
	restclient.AddMockups(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		BodyText:   `{ "id": 123, "name":"testing", "owner":{"login":"EBKopec" }}`,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
		},
	})
	request := repositories.CreateRepoRequest{Name: "testing"}

	output := make(chan repositories.CreateRepositoriesResult)

	service := reposService{}
	go service.createRepoConcurrent(request, output)

	result := <-output
	assert.NotNil(t, result)
	assert.Nil(t, result.Error)
	assert.NotNil(t, result.Response)
	assert.EqualValues(t, 123, result.Response.Id)
	assert.EqualValues(t, "testing", result.Response.Name)
	assert.EqualValues(t, "EBKopec", result.Response.Owner)
}

func TestHandleRepoResults(t *testing.T) {
	input := make(chan repositories.CreateRepositoriesResult)
	output := make(chan repositories.CreateReposResponse)

	var wg sync.WaitGroup

	service := reposService{}
	go service.handleRepoResults(&wg, input, output)

	wg.Add(1)
	go func() {
		input <- repositories.CreateRepositoriesResult{
			Error: errors.NewBadRequestError("invalid repository name"),
		}
	}()

	wg.Wait()
	close(input)

	result := <-output
	assert.NotNil(t, result)
	assert.EqualValues(t, 0, result.StatusCode)

	assert.EqualValues(t, 1, len(result.Results))
	assert.NotNil(t, result.Results[0].Error)
	assert.EqualValues(t, http.StatusBadRequest, result.Results[0].Error.Status())
	assert.EqualValues(t, "invalid repository name", result.Results[0].Error.Message())
}

func TestCreatReposInvalidRequests(t *testing.T) {
	request := []repositories.CreateRepoRequest{
		{},
		{Name: "   "},
	}

	result, err := RepositoryService.CreateRepos(request)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.EqualValues(t, http.StatusBadRequest, result.StatusCode)
	assert.EqualValues(t, 2, len(result.Results))

	assert.Nil(t, result.Results[0].Response)
	assert.EqualValues(t, http.StatusBadRequest, result.Results[0].Error.Status())
	assert.EqualValues(t, "invalid repository name", result.Results[0].Error.Message())

	assert.Nil(t, result.Results[1].Response)
	assert.EqualValues(t, http.StatusBadRequest, result.Results[1].Error.Status())
	assert.EqualValues(t, "invalid repository name", result.Results[1].Error.Message())
}

func TestCreatReposOneSuccessOneFail(t *testing.T) {
	restclient.FlushMocks()
	restclient.AddMockups(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		BodyText:   `{ "id": 123, "name":"testing", "owner":{"login":"EBKopec" }}`,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
		},
	})
	request := []repositories.CreateRepoRequest{
		{},
		{Name: "testing"},
	}

	result, err := RepositoryService.CreateRepos(request)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.EqualValues(t, http.StatusPartialContent, result.StatusCode)
	assert.EqualValues(t, 2, len(result.Results))

	for _, result := range result.Results {
		if result.Error != nil {
			assert.EqualValues(t, http.StatusBadRequest, result.Error.Status())
			assert.EqualValues(t, "invalid repository name", result.Error.Message())
			continue
		}

		assert.EqualValues(t, 123, result.Response.Id)
		assert.EqualValues(t, "testing", result.Response.Name)
		assert.EqualValues(t, "EBKopec", result.Response.Owner)
	}
}

func TestCreatReposAllSuccess(t *testing.T) {
	restclient.FlushMocks()
	restclient.AddMockups(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		BodyText:   `{ "id": 123, "name":"testing", "owner":{"login":"EBKopec" }}`,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
		},
	})
	requests := []repositories.CreateRepoRequest{
		{Name: "testing"},
		{Name: "testing"},
	}

	result, err := RepositoryService.CreateRepos(requests)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.EqualValues(t, http.StatusCreated, result.StatusCode)
	assert.EqualValues(t, 2, len(result.Results))

	assert.EqualValues(t, 123, result.Results[0].Response.Id)
	assert.EqualValues(t, "testing", result.Results[0].Response.Name)
	assert.EqualValues(t, "EBKopec", result.Results[0].Response.Owner)

	assert.EqualValues(t, 123, result.Results[1].Response.Id)
	assert.EqualValues(t, "testing", result.Results[1].Response.Name)
	assert.EqualValues(t, "EBKopec", result.Results[1].Response.Owner)

}

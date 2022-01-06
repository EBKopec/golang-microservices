package main

import (
	"bufio"
	"fmt"
	"github.com/evertonkopec/golang-microservices-main/src/api/domain/repositories"
	"github.com/evertonkopec/golang-microservices-main/src/api/services"
	"github.com/evertonkopec/golang-microservices-main/src/api/utils/errors"
	"os"
	"sync"
)

var (
	success map[string]string
	failed  map[string]errors.ApiError
)

type createRepoResult struct {
	Request repositories.CreateRepoRequest
	Result  *repositories.CreateRepoResponse
	Error   errors.ApiError
}

func getRequest() []repositories.CreateRepoRequest {
	result := make([]repositories.CreateRepoRequest, 0)

	file, err := os.Open("/home/everton-kopec/Documentos/Treinamentos/Golang/request.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		request := repositories.CreateRepoRequest{
			Name: line,
		}
		result = append(result, request)
	}

	return result
}

func main() {
	requests := getRequest()

	fmt.Println(fmt.Sprintf("about to process %d requests", len(requests)))

	input := make(chan createRepoResult)
	buffer := make(chan bool, 10)
	var wg sync.WaitGroup

	go handleResults(&wg, input)

	for _, request := range requests {
		buffer <- true
		wg.Add(1)
		go createRepo(buffer, input, request)
	}

	wg.Wait()
	close(input)

	// Now you can write success or failed maps to disk or notify them via email or anything you need to do.
}

func handleResults(wg *sync.WaitGroup, input chan createRepoResult) {
	for result := range input {
		if result.Error != nil {
			failed[result.Request.Name] = result.Error
		}else{
			success[result.Request.Name] = result.Result.Name
		}
		wg.Done()
	}
}

func createRepo(buffer chan bool, output chan createRepoResult, request repositories.CreateRepoRequest) {
	result, err := services.RepositoryService.CreateRepo(request)

	output <- createRepoResult{
		Request: request,
		Result:  result,
		Error:   err,
	}
	<-buffer
}

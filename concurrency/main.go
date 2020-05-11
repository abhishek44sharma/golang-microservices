package main

import (
	"bufio"
	"fmt"
	"golang-microservices/src/api/domain/repositories"
	"golang-microservices/src/api/services"
	"golang-microservices/src/api/utils/errors"
	"os"
	"sync"
)

var (
	success map[string]string
	failed  map[string]errors.ApiError
)

type CreateRepoResult struct {
	Request repositories.CreateRepoRequest
	Result  *repositories.CreateRepoResponse
	Error   errors.ApiError
}

func getRequests() []repositories.CreateRepoRequest {
	result := make([]repositories.CreateRepoRequest, 0)

	file, err := os.Open("/Users/abhishek/Desktop/Requests.txt")
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
	requests := getRequests()
	fmt.Printf(fmt.Sprintf("about to process %d requests", len(requests)))
	var wg sync.WaitGroup

	input := make(chan CreateRepoResult)
	buffer := make(chan bool, 10)

	go handleResults(&wg, input)

	for _, request := range requests {
		buffer <- true
		wg.Add(1)
		go CreateRepo(buffer, request, input)
	}
	wg.Wait()
	close(input)
	close(buffer)
	fmt.Println("Success: ", len(success))
	fmt.Println("Failure: ", len(failed))
}

func handleResults(wg *sync.WaitGroup, input chan CreateRepoResult) {
	success = make(map[string]string)
	failed = make(map[string]errors.ApiError)
	for result := range input {
		if result.Error != nil {
			failed[result.Request.Name] = result.Error
		} else {
			success[result.Request.Name] = result.Result.Name
		}
		wg.Done()
	}
}

func CreateRepo(buffer chan bool, request repositories.CreateRepoRequest, output chan CreateRepoResult) {
	result, err := services.RepositoryService.CreateRepo("xyz", request)
	output <- CreateRepoResult{
		Request: request,
		Result:  result,
		Error:   err,
	}
	<-buffer
}

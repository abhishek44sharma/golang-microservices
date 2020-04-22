package github_provider

import (
	"encoding/json"
	"fmt"
	"golang-microservices/src/api/clients/restclient"
	"golang-microservices/src/api/domain/github"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	headerAuthorization       = "Authorization"
	headerAuthorizationFormat = "token %s"
	urlCreateRepo             = "https://api.github.com/user/repos"
)

func GetAuthorizationHeader(accessToken string) string {
	return fmt.Sprintf(headerAuthorizationFormat, accessToken)
}

func CreateRepo(accessToken string, request github.CreateRepoRequest) (*github.CreateRepoResponse, *github.GitHubErrorResponse) {
	header := http.Header{}
	header.Set(headerAuthorization, GetAuthorizationHeader(accessToken))

	response, err := restclient.Post(urlCreateRepo, request, header)

	if err != nil {
		log.Println(fmt.Sprintf("Error when trying to create new repo in github: %s", err.Error()))
		return nil, &github.GitHubErrorResponse{StatusCode: http.StatusInternalServerError, Message: err.Error()}
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, &github.GitHubErrorResponse{StatusCode: http.StatusInternalServerError, Message: "invalid response body"}
	}

	defer response.Body.Close()

	if response.StatusCode > 299 {
		var errResponse github.GitHubErrorResponse
		if err := json.Unmarshal(bytes, &errResponse); err != nil {
			return nil, &github.GitHubErrorResponse{StatusCode: http.StatusInternalServerError, Message: "invalid json resonse body"}
		}
		errResponse.StatusCode = response.StatusCode
		return nil, &errResponse
	}

	var result github.CreateRepoResponse

	if err := json.Unmarshal(bytes, &result); err != nil {
		log.Println(fmt.Sprintf("error when trying to unmarshal successful response: %s", err.Error()))
		return nil, &github.GitHubErrorResponse{StatusCode: http.StatusInternalServerError, Message: "error when trying to unmarshal github create repo response"}
	}

	return &result, nil
}

package services

import (
	"golang-microservices/src/api/clients/restclient"
	"golang-microservices/src/api/domain/repositories"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
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

	assert.EqualValues(t, http.StatusBadRequest, err.Get_Status())
	assert.EqualValues(t, "invalid repository name", err.Get_Message())
}

func TestCreateRepoErrorFromGithub(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message": "Requires authentication","documentation_url": "https://developer.github.com/docs"}`)),
		},
	})

	request := repositories.CreateRepoRequest{
		Name: "testing",
	}

	result, err := RepositoryService.CreateRepo(request)
	assert.NotNil(t, err)
	assert.Nil(t, result)
	assert.EqualValues(t, http.StatusUnauthorized, err.Get_Status())
	assert.EqualValues(t, "Requires authentication", err.Get_Message())
}

func TestCreateRepoNoError(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id": 123, "name": "testing", "owner": {"login": "abhishek44sharma"}}`)),
		},
	})

	request := repositories.CreateRepoRequest{
		Name: "testing",
	}

	result, err := RepositoryService.CreateRepo(request)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.EqualValues(t, 123, result.Id)
	assert.EqualValues(t, "testing", result.Name)
	assert.EqualValues(t, "abhishek44sharma", result.Owner)
}

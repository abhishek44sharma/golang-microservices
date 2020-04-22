package github

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRepoTest(t *testing.T) {
	request := CreateRepoRequest{
		Name:        "golang tutorial",
		Description: "a golang introduction repository",
		Homepage:    "https://github.com",
		Private:     true,
		HasIssues:   true,
		HasProjects: true,
		HasWiki:     true,
	}

	bytes, err := json.Marshal(request)

	assert.Nil(t, err)
	assert.NotNil(t, bytes)

	target := CreateRepoRequest{}
	err = json.Unmarshal(bytes, &target)

	assert.Nil(t, err)
	assert.NotNil(t, target)
	assert.EqualValues(t, target, request)
}

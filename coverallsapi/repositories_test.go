package coverallsapi

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestRepositoryServiceGet(t *testing.T) {
	repository := &Repository{
		ID:       123,
		Name:     "user/fakerepo",
		HasBadge: true,
		Token:    "fake-repo-token",
	}
	responder, _ := httpmock.NewJsonResponder(200, repository)
	fakeUrl := "https://coveralls.io/api/repos/github/user/fakerepo"
	httpmock.RegisterResponder("GET", fakeUrl, responder)

	client := NewClient("fake token")
	httpmock.ActivateNonDefault(client.client.GetClient())
	defer httpmock.DeactivateAndReset()

	result, err := client.Repositories.Get(context.Background(), "github", "user/fakerepo")

	assert.Nil(t, err)
	assert.Equal(t, repository, result)
}

func TestRepositoryServiceAdd(t *testing.T) {
	repositoryConfig := &RepositoryConfig{
		Service:                         "github",
		Name:                            "user/fakerepo",
		CommentOnPullRequests:           true,
		SendBuildStatus:                 false,
		CommitStatusFailThreshold:       10.3,
		CommitStatusFailChangeThreshold: 9.1,
	}
	fakeUrl := "https://coveralls.io/api/repos"
	httpmock.RegisterResponder("POST", fakeUrl, func(req *http.Request) (*http.Response, error) {
		cfg := make(map[string]*RepositoryConfig)
		if err := json.NewDecoder(req.Body).Decode(&cfg); err != nil {
			return httpmock.NewStringResponse(400, ""), nil
		}

		assert.Equal(t, repositoryConfig, cfg["repo"])

		resp, err := httpmock.NewJsonResponse(200, cfg["repo"])
		if err != nil {
			return httpmock.NewStringResponse(500, ""), nil
		}
		return resp, nil
	})

	client := NewClient("fake token")
	httpmock.ActivateNonDefault(client.client.GetClient())
	defer httpmock.DeactivateAndReset()

	result, err := client.Repositories.Add(context.Background(), repositoryConfig)

	assert.Nil(t, err)
	assert.Equal(t, repositoryConfig, result)
}

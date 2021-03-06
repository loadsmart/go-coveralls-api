/*
Copyright (c) 2020 Loadsmart, Inc.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package coveralls

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
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
	failThreshold := 10.3
	repositoryConfig := &RepositoryConfig{
		Service:                         "github",
		Name:                            "user/fakerepo",
		CommentOnPullRequests:           true,
		SendBuildStatus:                 false,
		CommitStatusFailThreshold:       &failThreshold,
		CommitStatusFailChangeThreshold: nil,
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

func TestRepositoryConfigMarshall(t *testing.T) {
	var testCases = []struct {
		name string
		in   RepositoryConfig
		out  string
	}{
		{
			name: "simple",
			in:   RepositoryConfig{Service: "github", Name: "user/fakerepo"},
			out:  `{"service": "github", "name": "user/fakerepo", "send_build_status": false, "comment_on_pull_requests": false, "commit_status_fail_change_threshold": null, "commit_status_fail_threshold": null}`,
		},
		{
			name: "partial",
			in:   RepositoryConfig{Service: "github", Name: "user/fakerepo", SendBuildStatus: true, CommitStatusFailThreshold: pfloat64(10.3)},
			out:  `{"service": "github", "name": "user/fakerepo", "send_build_status": true, "comment_on_pull_requests": false, "commit_status_fail_change_threshold": null, "commit_status_fail_threshold": 10.3}`,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			content, err := json.Marshal(tt.in)
			assert.Nil(t, err)
			assert.JSONEq(t, tt.out, string(content))
		})
	}
}

func ExampleRepositoryService_Get() {
	// Instantiate the client with your _personal access token_
	client := NewClient("your-personal-access-token")
	// This returns information about a specific repository
	repository, err := client.Repositories.Get(context.Background(), "github", "user/repository")
	if err != nil {
		log.Fatalf("Error querying Coveralls API: %s\n", err)
	}

	fmt.Printf("Project has ID %d in Coveralls", repository.ID)
}

func pfloat64(v float64) *float64 {
	return &v
}

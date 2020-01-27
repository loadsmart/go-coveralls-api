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
	"fmt"
)

// RepositoryService holds information to access repository-related endpoints
type RepositoryService service

// Repository holds information about one specific repository
type Repository struct {
	ID       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	HasBadge bool   `json:"has_badge,omitempty"`
	Token    string `json:"token,omitempty"`
}

// RepositoryConfig represents config settings for a given repository
type RepositoryConfig struct {
	Service                         string  `json:"service"`                                       // Git provider. Options include: github, bitbucket, gitlab, stash, manual
	Name                            string  `json:"name"`                                          // Name of the repo. E.g. with Github, this is username/reponame.
	CommentOnPullRequests           bool    `json:"comment_on_pull_requests,omitempty"`            // Whether comments should be posted on pull requests (defaults to true)
	SendBuildStatus                 bool    `json:"send_build_status,omitempty"`                   // Whether build status should be sent to the git provider (defaults to true)
	CommitStatusFailThreshold       float64 `json:"commit_status_fail_threshold,omitempty"`        // Minimum coverage that must be present on a build for the build to pass (default is null, meaning any decrease is a failure)
	CommitStatusFailChangeThreshold float64 `json:"commit_status_fail_change_threshold,omitempty"` // If coverage decreases, the maximum allowed amount of decrease that will be allowed for the build to pass (default is null, meaning that any decrease is a failure)
}

// Get information about a repository already in Coveralls.
//
// Ctx is a context that's propagated to underlying client. You can use
// it to cancel the request in progress (when the program is terminated,
// for example).
//
// Svc indicates the service. Any value accepted by Coveralls API can be
// passed here. Soma valid inputs include 'github', 'bitbucket' or 'manual'.
//
// Repo is the repository name. In GitHub, for example, this is
// 'organization/repository'; other services could have different formats.
//
// If the request suceeded, it returns a Repository with the information
// available or an error if there was something wrong.
func (s *RepositoryService) Get(ctx context.Context, svc string, repo string) (*Repository, error) {
	url := fmt.Sprintf("%s/repos/%s/%s", s.client.baseURL, svc, repo)

	resp, err := s.client.client.R().
		SetContext(ctx).
		SetResult(&Repository{}).
		Get(url)

	if err != nil {
		return nil, err
	}

	return resp.Result().(*Repository), nil
}

// Add a repository to Coveralls
func (s *RepositoryService) Add(ctx context.Context, data *RepositoryConfig) (*RepositoryConfig, error) {
	url := fmt.Sprintf("%s/repos", s.client.baseURL)

	body := map[string]*RepositoryConfig{
		"repo": data,
	}

	resp, err := s.client.client.R().
		SetContext(ctx).
		SetBody(body).
		SetResult(&RepositoryConfig{}).
		Post(url)

	if err != nil {
		return nil, err
	}

	return resp.Result().(*RepositoryConfig), nil
}

// Update repository configuration in Coveralls
func (s *RepositoryService) Update(ctx context.Context, svc string, repo string, data *RepositoryConfig) (*RepositoryConfig, error) {
	url := fmt.Sprintf("%s/repos/%s/%s", s.client.baseURL, svc, repo)

	body := map[string]*RepositoryConfig{
		"repo": data,
	}

	resp, err := s.client.client.R().
		SetContext(ctx).
		SetBody(body).
		SetResult(&RepositoryConfig{}).
		Put(url)

	if err != nil {
		return nil, err
	}

	return resp.Result().(*RepositoryConfig), nil
}

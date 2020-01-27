package coverallsapi

import (
	"context"
	"fmt"
)

type RepositoryService service

type Repository struct {
	ID       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	HasBadge bool   `json:"has_badge,omitempty"`
	Token    string `json:"token,omitempty"`
}

type RepositoryConfig struct {
	Service                         string  `json:"service"`                                       // Git provider. Options include: github, bitbucket, gitlab, stash, manual
	Name                            string  `json:"name"`                                          // Name of the repo. E.g. with Github, this is username/reponame.
	CommentOnPullRequests           bool    `json:"comment_on_pull_requests,omitempty"`            // Whether comments should be posted on pull requests (defaults to true)
	SendBuildStatus                 bool    `json:"send_build_status,omitempty"`                   // Whether build status should be sent to the git provider (defaults to true)
	CommitStatusFailThreshold       float64 `json:"commit_status_fail_threshold,omitempty"`        // Minimum coverage that must be present on a build for the build to pass (default is null, meaning any decrease is a failure)
	CommitStatusFailChangeThreshold float64 `json:"commit_status_fail_change_threshold,omitempty"` // If coverage decreases, the maximum allowed amount of decrease that will be allowed for the build to pass (default is null, meaning that any decrease is a failure)
}

// Get information about a repository already in Coveralls
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

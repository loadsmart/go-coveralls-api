package coverallsapi

import (
	"fmt"
	"net/url"

	"github.com/go-resty/resty/v2"
)

const (
	defaultBaseURL = "https://coveralls.io/api"
)

type Client struct {
	client  *resty.Client
	baseURL *url.URL
	common  service

	Repositories *RepositoryService
}

type service struct {
	client *Client
}

// NewClient returns a new Coveralls API Client
// t is the Coveralls API token
func NewClient(t string) *Client {
	cli := resty.New()
	cli.SetHeader("Accept", "application/json")
	cli.SetHeader("Authorization", fmt.Sprintf("token %s", t))

	url, _ := url.Parse(defaultBaseURL)
	c := &Client{client: cli, baseURL: url}
	c.common.client = c

	c.Repositories = (*RepositoryService)(&c.common)
	return c
}

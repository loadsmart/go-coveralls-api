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

package coverallsapi

import (
	"fmt"
	"net/url"

	"github.com/go-resty/resty/v2"
)

const (
	defaultBaseURL = "https://coveralls.io/api"
)

// Client is used to provide a single interface to interact with Coveralls API
type Client struct {
	client  *resty.Client
	baseURL *url.URL // Base URL for Coveralls API
	common  service  // Share the same client instance among all services

	Repositories *RepositoryService // Service to interact with repository-related endpoints
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

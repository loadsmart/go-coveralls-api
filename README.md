# go-coveralls-api

[![GoDoc](https://godoc.org/github.com/loadsmart/go-coveralls-api?status.svg)](https://godoc.org/github.com/loadsmart/go-coveralls-api)

Client for [Coveralls API][] written in Go.

**Note**: the goal is to interact with administrative Coveralls API. To send coverage data, take a look at [goveralls][] project.

## Installation

Just follow the usual instructions for Go libraries:

```bash
go get github.com/loadsmart/go-coveralls-api
```

`go-coveralls-api` uses Go Modules and therefore requires Go 1.11+.

## Usage

To get the ID of a repo already configured in Coveralls

```go
import (
    "context"
    "fmt"
    "log"

    "github.com/loadsmart/go-coveralls-api"
)

client := coveralls.NewClient("your-personal-access-token")
repo, err := client.Repositories.Get(context.Background(), "github", "user/repository"))
if err != nil {
    log.Fatalf("Error querying Coveralls API: %s\n", err)
}

fmt.Printf("Project has ID %s in Coveralls", resp.ID)
```

Replace `your-personal-access-token` with your personal access token (can be found in your Coveralls account page).

## License

This work is copyrighted to Loadsmart, Inc. and licensed under MIT. For details see [LICENSE][] file.

[Coveralls API]: https://docs.coveralls.io/api-introduction
[goveralls]: https://github.com/mattn/goveralls
[LICENSE]: ./LICENSE

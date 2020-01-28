package coveralls

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClientWithAuthorizationHeader(t *testing.T) {
	client := NewClient("my-personal-token")

	authHeader := client.client.Header.Get("Authorization")
	assert.Equal(t, "token my-personal-token", authHeader)
}

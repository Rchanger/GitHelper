package mocks

import (
	"context"

	"github.com/google/go-github/github"
)

type MockGetDS struct {
	UserName string
	User     *github.User
	Response *github.Response
	Err      error
}

func (m *MockGetDS) Get(ctx context.Context, user string) (*github.User, *github.Response, error) {
	m.UserName = user
	return m.User, m.Response, m.Err
}

package mocks

import (
	"context"

	"github.com/google/go-github/github"
)

type MockCreateDS struct {
	Owner, Repo string
	Pull        *github.NewPullRequest
	Result      *github.PullRequest
	Response    *github.Response
	Err         error
}

func (m *MockCreateDS) Create(ctx context.Context, owner, repo string, pull *github.NewPullRequest) (*github.PullRequest, *github.Response, error) {
	m.Owner = owner
	m.Repo = repo
	m.Pull = pull
	return m.Result, m.Response, m.Err
}

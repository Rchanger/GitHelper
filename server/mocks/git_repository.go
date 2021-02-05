package mocks

import (
	"context"

	"github.com/google/go-github/github"
)

type MockListDS struct {
	User       string
	Repository []*github.Repository
	Response   *github.Response
	Err        error
}

func (m *MockListDS) List(ctx context.Context, user string, _ *github.RepositoryListOptions) ([]*github.Repository, *github.Response, error) {
	m.User = user
	return m.Repository, m.Response, m.Err
}

type MockGetCommitDS struct {
	Owner, Repo, Sha string
	Commit           *github.RepositoryCommit
	Response         *github.Response
	Err              error
}

func (m *MockGetCommitDS) GetCommit(ctx context.Context, owner, repo, sha string) (*github.RepositoryCommit, *github.Response, error) {
	m.Owner = owner
	m.Repo = repo
	m.Sha = sha
	return m.Commit, m.Response, m.Err
}

type MockGetBranchesDS struct {
	Owner, Repo string
	Branches    []*github.Branch
	Opt         *github.ListOptions
	Response    *github.Response
	Err         error
}

func (m *MockGetBranchesDS) ListBranches(ctx context.Context, owner string, repo string, opt *github.ListOptions) ([]*github.Branch, *github.Response, error) {
	m.Owner = owner
	m.Repo = repo
	m.Opt = opt
	return m.Branches, m.Response, m.Err
}

type GitRespositoryDS struct {
	*MockListDS
	*MockGetCommitDS
	*MockGetBranchesDS
}

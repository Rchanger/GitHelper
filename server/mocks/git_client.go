package mocks

import (
	"context"

	"github.com/google/go-github/github"
)

type MockGetRefDS struct {
	Author   string
	Repo     string
	BaseRef  string
	Ref      *github.Reference
	Response *github.Response
	Err      error
}

func (m *MockGetRefDS) GetRef(ctx context.Context, author, repo string, ref string) (*github.Reference, *github.Response, error) {
	m.BaseRef = ref
	m.Author = author
	m.Repo = repo
	return m.Ref, m.Response, m.Err
}

type MockCreateRefDS struct {
	Author   string
	Repo     string
	BaseRef  *github.Reference
	Ref      *github.Reference
	Response *github.Response
	Err      error
}

func (m *MockCreateRefDS) CreateRef(ctx context.Context, author, repo string, ref *github.Reference) (*github.Reference, *github.Response, error) {
	m.BaseRef = ref
	m.Author = author
	m.Repo = repo
	return m.Ref, m.Response, m.Err
}

type MockCreateTreeDS struct {
	Owner    string
	Repo     string
	BaseTree string
	Entries  []github.TreeEntry
	Tree     *github.Tree
	Response *github.Response
	Err      error
}

func (m *MockCreateTreeDS) CreateTree(ctx context.Context, owner, repo, baseTree string, entries []github.TreeEntry) (*github.Tree, *github.Response, error) {
	m.Owner = owner
	m.Repo = repo
	m.BaseTree = baseTree
	m.Entries = entries
	return m.Tree, m.Response, m.Err
}

type MockCreateCommitDS struct {
	Owner    string
	Repo     string
	Result   *github.Commit
	Commit   *github.Commit
	Response *github.Response
	Err      error
}

func (m *MockCreateCommitDS) CreateCommit(ctx context.Context, owner, repo string, commit *github.Commit) (*github.Commit, *github.Response, error) {
	m.Commit = commit
	m.Owner = owner
	m.Repo = repo
	return m.Result, m.Response, m.Err
}

type MockUpdateRefDS struct {
	Owner    string
	Repo     string
	BaseRef  *github.Reference
	Ref      *github.Reference
	Force    bool
	Response *github.Response
	Err      error
}

func (m *MockUpdateRefDS) UpdateRef(ctx context.Context, owner, repo string, ref *github.Reference, force bool) (*github.Reference, *github.Response, error) {
	m.Owner = owner
	m.Repo = repo
	m.BaseRef = ref
	m.Force = force
	return m.Ref, m.Response, m.Err
}

type GitClientDS struct {
	*MockCreateRefDS
	*MockGetRefDS
	*MockCreateTreeDS
	*MockCreateCommitDS
	*MockUpdateRefDS
}

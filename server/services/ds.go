package services

import (
	"context"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type DS struct {
	GitClientDS      GitClientDS
	GitRepositoryDS  GitRepositoryDS
	GitPullRequestDS GitPullRequestDS
	AuthConfigDS     AuthConfigDS
	GitUserDS        GitUserDS
}

type GitClientDS interface {
	GetRef(ctx context.Context, author, repo string, ref string) (*github.Reference, *github.Response, error)
	CreateRef(ctx context.Context, author, repo string, ref *github.Reference) (*github.Reference, *github.Response, error)
	CreateTree(ctx context.Context, owner, repo, baseTree string, entries []github.TreeEntry) (*github.Tree, *github.Response, error)
	CreateCommit(ctx context.Context, owner, repo string, commit *github.Commit) (*github.Commit, *github.Response, error)
	UpdateRef(ctx context.Context, owner, repo string, ref *github.Reference, force bool) (*github.Reference, *github.Response, error)
}

type GitRepositoryDS interface {
	List(ctx context.Context, user string, opts *github.RepositoryListOptions) ([]*github.Repository, *github.Response, error)
	ListBranches(ctx context.Context, owner string, repo string, opt *github.ListOptions) ([]*github.Branch, *github.Response, error)
	GetCommit(ctx context.Context, owner, repo, sha string) (*github.RepositoryCommit, *github.Response, error)
}

type GitPullRequestDS interface {
	Create(ctx context.Context, owner, repo string, pull *github.NewPullRequest) (*github.PullRequest, *github.Response, error)
}

type GitUserDS interface {
	Get(ctx context.Context, user string) (*github.User, *github.Response, error)
}

type AuthConfigDS interface {
	Exchange(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error)
}

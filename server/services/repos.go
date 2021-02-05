package services

import (
	"context"

	"github.com/google/go-github/github"
)

func (ds DS) GetRepos(ctx context.Context) ([]*github.Repository, error) {
	repos, _, err := ds.GitRepositoryDS.List(ctx, "", nil)
	if err != nil {
		return nil, err
	}
	return repos, nil
}

func (ds DS) GetBranches(ctx context.Context, owner, repo string) ([]*github.Branch, error) {
	branches, _, err := ds.GitRepositoryDS.ListBranches(ctx, owner, repo, nil)
	if err != nil {
		return nil, err
	}
	return branches, nil
}

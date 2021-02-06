package services

import (
	"GitHelper/server/models"
	"context"

	"github.com/google/go-github/github"
)

func (ds DS) NewPullRequest(ctx context.Context, params models.PullRequestSubmitParams) (*github.PullRequest, error) {
	prParms := &github.NewPullRequest{
		Title:               &params.Title,
		Head:                &params.Newbranch,
		Base:                &params.BaseBranch,
		Body:                &params.Description,
		MaintainerCanModify: github.Bool(true),
	}

	newPR, _, err := ds.GitPullRequestDS.Create(ctx, params.Author, params.Repo, prParms)
	if err != nil {
		return nil, err
	}

	return newPR, nil
}

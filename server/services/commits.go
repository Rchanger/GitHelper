package services

import (
	"GitHelper/server/models"
	"context"
	"fmt"
	"time"

	"github.com/google/go-github/github"
)

// GetCommit - 70
func (ds DS) FetchCommit(ctx context.Context, author, repo string, ref *github.Reference) (*github.RepositoryCommit, error) {
	c, _, err := ds.GitRepositoryDS.GetCommit(ctx, author, repo, *ref.Object.SHA)
	if err != nil {
		return nil, err
	}

	c.Commit.SHA = c.SHA
	return c, nil
}

// CreateCommit - 84
func (ds DS) CreateNewCommit(ctx context.Context, params models.CommitSubmitParams) (*github.Commit, error) {
	commitAuthor := &github.CommitAuthor{
		Date:  &params.CommitDateTime,
		Name:  &params.Author,
		Email: &params.Email,
	}
	commit := &github.Commit{
		Author:  commitAuthor,
		Message: &params.CommitMsg,
		Tree:    params.Tree,
		Parents: []github.Commit{*params.ParentCommit.Commit},
	}

	newCommit, _, err := ds.GitClientDS.CreateCommit(ctx, params.Author, params.Repo, commit)
	if err != nil {
		return nil, err
	}

	return newCommit, nil
}

func (ds DS) CreatePushCommit(ctx context.Context, params models.CommitSubmitParams) error {
	params.CommitDateTime = time.Now()
	branchRef, err := ds.FetchBranchRef(ctx, params.Author, params.Repo, models.HeadRef+params.Branch)
	if err != nil {
		fmt.Println("FetchBranchRef", err)
		return err
	}
	tree, err := ds.CreateNewBranchTree(ctx, params.Author, params.Repo, branchRef, params.Files)
	if err != nil {
		fmt.Println("CreateNewBranchTree", err)
		return err
	}
	params.Tree = tree

	parentCommit, err := ds.FetchCommit(ctx, params.Author, params.Repo, branchRef)
	if err != nil {
		fmt.Println("FetchCommit", err)
		return err
	}
	params.ParentCommit = parentCommit

	newCommit, err := ds.CreateNewCommit(ctx, params)
	if err != nil {
		fmt.Println("CreateNewCommit", err)
		return err
	}

	branchRef.Object.SHA = newCommit.SHA

	err = ds.AttachCommitToRef(ctx, params.Author, params.Repo, branchRef, false)
	if err != nil {
		fmt.Println("AttachCommitToRef", err)
		return err
	}

	return nil
}

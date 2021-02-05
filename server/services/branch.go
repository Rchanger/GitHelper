package services

import (
	"GitHelper/server/models"
	"context"
	"io/ioutil"
	"path/filepath"

	"github.com/google/go-github/github"
)

// GetRef - 34
func (ds DS) FetchBranchRef(ctx context.Context, author, repo string, branch string) (*github.Reference, error) {
	ref, _, err := ds.GitClientDS.GetRef(ctx, author, repo, branch)
	if err != nil {
		return nil, err
	}
	return ref, nil
}

// CreateRef - 44
func (ds DS) CreateNewBranch(ctx context.Context, author, repo string, ref *github.Reference) (*github.Reference, error) {
	ref, _, err := ds.GitClientDS.CreateRef(ctx, author, repo, ref)
	if err != nil {
		return nil, err
	}
	return ref, nil
}

// CreateTree -63
func (ds DS) CreateNewBranchTree(ctx context.Context, author, repo string, newBranchRef *github.Reference, modifiedFilesPath []string) (*github.Tree, error) {
	modifiedFilesEntries := []github.TreeEntry{}
	for _, path := range modifiedFilesPath {
		bs, err := ioutil.ReadFile(models.ClonePath + repo + "/" + path)
		if err != nil {
			return nil, err
		}
		fileName := filepath.Base(path)
		modifiedFilesEntries = append(modifiedFilesEntries, github.TreeEntry{
			Path:    github.String(fileName),
			Type:    github.String("blob"),
			Content: github.String(string(bs)),
			Mode:    github.String("100644"),
		})
	}

	tree, _, err := ds.GitClientDS.CreateTree(ctx, author, repo, *newBranchRef.Object.SHA, modifiedFilesEntries)
	if err != nil {
		return nil, err
	}
	return tree, nil
}

//UpdateRef - 94
func (ds DS) AttachCommitToRef(ctx context.Context, author, repo string, ref *github.Reference, force bool) error {
	_, _, err := ds.GitClientDS.UpdateRef(ctx, author, repo, ref, force)
	return err
}

package services

import (
	"context"

	"github.com/google/go-github/github"
)

func (ds DS) GetUser(ctx context.Context) (*github.User, error) {
	user, _, err := ds.GitUserDS.Get(ctx, "")
	if err != nil {
		return nil, err
	}
	return user, nil
}

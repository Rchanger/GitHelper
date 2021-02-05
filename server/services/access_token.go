package services

import (
	"context"

	"golang.org/x/oauth2"
)

func (ds DS) GetAccessToken(ctx context.Context, code string) (*oauth2.Token, error) {
	return ds.AuthConfigDS.Exchange(ctx, code)
}

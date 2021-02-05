package models

import (
	"context"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

const (
	Scopes        = "public_repo"
	DataStructure = "DS"
	MasterRef     = "refs/heads/master"
	HeadRef       = "heads/"
	ClonePath     = "/home/repos/"
)

var (
	Config     AppConfig
	Token      = ""
	Outfile, _ = os.OpenFile("git_helper.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
)

// GetConfigFilePath - returns config filepath
func GetConfigFilePath() string {
	return "config/config.toml"
}

func NewAuthConfig() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     Config.ClientId,
		ClientSecret: Config.ClientSecret,
		Scopes:       []string{Scopes},
		Endpoint: oauth2.Endpoint{
			TokenURL: Config.TokenURL,
			AuthURL:  Config.AuthURL,
		},
	}
}

func NewGithubClient(ctx context.Context) *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: Token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	return client

}

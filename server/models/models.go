package models

import (
	"time"

	"github.com/google/go-github/github"
)

type CommitRequestParams struct {
	Branch        string   `json:"branch,omitempty"`
	CommitMessage string   `json:"commit_message,omitempty"`
	Repo          string   `json:"repo,omitempty"`
	Files         []string `json:"files,omitempty"`
}

type CommitSubmitParams struct {
	Branch         string
	Author         string
	Email          string
	CommitMsg      string
	Repo           string
	CommitDateTime time.Time
	Tree           *github.Tree
	ParentCommit   *github.RepositoryCommit
	Files          []string
}

type PullRequestParams struct {
	Repo        string `json:"repo,omitempty"`
	Description string `json:"description,omitempty"`
	Title       string `json:"title,omitempty"`
	BaseBranch  string `json:"base_branch,omitempty"`
	NewBranch   string `json:"new_branch,omitempty"`
}

type PullRequestSubmitParams struct {
	Author      string
	Repo        string
	Description string
	Email       string
	Title       string
	BaseBranch  string
	Newbranch   string
	CanModify   bool
}

type BranchParams struct {
	Repo      string `json:"repo,omitempty"`
	NewBranch string `json:"new_branch,omitempty"`
}

package services

import (
	"GitHelper/server/mocks"
	"context"
	"errors"
	"testing"

	"github.com/google/go-github/github"
	"github.com/stretchr/testify/assert"
)

func Test_GetRepos_DS(t *testing.T) {

	t.Run("happy path", func(t *testing.T) {

		repoName := "test"

		repos := []*github.Repository{
			&github.Repository{Name: &repoName},
		}

		m := mocks.GitRespositoryDS{
			MockListDS: &mocks.MockListDS{
				Repository: repos,
			},
		}

		ds := DS{
			GitRepositoryDS: m,
		}
		ctx := context.Background()
		repoRespose, err := ds.GetRepos(ctx)
		assert.NoError(t, err)
		assert.Equal(t, repos, repoRespose)

	})

	t.Run("sad path - get error", func(t *testing.T) {

		m := mocks.GitRespositoryDS{
			MockListDS: &mocks.MockListDS{Err: errors.New("expectation error")},
		}
		ds := DS{
			GitRepositoryDS: m,
		}
		ctx := context.Background()
		_, err := ds.GetRepos(ctx)
		assert.Error(t, errors.New("expectation error"), err)
	})
}

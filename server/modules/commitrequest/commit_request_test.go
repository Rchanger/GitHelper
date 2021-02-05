package commitrequest

import (
	"GitHelper/server/mocks"
	"GitHelper/server/models"
	"GitHelper/server/services"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-github/github"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func Test_CreatePushCommit(t *testing.T) {
	e := echo.New()

	t.Run("happy path", func(t *testing.T) {
		loggedInUser := "test_user"

		m := mocks.GitClientDS{
			MockCreateRefDS: &mocks.MockCreateRefDS{},
			MockGetRefDS: &mocks.MockGetRefDS{
				Ref: &github.Reference{
					Object: &github.GitObject{
						SHA: &loggedInUser,
					},
				},
			},
			MockCreateTreeDS: &mocks.MockCreateTreeDS{},
			MockUpdateRefDS:  &mocks.MockUpdateRefDS{},
			MockCreateCommitDS: &mocks.MockCreateCommitDS{
				Result: &github.Commit{
					SHA: &loggedInUser,
				},
			},
		}

		m2 := mocks.GitRespositoryDS{
			MockGetCommitDS: &mocks.MockGetCommitDS{
				Commit: &github.RepositoryCommit{
					Commit: &github.Commit{
						SHA: &loggedInUser,
					},
					SHA: &loggedInUser,
				},
			},
		}

		loggedInUserEmail := "test_user"
		user := github.User{
			Login: &loggedInUser,
			Email: &loggedInUserEmail,
		}

		m1 := &mocks.MockGetDS{
			User: &user,
		}

		ds := services.DS{
			GitClientDS:     m,
			GitUserDS:       m1,
			GitRepositoryDS: m2,
		}

		ctx := context.Background()
		ctx = context.WithValue(ctx, models.DataStructure, ds)
		branch := models.CommitRequestParams{
			Branch:        "test",
			CommitMessage: "Test",
			Repo:          "test",
			Files:         []string{"/home/srokade/workspace/goworkspace/src/GitHelper/server/services/commits.go"},
		}
		requestByte, _ := json.Marshal(branch)
		requestReader := bytes.NewReader(requestByte)
		req := httptest.NewRequest(http.MethodPost, "/commit", requestReader)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		req = req.WithContext(ctx)

		echoCtx := e.NewContext(req, rec)

		CreatePushCommit(echoCtx)
		assert.Equal(t, http.StatusCreated, rec.Code)

	})

	t.Run("sad path - get error", func(t *testing.T) {
		loggedInUser := "test_user"
		loggedInUserEmail := "test_user"
		m := mocks.GitClientDS{
			MockCreateRefDS: &mocks.MockCreateRefDS{},
			MockGetRefDS: &mocks.MockGetRefDS{
				Ref: &github.Reference{
					Object: &github.GitObject{
						SHA: &loggedInUser,
					},
				},
			},
			MockCreateTreeDS:   &mocks.MockCreateTreeDS{},
			MockUpdateRefDS:    &mocks.MockUpdateRefDS{},
			MockCreateCommitDS: &mocks.MockCreateCommitDS{Err: errors.New("expectation error")},
		}

		m2 := mocks.GitRespositoryDS{
			MockGetCommitDS: &mocks.MockGetCommitDS{
				Commit: &github.RepositoryCommit{
					Commit: &github.Commit{
						SHA: &loggedInUser,
					},
					SHA: &loggedInUser,
				},
			},
		}
		user := github.User{
			Login: &loggedInUser,
			Email: &loggedInUserEmail,
		}

		m1 := &mocks.MockGetDS{
			User: &user,
		}

		ds := services.DS{
			GitClientDS:     m,
			GitUserDS:       m1,
			GitRepositoryDS: m2,
		}

		ctx := context.Background()
		ctx = context.WithValue(ctx, "DS", ds)

		branch := models.CommitRequestParams{
			Branch:        "test",
			CommitMessage: "Test",
			Repo:          "test",
			Files:         []string{"/home/srokade/workspace/goworkspace/src/GitHelper/server/services/commits.go"},
		}
		requestByte, _ := json.Marshal(branch)
		requestReader := bytes.NewReader(requestByte)
		req := httptest.NewRequest(http.MethodPost, "/commit", requestReader)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		req = req.WithContext(ctx)

		echoCtx := e.NewContext(req, rec)

		CreatePushCommit(echoCtx)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})

}

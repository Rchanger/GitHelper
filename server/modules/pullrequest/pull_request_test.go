package pullrequest

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

func Test_CreatePullRequest(t *testing.T) {
	e := echo.New()

	t.Run("happy path", func(t *testing.T) {
		loggedInUser := "test_user"

		m := &mocks.MockCreateDS{}

		loggedInUserEmail := "test_user"
		user := github.User{
			Login: &loggedInUser,
			Email: &loggedInUserEmail,
		}

		m1 := &mocks.MockGetDS{
			User: &user,
		}

		ds := services.DS{
			GitPullRequestDS: m,
			GitUserDS:        m1,
		}

		ctx := context.Background()
		ctx = context.WithValue(ctx, models.DataStructure, ds)
		pull := models.PullRequestParams{
			Title:       "Test pull",
			Repo:        "test",
			BaseBranch:  "dev",
			NewBranch:   "test",
			Description: "test pull",
		}
		requestByte, _ := json.Marshal(pull)
		requestReader := bytes.NewReader(requestByte)
		req := httptest.NewRequest(http.MethodPost, "/pullRequest", requestReader)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		req = req.WithContext(ctx)

		echoCtx := e.NewContext(req, rec)

		CreatePullRequest(echoCtx)
		assert.Equal(t, http.StatusCreated, rec.Code)

	})

	t.Run("sad path - get error", func(t *testing.T) {
		loggedInUser := "test_user"
		loggedInUserEmail := "test_user"

		m := &mocks.MockCreateDS{Err: errors.New("expectation error")}

		user := github.User{
			Login: &loggedInUser,
			Email: &loggedInUserEmail,
		}

		m1 := &mocks.MockGetDS{
			User: &user,
		}

		ds := services.DS{
			GitPullRequestDS: m,
			GitUserDS:        m1,
		}

		ctx := context.Background()
		ctx = context.WithValue(ctx, "DS", ds)

		pull := models.PullRequestParams{
			Title:       "Test pull",
			Repo:        "test",
			BaseBranch:  "dev",
			NewBranch:   "test",
			Description: "test pull",
		}
		requestByte, _ := json.Marshal(pull)
		requestReader := bytes.NewReader(requestByte)
		req := httptest.NewRequest(http.MethodPost, "/pullRequest", requestReader)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		req = req.WithContext(ctx)

		echoCtx := e.NewContext(req, rec)

		CreatePullRequest(echoCtx)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})

}

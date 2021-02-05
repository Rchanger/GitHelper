package branch

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

	"github.com/stretchr/testify/require"

	"github.com/google/go-github/github"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func Test_CreateBranch(t *testing.T) {
	e := echo.New()

	t.Run("happy path", func(t *testing.T) {

		loggedInUser := "test_user"
		user := github.User{
			Login: &loggedInUser,
		}
		m := mocks.GitClientDS{
			MockCreateRefDS: &mocks.MockCreateRefDS{},
			MockGetRefDS: &mocks.MockGetRefDS{
				Ref: &github.Reference{
					Object: &github.GitObject{
						SHA: &loggedInUser,
					},
				},
			},
		}
		m1 := &mocks.MockGetDS{
			User: &user,
		}
		ds := services.DS{
			GitClientDS: m,
			GitUserDS:   m1,
		}

		ctx := context.Background()
		ctx = context.WithValue(ctx, models.DataStructure, ds)
		branch := models.BranchParams{
			Repo:      "test",
			NewBranch: "dev",
		}
		requestByte, err := json.Marshal(branch)
		require.NoError(t, err)
		requestReader := bytes.NewReader(requestByte)
		req := httptest.NewRequest(http.MethodPost, "/branch", requestReader)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		req = req.WithContext(ctx)

		echoCtx := e.NewContext(req, rec)

		CreateBranch(echoCtx)
		assert.Equal(t, http.StatusOK, rec.Code)

	})

	t.Run("sad path - get error", func(t *testing.T) {

		loggedInUser := "test_user"
		user := github.User{
			Login: &loggedInUser,
		}
		m1 := &mocks.MockGetDS{
			User: &user,
		}
		m := mocks.GitClientDS{
			MockCreateRefDS: &mocks.MockCreateRefDS{Err: errors.New("expectation error")},
			MockGetRefDS: &mocks.MockGetRefDS{
				Ref: &github.Reference{
					Object: &github.GitObject{
						SHA: &loggedInUser,
					},
				},
			},
		}
		ds := services.DS{
			GitClientDS: m,
			GitUserDS:   m1,
		}

		ctx := context.Background()
		ctx = context.WithValue(ctx, "DS", ds)

		branch := models.BranchParams{
			Repo:      "test",
			NewBranch: "dev",
		}
		requestByte, _ := json.Marshal(branch)
		requestReader := bytes.NewReader(requestByte)
		req := httptest.NewRequest(http.MethodPost, "/branch", requestReader)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		req = req.WithContext(ctx)

		echoCtx := e.NewContext(req, rec)

		CreateBranch(echoCtx)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})

}

package repository

import (
	"GitHelper/server/mocks"
	"GitHelper/server/models"
	"GitHelper/server/services"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func Test_GetRepos(t *testing.T) {
	e := echo.New()

	t.Run("happy path", func(t *testing.T) {

		m := mocks.GitRespositoryDS{
			MockListDS: &mocks.MockListDS{},
		}

		ds := services.DS{
			GitRepositoryDS: m,
		}

		ctx := context.Background()
		ctx = context.WithValue(ctx, models.DataStructure, ds)

		req := httptest.NewRequest(http.MethodGet, "/repos", nil)
		rec := httptest.NewRecorder()
		req = req.WithContext(ctx)

		echoCtx := e.NewContext(req, rec)

		GetRepos(echoCtx)
		assert.Equal(t, http.StatusOK, rec.Code)

	})

	t.Run("sad path - get error", func(t *testing.T) {

		m := mocks.GitRespositoryDS{
			MockListDS: &mocks.MockListDS{Err: errors.New("expectation error")},
		}
		ds := services.DS{
			GitRepositoryDS: m,
		}

		ctx := context.Background()
		ctx = context.WithValue(ctx, "DS", ds)

		req := httptest.NewRequest(http.MethodGet, "/repos", nil)
		rec := httptest.NewRecorder()
		req = req.WithContext(ctx)

		echoCtx := e.NewContext(req, rec)

		GetRepos(echoCtx)
		assert.Equal(t, http.StatusExpectationFailed, rec.Code)
	})

}

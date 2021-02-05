package authentication

import (
	"GitHelper/server/mocks"
	"GitHelper/server/models"
	"GitHelper/server/services"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func Test_SaveAccessToken(t *testing.T) {
	e := echo.New()

	t.Run("happy path", func(t *testing.T) {
		m := &mocks.MockExchangeDS{}
		ds := services.DS{
			AuthConfigDS: m,
		}

		ctx := context.Background()
		ctx = context.WithValue(ctx, models.DataStructure, ds)

		req := httptest.NewRequest(http.MethodPost, "/auth/getToken?code=foo", nil)
		rec := httptest.NewRecorder()
		req = req.WithContext(ctx)

		echoCtx := e.NewContext(req, rec)

		SaveAccessToken(echoCtx)
		assert.Equal(t, http.StatusPermanentRedirect, rec.Code)

	})

	t.Run("sad path - exchange error", func(t *testing.T) {

		m := &mocks.MockExchangeDS{Err: assert.AnError}
		ds := services.DS{
			AuthConfigDS: m,
		}

		ctx := context.Background()
		ctx = context.WithValue(ctx, "DS", ds)

		req := httptest.NewRequest(http.MethodPost, "/auth/getToken?code=foo", nil)
		rec := httptest.NewRecorder()
		req = req.WithContext(ctx)

		echoCtx := e.NewContext(req, rec)

		SaveAccessToken(echoCtx)
		assert.Equal(t, http.StatusExpectationFailed, rec.Code)
	})

}

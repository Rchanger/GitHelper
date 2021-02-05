package authentication

import (
	"GitHelper/server/common"
	"GitHelper/server/models"
	"net/http"

	"github.com/labstack/echo"
	"golang.org/x/oauth2"
)

func Init(o *echo.Group) {
	o.GET("/auth", AuthenticateUser)
	o.GET("/auth/token", SaveAccessToken)
}

func AuthenticateUser(c echo.Context) error {
	authURL := models.NewAuthConfig().AuthCodeURL("state", oauth2.AccessTypeOffline)
	return c.Redirect(http.StatusPermanentRedirect, authURL)
}

func SaveAccessToken(c echo.Context) error {
	ctx := c.Request().Context()

	code := c.FormValue("code")
	if code == "" {
		c.Logger().Error("unable to get access code")
		return c.JSON(http.StatusForbidden, "unable to get access code")
	}
	ds, err := common.GetServiceDS(ctx)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	token, err := ds.GetAccessToken(ctx, code)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusExpectationFailed, err)
	}

	models.Token = token.AccessToken
	return c.Redirect(http.StatusPermanentRedirect, models.Config.ClientRedirectURL)
}

package repository

import (
	"GitHelper/server/common"
	"net/http"

	"github.com/labstack/echo"
)

func Init(o *echo.Group) {
	o.GET("/repos", GetRepos)
	o.GET("/repos/:id/branches", GetBranches)

}

func GetRepos(c echo.Context) error {
	ctx := c.Request().Context()
	ds, err := common.GetServiceDS(ctx)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusExpectationFailed, err)
	}
	repos, err := ds.GetRepos(ctx)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusExpectationFailed, err)
	}
	return c.JSON(http.StatusOK, repos)
}

func GetBranches(c echo.Context) error {
	ctx := c.Request().Context()
	repo := c.Param("id")
	ds, err := common.GetServiceDS(ctx)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusExpectationFailed, err)
	}
	c.Logger().Error("Invalid token")
	user, err := ds.GetUser(ctx)
	if err != nil {
		c.Logger().Error("Invalid token")
		return c.JSON(http.StatusUnauthorized, "Invalid token")
	}
	branches, err := ds.GetBranches(ctx, *user.Login, repo)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusExpectationFailed, err)
	}
	return c.JSON(http.StatusOK, branches)
}

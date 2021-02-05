package pullrequest

import (
	"GitHelper/server/common"
	"GitHelper/server/models"
	"net/http"

	"github.com/labstack/echo"
)

func Init(o *echo.Group) {
	o.POST("/pullRequest", CreatePullRequest)
}

func CreatePullRequest(c echo.Context) error {
	ctx := c.Request().Context()

	params := models.PullRequestParams{}

	if err := c.Bind(&params); err != nil {
		c.Logger().Error("Invalid request params")
		return c.JSON(http.StatusBadRequest, "invalid request params")
	}

	params.BaseBranch = "master"
	ds, err := common.GetServiceDS(ctx)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	user, err := ds.GetUser(ctx)
	if err != nil {
		c.Logger().Error("Invalid token")
		return c.JSON(http.StatusUnauthorized, "Invalid token")
	}

	submitRequest := models.PullRequestSubmitParams{
		Repo:        params.Repo,
		Newbranch:   params.NewBranch,
		Title:       params.Title,
		Description: params.Description,
		Email:       *user.Email,
		Author:      *user.Login,
		BaseBranch:  "master",
	}

	prURL, err := ds.NewPullRequest(ctx, submitRequest)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, prURL)
}

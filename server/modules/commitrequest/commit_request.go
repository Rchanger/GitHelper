package commitrequest

import (
	"GitHelper/server/common"
	"GitHelper/server/models"
	"net/http"

	"github.com/labstack/echo"
)

func Init(o *echo.Group) {
	o.POST("/commit", CreatePushCommit)
}

func CreatePushCommit(c echo.Context) error {
	ctx := c.Request().Context()
	params := models.CommitRequestParams{}

	if err := c.Bind(&params); err != nil {
		c.Logger().Error("Invalid request params")
		return c.JSON(http.StatusBadRequest, "Invalid request params")
	}

	ds, err := common.GetServiceDS(ctx)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusExpectationFailed, err)
	}

	user, err := ds.GetUser(ctx)
	if err != nil {
		c.Logger().Error("Invalid token")
		return c.JSON(http.StatusUnauthorized, "Invalid token")
	}

	submitRequest := models.CommitSubmitParams{
		Email:     *user.Email,
		Author:    *user.Login,
		Branch:    params.Branch,
		CommitMsg: params.CommitMessage,
		Files:     params.Files,
		Repo:      params.Repo,
	}

	err = ds.CreatePushCommit(ctx, submitRequest)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, "successfully created commit")
}

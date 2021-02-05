package branch

import (
	"GitHelper/server/common"
	"GitHelper/server/models"
	"net/http"

	"github.com/google/go-github/github"
	"github.com/labstack/echo"
)

func Init(o *echo.Group) {
	o.POST("/branch", CreateBranch)

}

func CreateBranch(c echo.Context) error {
	ctx := c.Request().Context()
	params := models.BranchParams{}

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
	masterRef, err := ds.FetchBranchRef(ctx, *user.Login, params.Repo, models.MasterRef)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	newBranchRef := &github.Reference{Ref: github.String("refs/heads/" + params.NewBranch), Object: &github.GitObject{SHA: masterRef.Object.SHA}}
	_, err = ds.CreateNewBranch(ctx, *user.Login, params.Repo, newBranchRef)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, "Success")
}

package middleware

import (
	"GitHelper/server/models"
	"GitHelper/server/services"
	"context"
	"net/http"

	"github.com/labstack/echo"
)

func CustomContextBuilder(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()
		reqCtx := req.Context()
		authConfig := models.NewAuthConfig()

		serviceDS := services.DS{
			AuthConfigDS: authConfig,
		}
		if req.URL.Path != "/auth" && req.URL.Path != "/auth/token" && models.Token == "" {
			c.Logger().Errorf("unauthorized access to %s", req.URL.Path)
			return c.JSON(http.StatusUnauthorized, "unauthorized user")
		}
		if models.Token != "" {
			gitClient := models.NewGithubClient(reqCtx)
			serviceDS.GitClientDS = gitClient.Git
			serviceDS.GitRepositoryDS = gitClient.Repositories
			serviceDS.GitPullRequestDS = gitClient.PullRequests
			serviceDS.GitUserDS = gitClient.Users
		}
		ctx := context.WithValue(reqCtx, models.DataStructure, serviceDS)
		newReq := req.WithContext(ctx)
		c.SetRequest(newReq)
		return next(c)
	}
}

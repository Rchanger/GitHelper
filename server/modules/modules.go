package modules

import (
	"GitHelper/server/modules/authentication"
	"GitHelper/server/modules/branch"
	"GitHelper/server/modules/commitrequest"
	"GitHelper/server/modules/pullrequest"
	"GitHelper/server/modules/repository"

	"github.com/labstack/echo"
)

func Init(o *echo.Group) {
	authentication.Init(o)
	repository.Init(o)
	branch.Init(o)
	pullrequest.Init(o)
	commitrequest.Init(o)

}

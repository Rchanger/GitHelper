package routes

import (
	"GitHelper/server/modules"

	"github.com/labstack/echo"
)

func Init(e *echo.Echo) {
	o := e.Group("")

	modules.Init(o)
}

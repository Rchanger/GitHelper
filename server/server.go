package main

import (
	"GitHelper/server/middleware"
	"GitHelper/server/models"
	"GitHelper/server/routes"
	"net/http"

	"github.com/BurntSushi/toml"
	"github.com/labstack/echo"
	echoMW "github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CustomContextBuilder)
	e.Use(echoMW.CORSWithConfig(echoMW.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))
	e.Use(echoMW.LoggerWithConfig(echoMW.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
		Output: models.Outfile,
	}))
	e.Logger.SetOutput(models.Outfile)
	e.Logger.SetLevel(1)
	e.Logger.SetHeader("[${time_rfc3339_nano}-${level}-${short_file}-${line}]")

	if err := setconfig(); err != nil {
		panic(err)
	}

	routes.Init(e)
	e.HideBanner = true
	if err := e.Start(":" + models.Config.AppPort); err != nil {
		panic(err)
	}
}

// setconfig - reads config file and assigns to config model
func setconfig() error {
	_, err := toml.DecodeFile(models.GetConfigFilePath(), &models.Config)
	if err != nil {
		return err
	}
	return nil
}

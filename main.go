package main

import (
	"kasir_api/config"
	"kasir_api/routes"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func main() {
	// ECHO INSTANCE
	e := echo.New()

	// MIDDLEWARES
	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	//ROUTES
	routes.APIRoutes(e)

	// VIPER CONFIG
	appConfig := config.ViperConfig()

	// ASSIGN TO VARIABLES
	api_host := appConfig.APIHost
	api_port := appConfig.APIPort

	// START SERVER
	e.Logger.Fatal(e.Start(api_host + ":" + api_port))
}

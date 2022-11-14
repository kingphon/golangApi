package main

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"golangApi/config"
	_ "golangApi/docs"
	"golangApi/module/database"
	"golangApi/route"
)

func init() {
	config.Init()
	database.Connect()
}

// @title Swagger Document Admin API
// @version 1.0
// @description Document Admin server
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	envVars := config.GetEnv()
	e := echo.New()

	// Route ...
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	route.Route(e)

	// Start server
	e.Logger.Fatal(e.Start(envVars.AppPort))
}

package main

import (
	"github.com/Poul-george/go-api/api/echoset"
	"github.com/Poul-george/go-api/api/presentation/controller/users"
	"github.com/Poul-george/go-api/api/request_get"
	"github.com/labstack/echo/v4/middleware"
)

var e = echoset.CreateMux()

func main() {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/users", request_get.ResponesUsers)
	e.POST("/users", users.Controller)
	e.Logger.Fatal(e.Start(":1324"))
}

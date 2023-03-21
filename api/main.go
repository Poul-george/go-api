package main

import (
	"fmt"

	"github.com/Poul-george/go-api/api/config"
)

// "github.com/Poul-george/go-api/api/infrastructure/echo/echoset"
// "github.com/Poul-george/go-api/api/presentation/controller/user/create"
// "github.com/Poul-george/go-api/api/presentation/controller/user/list"
// "github.com/labstack/echo/v4/middleware"

// var e = echoset.CreateMux()

func main() {
	fmt.Println(config.GetConfig())
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	// e.GET("/users", list.Get)
	// e.POST("/users", create.Post)
	// e.Logger.Fatal(e.Start(":1324"))
}

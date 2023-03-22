package main

import "C"
import (
	"fmt"
	"net/http"
	"time"

	"github.com/Poul-george/go-api/api/config"
	"github.com/Poul-george/go-api/api/infrastructure/echo/echoset"
	cmdConfig "github.com/Poul-george/go-api/api/presentation/cmd/config"
)

var e = echoset.CreateMux()

func main() {
	fmt.Println(config.GetConfig())

	c := config.GetServerConfig()
	e := cmdConfig.NewEchoServer()
	s := http.Server{
		Addr:        c.StartAddress,
		Handler:     e,
		IdleTimeout: time.Duration(c.IdleTimeout),
	}
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		fmt.Printf("s.ListenAndServe() Error = %v \n", err.Error())
		e.Logger.Fatal(err)
	}
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	// e.GET("/users", list.Get)
	// e.POST("/users", create.Post)
	// e.Logger.Fatal(e.Start(":1324"))
}

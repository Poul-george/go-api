package main

import (
	"github.com/Poul-george/go-api/api/echoset"
	"github.com/Poul-george/go-api/api/request_get"
)

var e = echoset.CreateMux()

func main() {
	e.GET("/users", request_get.ResponesUsers)
	e.Logger.Fatal(e.Start(":1324"))
}

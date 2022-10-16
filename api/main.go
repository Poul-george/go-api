package main

import (
    "github.com/Poul-george/go-api/api/echoset"
    "github.com/Poul-george/go-api/api/request_get"
    // "fmt"
)

var e = echoset.CreateMux()

func main() {
	e.GET("/", request_get.ArticleIndex)
	e.GET("/json", request_get.GetJson)

	e.Logger.Fatal(e.Start(":1324"))
}

//go run main.go

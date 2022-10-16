package main

import (
	// "net/http"
	// "encoding/json"

	// "github.com/labstack/echo/v4"
	// "github.com/labstack/echo/v4/middleware"
	// "time"

    // "github.com/Poul-george/go_api_app/tree/master/go/api/echoset"
    "api/echoset"
    "fmt"
)

// type Book struct {
// 	Id string
// 	Name string
// 	Author string
// 	Day time.Time
// }

var e = echoset.test()

func main() {
	// e.GET("/", articleIndex)
	// e.GET("/json", getJson)

	// e.Logger.Fatal(e.Start(":1324"))
	fmt.Println(e)
}

// func createMux() *echo.Echo {
// 	e := echo.New()

// 	var allowedOrigins = []string{
//         "http://localhost:1324",
//         "http://localhost:3000",
//     }

// 	// CORS用のmiddlewareはあるものの、403を勝手に返してくれるわけではない。
//     // 以下のレスポンスヘッダを付与する役割がある。
//     // Access-Control-Allow-Credentials: true
//     // Access-Control-Allow-Origin: https://www.example.com:8080
//     e.Use(
//         middleware.CORSWithConfig(middleware.CORSConfig{
//             AllowCredentials: true,
//             AllowOrigins:     allowedOrigins,
//             AllowHeaders: []string{
//                 echo.HeaderAccessControlAllowHeaders,
//                 echo.HeaderContentType,
//                 echo.HeaderContentLength,
//                 echo.HeaderAcceptEncoding,
//                 echo.HeaderXCSRFToken,
//                 echo.HeaderAuthorization,
//             },
//             AllowMethods: []string{
//                 http.MethodGet,
//                 http.MethodPut,
//                 http.MethodPatch,
//                 http.MethodPost,
//                 http.MethodDelete,
//             },
//             MaxAge: 86400,
//         }),
//     )

//     // echoで起動しているAPIサーバーに、Originが不正な場合に403を返却させるには、自分でミドルウェアを書く必要がある
//     e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
//         return func(c echo.Context) error {
//             // Originヘッダの中身を取得
//             origin := c.Request().Header.Get(echo.HeaderOrigin)
//             // 許可しているOriginの中で、リクエストヘッダのOriginと一致するものがあれば処理を継続
//             for _, o := range allowedOrigins {
//                 if origin == o {
//                     return next(c)
//                 }
//             }
//             // 一致しているものがなかった場合は403(Forbidden)を返却する
//             // ginと処理を合わせるなら return c.NoContent(http.StatusForbidden) のがいいかも。
//             return c.String(http.StatusForbidden, "forbidden")
//         }
//     })

// 	e.Use(middleware.Recover())
// 	e.Use(middleware.Logger())
// 	e.Use(middleware.Gzip())

// 	return e
// }

// func articleIndex(c echo.Context) error {
// 	return c.String(http.StatusOK, "hello css html")
// }
// func getJson(c echo.Context) error {
// 	array := make([]Book, 3)
// 	for i := 0;i < len(array);i++ {
// 		array[i].Id = "vol1010tttx7ym9l"
// 		array[i].Name = "C++ programming language"
// 		array[i].Author = "Bjarne Stroutsrup"
// 		array[i].Day = time.Now()
// 	}

//   	res, _ := json.Marshal(array)
// 	return c.String(http.StatusOK, string(res))
// }

//go run main.go

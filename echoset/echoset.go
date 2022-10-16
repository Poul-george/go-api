package echoset

import (
	// "net/http"
	// // "encoding/json"

	// "github.com/labstack/echo/v4"
	// "github.com/labstack/echo/v4/middleware"
	"fmt"
)

func Foo() {
	fmt.Println("Foo!")
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
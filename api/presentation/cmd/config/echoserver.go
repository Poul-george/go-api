package config

//echoのcros設定のファイル airでの実装を考慮

import (
	"github.com/Poul-george/go-api/api/config"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewEchoServer() *echo.Echo {
	e := SetMiddlleware()
	router(e)

	return e
}

func SetMiddlleware() *echo.Echo {
	e := echo.New()
	cnf := config.GetServerConfig()

	// CORS用のmiddlewareはあるものの、403を勝手に返してくれるわけではない。
	e.Use(
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowCredentials: true,
			AllowOrigins:     cnf.EchoAllowOrigin,
			AllowHeaders: []string{
				echo.HeaderAccessControlAllowHeaders,
				echo.HeaderContentType,
				echo.HeaderContentLength,
				echo.HeaderAcceptEncoding,
				echo.HeaderXCSRFToken,
				echo.HeaderAuthorization,
			},
			AllowMethods: []string{
				http.MethodGet,
				http.MethodPut,
				http.MethodPatch,
				http.MethodPost,
				http.MethodDelete,
			},
			MaxAge: 86400,
		}),
	)

	// echoで起動しているAPIサーバーに、Originが不正な場合に403を返却させるには、自分でミドルウェアを書く必要がある
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Originヘッダの中身を取得
			origin := c.Request().Header.Get(echo.HeaderOrigin)
			// // 許可しているOriginの中で、リクエストヘッダのOriginと一致するものがあれば処理を継続
			for _, o := range cnf.EchoAllowOrigin {
				if origin == o {
					return next(c)
				}
			}
			// 一致しているものがなかった場合は403(Forbidden)を返却する
			return c.String(http.StatusForbidden, "forbidden")
		}
	})

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	return e
}

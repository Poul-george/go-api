package config

import (
	appapi "github.com/Poul-george/go-api/api/presentation/appapi/echoserver/router"
	"github.com/labstack/echo/v4"
)

func router(e *echo.Echo) {
	appapi.RouterV1Api(e.Group("api/v1"))
}

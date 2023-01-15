package list

import (
	"net/http"

	usecase "github.com/Poul-george/go-api/api/core/usecase/api/user/list"
	"github.com/labstack/echo/v4"
)

// type Prameter struct {
// 	UserId string `json:"user_id"`
// }

func Get(c echo.Context) error {
	// var p Prameter
	// if err := c.Bind(&p); err != nil {
	// 	return err
	// }

	out, err := usecase.Do()
	if err != nil {
		return c.JSON(http.StatusBadRequest, "リクエストができませんでした。")
	}

	res := NewResponse(*out)

	return c.JSON(http.StatusOK, res)
}

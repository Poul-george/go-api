package create

import (
	"net/http"

	usecase "github.com/Poul-george/go-api/api/core/usecase/api/user/post"
	"github.com/Poul-george/go-api/api/presentation/input"
	"github.com/labstack/echo/v4"
)

type Prameter struct {
	Name        string `json:"name"`
	Password    string `json:"password"`
	MailAddress string `json:"email"`
	Comments    string `json:"comments"`
}

func Post(c echo.Context) error {
	var p Prameter
	if err := c.Bind(&p); err != nil {
		return err
	}

	err := usecase.Do(input.Input{
		Name:        p.Name,
		Password:    p.Password,
		MailAddress: p.MailAddress,
		Comments:    p.Comments,
	})
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, "正常に登録できました。")
}

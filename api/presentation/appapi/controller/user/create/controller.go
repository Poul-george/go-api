package create

import (
	"net/http"

	"github.com/Poul-george/go-api/api/core/usecase/api/user/post"
	usecase "github.com/Poul-george/go-api/api/core/usecase/api/user/post"
	"github.com/labstack/echo/v4"
)

type Prameter struct {
	Name        string `json:"name"`
	Password    string `json:"password"`
	MailAddress string `json:"email"`
	Comments    string `json:"comments"`
}

type Controller struct {
	UseCase usecase.UseCase
}

func (c *Controller) Post(ctx echo.Context) error {
	var p Prameter
	if err := ctx.Bind(&p); err != nil {
		return err
	}

	err := c.UseCase.Do(post.Input{
		Name:        p.Name,
		Password:    p.Password,
		MailAddress: p.MailAddress,
		Comments:    p.Comments,
	})
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, "正常に登録できました。")
}

package list

import (
	"net/http"

	usecase "github.com/Poul-george/go-api/api/core/usecase/api/user/list"
	"github.com/Poul-george/go-api/api/presentation/appapi/echoserver/context"
)

// type Prameter struct {
// 	UserId string `json:"user_id"`
// }

type Controller struct {
	UseCase usecase.UseCase
}

func (c Controller) Get(ctx context.Context) error {
	// var p Prameter
	// if err := c.Bind(&p); err != nil {
	// 	return err
	// }

	out, err := c.UseCase.Do(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "リクエストができませんでした。")
	}

	res := NewResponse(*out)

	return ctx.JSON(http.StatusOK, res)
}

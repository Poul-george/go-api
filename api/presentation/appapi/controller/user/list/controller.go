package list

import (
	"net/http"

	echoContext "github.com/Poul-george/go-api/api/presentation/appapi/echoserver/context"

	usecase "github.com/Poul-george/go-api/api/core/usecase/api/user/list"
)

type Prameter struct {
	UserID         *uint64 `json:"user_id" form:"user_id"`
	ExternalUserID *string `json:"external_user_id" form:"external_user_id"`
}

type Controller struct {
	UseCase usecase.UseCase
}

func (c Controller) Get(ctx echoContext.Context) error {
	var p Prameter
	if err := ctx.Bind(&p); err != nil {
		return err
	}

	out, err := c.UseCase.Do(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "リクエストができませんでした。")
	}

	res := NewResponse(*out)

	return ctx.JSON(http.StatusOK, res)
}

package detail

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Poul-george/go-api/api/core/common/types/identifier"

	echoContext "github.com/Poul-george/go-api/api/presentation/appapi/echoserver/context"

	usecase "github.com/Poul-george/go-api/api/core/usecase/api/user/detail"
)

type Prameter struct {
	UserID         uint64 `json:"user_id" query:"user_id"`
	ExternalUserID string `json:"external_user_id" query:"external_user_id"`
}

type Controller struct {
	UseCase usecase.UseCase
}

func (c Controller) Get(ctx echoContext.Context) error {
	var p Prameter
	if err := ctx.Bind(&p); err != nil {
		fmt.Println("バインドエラー")
		return err
	}

	if p.ExternalUserID == "" && p.UserID == 0 {
		return errors.New("user detail bad request")
	}

	out, err := c.UseCase.Do(ctx.Request().Context(), usecase.Input{
		ExternalUserID: identifier.ExternalUserID(p.ExternalUserID),
		UserID:         identifier.UserID(p.UserID),
	})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "リクエストができませんでした。")
	}

	res := NewResponse(*out)

	return ctx.JSON(http.StatusOK, res)
}

package inject

import (
	userListUseCase "github.com/Poul-george/go-api/api/core/usecase/api/user/list"
	userList "github.com/Poul-george/go-api/api/presentation/appapi/controller/user/list"
	"github.com/labstack/echo/v4"
)

func (i *Injector) V1UserListController() echo.HandlerFunc {
	return newHandlerFunc(userList.Controller{
		UseCase: *userListUseCase.NewUseCase(i.userRepository()),
	}.Get)
}

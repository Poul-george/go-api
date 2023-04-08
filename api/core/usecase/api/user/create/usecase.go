package create

import (
	"context"
	"fmt"

	"github.com/Poul-george/go-api/api/core/domain/model"
	"github.com/Poul-george/go-api/api/core/domain/repository/user"
)

type UseCase struct {
	UserRepository user.Repository
}

func NewUseCase(
	userRepository user.Repository,
) *UseCase {
	return &UseCase{
		UserRepository: userRepository,
	}
}

func (u *UseCase) Do(ctx context.Context, ip Input) error {
	res, err := InputData(ip)
	if err != nil {
		return err
	}

	user, err1 := model.NewUser(res.ExternalUserID, res.Name, res.Password, res.MailAddress, res.Comments)
	if err1 != nil {
		return err1
	}

	fmt.Printf("--------------------{[%v]}----------------------------00000\n", u)

	err2 := u.UserRepository.Create(ctx, user)
	if err2 != nil {
		return err2
	}

	return nil
}

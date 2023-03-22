package post

import (
	"fmt"

	"github.com/Poul-george/go-api/api/core/domain/model"
	// "github.com/Poul-george/go-api/api/infrastructure/repository/user"
	"github.com/Poul-george/go-api/api/core/domain/repository/user"
)

type UseCase struct {
	UserRepository user.Repository
}

func (u *UseCase) Do(ip Input) error {
	res, err := InputData(ip)
	if err != nil {
		return err
	}

	user, err1 := model.NewUser(res.Name, res.Password, res.MailAddress, res.Comments)
	if err1 != nil {
		return err1
	}

	fmt.Printf("--------------------{[%v]}----------------------------00000\n", u)

	err2 := u.UserRepository.Create(user)
	if err2 != nil {
		return err2
	}

	return nil
}

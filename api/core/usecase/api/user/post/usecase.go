package post

import (
	"fmt"

	"github.com/Poul-george/go-api/api/core/domain/model"
	"github.com/Poul-george/go-api/api/infrastructure/repository/user"
	"github.com/Poul-george/go-api/api/presentation/input"
)

func Do(ip input.Input) error {
	res, err := input.InputData(ip)
	if err != nil {
		return err
	}

	u, err1 := model.NewUser(res.Name, res.Password, res.MailAddress, res.Comments)
	if err1 != nil {
		return err1
	}

	fmt.Printf("--------------------{[%v]}----------------------------00000\n", u)

	err2 := user.Create(u)
	if err2 != nil {
		return err2
	}

	return nil
}

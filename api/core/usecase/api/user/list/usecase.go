package list

import (
	"fmt"

	"github.com/Poul-george/go-api/api/infrastructure/repository/user"
)

func Do() (*[]user.Users, error) {
	user, err := user.List()
	fmt.Printf("--------------------{[%v]}----------------------------00000\n", user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

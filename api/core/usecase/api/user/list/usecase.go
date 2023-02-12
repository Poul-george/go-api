package list

import (
	"fmt"

	"github.com/Poul-george/go-api/api/infrastructure/repository/user"
)

func Do() (*Output, error) {
	users, err := user.List()
	fmt.Printf("--------------------{[%v]}----------------------------00000\n", users)

	if err != nil {
		return nil, err
	}

	return NewOutput(users), nil
}

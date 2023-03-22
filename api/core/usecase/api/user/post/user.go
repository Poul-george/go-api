package post

import (
	"github.com/Poul-george/go-api/api/core/domain/model"
)

type Input model.User

func InputData(ip Input) (*Input, error) {
	// var input Input
	// input = ip

	return &ip, nil
}

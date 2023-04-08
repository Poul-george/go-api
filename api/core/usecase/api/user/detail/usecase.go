package detail

import (
	"context"
	"fmt"

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

func (u *UseCase) Do(ctx context.Context, input Input) (*Output, error) {
	user, err := u.UserRepository.FindByID(ctx, input.ExternalUserID, input.UserID)
	fmt.Printf("--------------------{[%v]}----------------------------00000\n", user)

	if err != nil {
		return nil, err
	}

	return NewOutput(*user), nil
}

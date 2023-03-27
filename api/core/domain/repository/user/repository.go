package user

import (
	"context"

	"github.com/Poul-george/go-api/api/core/domain/model"
	"github.com/Poul-george/go-api/api/infrastructure/data/persistence/gorm/table"
)

type Repository interface {
	Create(context.Context, *model.User) error
	List(context.Context) ([]table.User, error)
}

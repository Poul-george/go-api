package user

import (
	"context"

	"github.com/Poul-george/go-api/api/core/domain/model"
	"github.com/Poul-george/go-api/api/infrastructure/data/persistence/gorm/handler"
	"github.com/Poul-george/go-api/api/infrastructure/data/persistence/gorm/table"
)

type Repository struct {
	handler *handler.Handler
}

func NewRepository(h *handler.Handler) *Repository {
	return &Repository{handler: h}
}

func (r *Repository) Create(ctx context.Context, user *model.User) error {
	createUser := ToUserForCreate(*user)
	if err := r.handler.Writer(ctx).Create(createUser).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) List(ctx context.Context) ([]table.User, error) {
	users := []table.User{}
	err := r.handler.Reader(ctx).
		Model(&table.User{}).
		Order("id").
		Find(&users).Error
	if err != nil {
		return nil, nil
	}
	return users, nil
}

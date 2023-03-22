package user

import (
	"context"

	"github.com/Poul-george/go-api/api/core/domain/model"
	"github.com/Poul-george/go-api/api/infrastructure/data/persistence/gorm/handler"
	"github.com/Poul-george/go-api/api/infrastructure/data/persistence/gorm/table"
	"github.com/Poul-george/go-api/api/infrastructure/mysql/mysql_setting"
)

type Repository struct {
	handler *handler.Handler
}

func NewRepository(h *handler.Handler) *Repository {
	return &Repository{handler: h}
}

func (r Repository) Create(user *model.User) error {
	db := mysql_setting.Mysql_Connect()
	err := db.DB.Create(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) List(ctx context.Context) ([]table.User, error) {
	users := []table.User{}
	err := r.handler.Reader(ctx).
		Model(&table.User{}).
		Order("id desc").
		Find(&users).Error
	if err != nil {
		return nil, err
	}
	// db := mysql_setting.Mysql_Connect()
	// err := db.DB.Find(&users).Order("id desc").Error
	// if err != nil {
	// 	return nil, err
	// }

	return users, nil
}
package user

import (
	"github.com/Poul-george/go-api/api/core/domain/model"
	"github.com/Poul-george/go-api/api/infrastructure/data/persistence/gorm/table"
)

func ToUserForCreate(u model.User) *table.User {
	t := table.User{
		ExternalUserID: u.ExternalUserID.String(),
		Name:           u.Name,
		Password:       u.Password,
		MailAddress:    u.MailAddress,
		Comments:       u.Comments,
	}
	return &t
}

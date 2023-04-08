package detail

import (
	"time"

	"github.com/Poul-george/go-api/api/infrastructure/data/persistence/gorm/table"
)

type Output struct {
	ID             int
	ExternalUserID string
	Name           string
	MailAddress    string
	Comments       string
	UpdatedAt      time.Time
}

func NewOutput(user table.User) *Output {
	return &Output{
		ID:             user.ID,
		ExternalUserID: user.ExternalUserID,
		Name:           user.Name,
		MailAddress:    user.MailAddress,
		Comments:       user.Comments,
		UpdatedAt:      user.UpdatedAt,
	}
}

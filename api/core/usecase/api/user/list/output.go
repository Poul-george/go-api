package list

import (
	"time"

	"github.com/Poul-george/go-api/api/infrastructure/data/persistence/gorm/table"
)

type Output struct {
	Users []OutputUser
}

type OutputUser struct {
	ID             int
	ExternalUserID string
	Name           string
	MailAddress    string
	Comments       string
	UpdatedAt      time.Time
}

func NewOutput(users []table.User) *Output {
	outputUsers := make([]OutputUser, len(users))

	for i, u := range users {
		outputUsers[i] = OutputUser{
			ID:             u.ID,
			ExternalUserID: u.ExternalUserID,
			Name:           u.Name,
			MailAddress:    u.MailAddress,
			Comments:       u.Comments,
			UpdatedAt:      u.UpdatedAt,
		}
	}

	return &Output{Users: outputUsers}
}

package list

import (
	"time"

	"github.com/Poul-george/go-api/api/infrastructure/repository/user"
)

type Response []usr

type usr struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Password    string    `json:"password"`
	MailAddress string    `json:"mail_address"`
	Comments    string    `json:"comments"`
	UpdatedAt   time.Time `json:"latest_day"`
}

func NewResponse(out []user.Users) Response {
	users := make([]usr, len(out))
	for i, u := range out {
		users[i] = usr{
			ID:          u.Id,
			Name:        u.Name,
			Password:    u.Password,
			MailAddress: u.MailAddress,
			Comments:    u.Comments,
			UpdatedAt:   u.UpdatedAt,
		}
	}

	return users
}

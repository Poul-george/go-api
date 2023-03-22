package list

import (
	"time"

	"github.com/Poul-george/go-api/api/core/usecase/api/user/list"
)

type Response []UserResponse

type UserResponse struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Password    string    `json:"password"`
	MailAddress string    `json:"mail_address"`
	Comments    string    `json:"comments"`
	UpdatedAt   time.Time `json:"latest_day"`
}

func NewResponse(out list.Output) Response {
	users := make([]UserResponse, len(out.Users))
	for i, u := range out.Users {
		users[i] = UserResponse{
			ID:          u.Id,
			Name:        u.Name,
			MailAddress: u.MailAddress,
			Comments:    u.Comments,
			UpdatedAt:   u.UpdatedAt,
		}
	}

	return users
}

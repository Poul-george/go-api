package table

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID             int
	ExternalUserID string
	Name           string
	Password       string
	MailAddress    string
	Comments       string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt
}

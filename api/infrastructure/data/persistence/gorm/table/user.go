package table

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id          int            `json:"Id"`
	Name        string         `json:"Name"`
	Password    string         `json:"Password"`
	MailAddress string         `json:"MailAddress"`
	Comments    string         `json:"Comments"`
	CreatedAt   time.Time      `json:"CreatedAt"`
	UpdatedAt   time.Time      `json:"UpdatedAt"`
	DeletedAt   gorm.DeletedAt `json:"DeletedAt"`
}

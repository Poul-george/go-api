package user

import (
	"time"

	"github.com/Poul-george/go-api/api/core/domain/model"
	"github.com/Poul-george/go-api/api/infrastructure/mysql/mysql_setting"
)

type Users struct {
	Id          int       `json:"Id"`
	Name        string    `json:"Name"`
	Password    string    `json:"Password"`
	MailAddress string    `json:"MailAddress"`
	Comments    string    `json:"Comments"`
	CreatedAt   time.Time `json:"CreatedAt"`
	UpdatedAt   time.Time `json:"UpdatedAt"`
}

func Create(user *model.User) error {
	db := mysql_setting.Mysql_Connect()
	err := db.DB.Create(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func List() ([]Users, error) {
	users := []Users{}
	db := mysql_setting.Mysql_Connect()
	err := db.DB.Table("users").Select("id,name,password,mail_address,comments,created_at,updated_at").Scan(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

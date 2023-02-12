package user

import (
	"github.com/Poul-george/go-api/api/core/domain/model"
	"github.com/Poul-george/go-api/api/infrastructure/data/persistence/gorm/table"
	"github.com/Poul-george/go-api/api/infrastructure/mysql/mysql_setting"
)

func Create(user *model.User) error {
	db := mysql_setting.Mysql_Connect()
	err := db.DB.Create(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func List() ([]table.User, error) {
	users := []table.User{}
	db := mysql_setting.Mysql_Connect()
	err := db.DB.Find(&users).Order("id desc").Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

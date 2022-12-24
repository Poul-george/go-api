package request_get

import (
	"encoding/json"
	"net/http"

	"github.com/Poul-george/go-api/api/infrastructure/repository"
	"github.com/Poul-george/go-api/api/mysql_setting"
	"github.com/labstack/echo/v4"
)

func GetUsers(db *mysql_setting.SQLHandler) []repository.Users {
	result := []repository.Users{}
	_ = db.DB.Table("users").Select("id,name,password,mail_address,comments,created_at,updated_at").Scan(&result).Error
	return result
}

func ResponesUsers(c echo.Context) error {
	result := GetUsers(mysql_setting.Mysql_Connect())
	res, _ := json.Marshal(result)

	return c.String(http.StatusOK, string(res))
}

package request_get

import (
	"net/http"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/Poul-george/go-api/api/mysql_set"
)

func GetUsers(c echo.Context) error {
	db := mysql_set.Mysql_Connect()
	result := []*mysql_set.Users{}
	_ = db.DB.Table("users").Select("id,name,author,created_at").Scan(&result).Error

	all_users := make([]mysql_set.Users, len(result))

	for i, r := range result {
		all_users[i].Id = r.Id
		all_users[i].Name = r.Name
		all_users[i].Author = r.Author
		all_users[i].CreatedAt = r.CreatedAt
	}

	res, _ := json.Marshal(all_users)
	return c.String(http.StatusOK, string(res))
}
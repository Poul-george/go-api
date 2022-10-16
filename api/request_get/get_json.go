package request_get

import (
	"net/http"
	"encoding/json"

	"github.com/labstack/echo/v4"
	"time"
)

type Book struct {
	Id string
	Name string
	Author string
	Day time.Time
}

 func ArticleIndex(c echo.Context) error {
	return c.String(http.StatusOK, "hello css html")
}

func GetJson(c echo.Context) error {
	array := make([]Book, 3)
	for i := 0;i < len(array);i++ {
		array[i].Id = "vol1010tttx7ym9l"
		array[i].Name = "C++ programming language"
		array[i].Author = "Bjarne Stroutsrup"
		array[i].Day = time.Now()
	}

	res, _ := json.Marshal(array)
	return c.String(http.StatusOK, string(res))
}
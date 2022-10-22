package mysql_set

import (
	"gorm.io/gorm"
  	"gorm.io/driver/mysql"
	  "fmt"
	  "os"
	  "time"
)

func TestDBSet() (database *gorm.DB) {
	user := os.Getenv("MYSQL_USER")
    password := os.Getenv("MYSQL_PASSWORD")
    host := os.Getenv("MYSQL_HOST")
    port := os.Getenv("MYSQL_PORT")
    dbName := os.Getenv("MYSQL_DATABASE")
    dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?parseTime=true&loc=Asia%2FTokyo"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	count := 0
	fmt.Printf("%v\n",err)
	if err != nil {
		for {
			if err == nil {
				fmt.Println("")
				break
			}
			fmt.Print(".")
			time.Sleep(time.Second)
			count++
			if count > 180 {
				fmt.Println("")
				fmt.Println("DB接続失敗")
				panic(err)
			}
			// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
			if count > 5 {break}
		}
	}
	fmt.Println("DB接続成功")

	return db
}
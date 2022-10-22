package mysql_set

import (
	"gorm.io/gorm"
  	"gorm.io/driver/mysql"
	  "fmt"
	  "os"
	  "time"
)

type Users struct {
    Id int `json:"Id"`
    Name string `json:"Name"`
    Author string `json:"Author"`
    CreatedAt time.Time `json:"CreatedAt"`
}

// SQLHandler ...
type SQLHandler struct {
    DB  *gorm.DB
    Err error
}

var dbConn *SQLHandler
var Db *gorm.DB

// DBOpen は DB connectionを張る。
func DBOpen() {
    dbConn = NewSQLHandler()
}

// DBClose は DB connectionを張る。
func DBClose() {
    sqlDB, _ := dbConn.DB.DB()
    sqlDB.Close()
}

func GetAllItems(users *Users)  {
    Db.Find(&users)
}

// NewSQLHandler ...
func NewSQLHandler() *SQLHandler {
    user := os.Getenv("MYSQL_USER")
    password := os.Getenv("MYSQL_PASSWORD")
    host := os.Getenv("MYSQL_HOST")
    port := os.Getenv("MYSQL_PORT")
    dbName := os.Getenv("MYSQL_DATABASE")
    dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?parseTime=true&loc=Asia%2FTokyo"

	
    var db *gorm.DB
    var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
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
			db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
			if count > 5 {break}
		}
	}
	fmt.Println("DB接続成功")

    sqlDB, _ := db.DB()
    //コネクションプールの最大接続数を設定。
    sqlDB.SetMaxIdleConns(100)
    //接続の最大数を設定。 nに0以下の値を設定で、接続数は無制限。
    sqlDB.SetMaxOpenConns(100)
    //接続の再利用が可能な時間を設定。dに0以下の値を設定で、ずっと再利用可能。
    sqlDB.SetConnMaxLifetime(100 * time.Second)

    mysqlHandler := new(SQLHandler)
    // db.Logger.LogMode(4)
    mysqlHandler.DB = db

    return mysqlHandler
}

// GetDBConn ...
func GetDBConn() *SQLHandler {
    return dbConn
}

// BeginTransaction ...
func BeginTransaction() *gorm.DB {
    dbConn.DB = dbConn.DB.Begin()
    return dbConn.DB
}

// Rollback ...
func RollBack() {
    dbConn.DB.Rollback()
}
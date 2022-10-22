package mysql_set

import (

)

func Mysql_Connect() *SQLHandler {
	DBOpen()
	DB := GetDBConn()
	return DB
}
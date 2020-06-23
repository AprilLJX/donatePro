package db

import (
"database/sql"
_ "github.com/go-sql-driver/mysql"
"log"
)

var DB *sql.DB
func init() {
	var err error
	DB,err = sql.Open("mysql","root:root@tcp(localhost:3306)/donate_google?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
		log.Panic(err)
	}
	DB.SetMaxOpenConns(30)
	DB.SetMaxIdleConns(10)
	err = DB.Ping()
	if err != nil{
		log.Panic(err)
	}
}

package dao

import (
	"Donate_gin/db"
)

func GetPswDao(account string) (password string,err error) {
	err = db.DB.QueryRow("SELECT password FROM donor WHERE account = ?",account).Scan(&password)
	return
}

func GetUser()  {

}


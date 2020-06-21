package dao

import (
	"Donate_gin/db"
	"Donate_gin/entity"
)

//type UserLogin struct {
//	Account string `json:"account" form:"account"`
//	Password string `json:"password" form:"password"`
//}



func  GetRow(account string,password string) (id int, err error)  {
	u := entity.UserLogin{
		Account:account,
		Password:password,
	}
	err = db.DB.QueryRow("SELECT id FROM user_login WHERE account = ? and password=?",u.Account,u.Password).Scan(&id)
	return
}



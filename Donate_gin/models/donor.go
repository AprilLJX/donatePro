package models

import (
	"Donate_gin/dao"
)

func Login(account string,password string) (id int,err error){

	id, err = dao.GetRow(account,password)
	return
}



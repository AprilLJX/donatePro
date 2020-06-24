package models

import (
	"Donate_gin/dao"
	"Donate_gin/entity"
)

//接收api传来的数据，进行一些逻辑处理
func AdLoginModel(account string,password string)(id int,err error)  {
	admin := entity.Administrator{
		Account:  account,
		Password: password,
	}

	admin, err = dao.AdLoginDao(admin)
	id = admin.Id
	return
}

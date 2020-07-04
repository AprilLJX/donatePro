package dao

import (
	"Donate_gin/db"
	"Donate_gin/entity"
)


//数据库处理相关
func AdLoginDao(a1 entity.Administrator) (a2 entity.Administrator,err error) {
	a2 = a1
	err = db.DB.QueryRow("SELECT id FROM administrator WHERE account = ? and password=?",a1.Account,a1.Password).Scan(&a2.Id)
	return

}

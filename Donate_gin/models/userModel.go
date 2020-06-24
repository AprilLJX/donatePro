package models

import (
	"Donate_gin/dao"
	"errors"
)

func UserLoginModel(account string,password string)(err error)  {
	password_, err := dao.GetPswDao(account)
	//todo 加密模块
	if password == password_{



	}else {
		return errors.New("账号或密码错误")
	}
}

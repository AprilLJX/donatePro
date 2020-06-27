package models

import (
	"Donate_gin/dao"
	"errors"
)

//捐赠方登录
func DonorLoginModel(account string,password string)(donormap map[string]string,err error)  {
	donorID ,password_, err := dao.GetDonorPswDao(account)
	//todo 加密模块
	if err != nil{
		return
	}

	donormap = make(map[string]string)
	if password == password_{
		donor,_ := dao.GetDonorDao(donorID)

		donormap["account"] = donor.Account
		donormap["nickname"] = donor.Nickname
		donormap["id_number"] = donor.IdNumber
		donormap["cur_residence"] = donor.CurResidence
		donormap["city"] = donor.City
		donormap["avatar"] = donor.Avatar
		donormap["love_value"] = donor.LoveValue
		donormap["profile"] = donor.Profile
		return

	}else {
		return nil,errors.New("账号或密码错误")
	}

}

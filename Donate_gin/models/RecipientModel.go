package models

import (
	"Donate_gin/dao"
	"crypto/md5"
	"encoding/hex"
	"errors"
)

//捐赠方登录
func RecipientLoginModel(account string,password string)(donormap map[string]interface{},err error)  {
	donorID ,password_, err := dao.GetDonorPswDao(account)

	ctx := md5.New()
	ctx.Write([]byte(password))
	password =  hex.EncodeToString(ctx.Sum(nil))

	if err != nil{
		return
	}

	donormap = make(map[string]interface{})
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

func RecipientRegisterModel(account string,password string,name string,idNumber string,company string,categpry string,creditCode string,address string,profile string) (recipientId int64,err error) {

	//todo 判断公司
	ctx := md5.New()
	ctx.Write([]byte(password))
	password =  hex.EncodeToString(ctx.Sum(nil))


	recipientId,err = dao.RecipientRegisterDao(account,password,name,idNumber,company,categpry,creditCode,address,profile)
	return

}

func RecipientVerify(account string,code string) (err error) {
	gencode := Code[account]

	if code == gencode{

		return nil
	}else {
		return errors.New("验证码错误")
	}

}

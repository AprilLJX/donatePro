package models

import (
	"Donate_gin/dao"
	"crypto/md5"
	"encoding/hex"
	"errors"
)

//捐赠方登录
func RecipientLoginModel(account string,password string)(recipientId int,err error)  {
	recipientId ,password_, err := dao.GetRecipientPswDao(account)

	ctx := md5.New()
	ctx.Write([]byte(password))
	password =  hex.EncodeToString(ctx.Sum(nil))

	if err != nil{
		return
	}

	if password == password_{
		return recipientId,err

	}else {
		return 0,errors.New("账号或密码错误")
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

package models

import (
	"Donate_gin/dao"
	"crypto/md5"
	"encoding/hex"
	"errors"
)

//捐赠方登录
func DonorLoginModel(account string,password string)(donorID int,err error)  {
	donorID ,password_, err := dao.GetDonorPswDao(account)

	ctx := md5.New()
	ctx.Write([]byte(password))
	password =  hex.EncodeToString(ctx.Sum(nil))

	if err != nil{
		return
	}

	if password == password_{
		return donorID,err

	}else {
		return 0,errors.New("账号或密码错误")
	}

}



//发送验证码
//获取验证码
//判断验证码
func DonorRegister(account string,code string,pwd string,name string,IdNumber string,city string) (donationID int64,err error) {
	gencode := Code[account]

	if code == gencode{
		ctx := md5.New()
		ctx.Write([]byte(pwd))
		password :=  hex.EncodeToString(ctx.Sum(nil))
		donationID ,err = dao.AddDonorDao(account,password,name,IdNumber,city)
		if err!= nil{
			return 0,err
		}
		return
	}else {
		return 0,errors.New("验证码错误")
	}

}

package models

import (
	"Donate_gin/dao"
	"Donate_gin/util"
	"errors"
	"fmt"
)

var Code map[string]string

func SendDonorSMSModel(account string) (err error)  {
	donorID,_ := dao.GetDonorAccountDao(account)
	if donorID == 0{
		code := util.GenValidateCode(6)
		fmt.Printf("验证码为:%v\n:",code)
		Code[account] = code
		return nil
	}else {
		return errors.New("账号已存在")
	}
}

func SendRecipientSMSModel(account string) (err error)  {
	donorID,_ := dao.GetRecipientAccountDao(account)
	if donorID == 0{
		code := util.GenValidateCode(6)
		fmt.Printf("验证码为:%v\n:",code)
		Code[account] = code
		return nil
	}else {
		return errors.New("账号已存在")
	}
}

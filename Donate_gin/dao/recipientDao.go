package dao

import (
	"Donate_gin/db"
	"Donate_gin/entity"
	"errors"
	"fmt"
)

//获取接收方用户信息
func GetRecipientDao(recipientID int) (recipient entity.Recipient,err error) {
	recipient.RecipientId = recipientID

	err = db.DB.QueryRow("SELECT account,name,id_number,company,credit_code,com_address,com_profile,recipient_num " +
		"FROM recipient WHERE recipient_id = ?",recipientID).Scan(
		&recipient.Account,&recipient.Name,&recipient.IdNumber,&recipient.Company,&recipient.CreditCode,&recipient.ComAddress,&recipient.ComProfile,&recipient.RecipientNum)

	return
}

func GetRecipientAccountDao(account string) (recipientId int,err error) {
	err = db.DB.QueryRow("SELECT recipient_id FROM recipient WHERE account = ?",account).Scan(&recipientId)
	return
}


func RecipientRegisterDao(account string,password string,name string,idNumber string,company string,categpry string,creditCode string,address string,profile string) (recipientId int64,err error) {
	sqlStr := "insert into recipient(account,password, name,id_number,company,com_category,credit_code,com_address,com_profile) " +
		"values (?,?,?,?,?,?,?,?,?)"

	ret, err := db.DB.Exec(sqlStr, account, password,name, idNumber,company,categpry,creditCode,address,profile)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return 0,errors.New("insert new donor failed")
	}
	recipientId, err = ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return

	}
	return
}

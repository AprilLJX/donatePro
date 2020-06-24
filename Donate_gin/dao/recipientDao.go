package dao

import (
	"Donate_gin/db"
	"Donate_gin/entity"
)

func GetRecipientDao(recipientID int) (recipient entity.Recipient,err error) {
	recipient.RecipientId = recipientID

	err = db.DB.QueryRow("SELECT account,name,id_number,company,credit_code,com_address,com_profile,recipient_num " +
		"FROM recipient WHERE recipient_id = ?",recipientID).Scan(
		&recipient.Account,&recipient.Name,&recipient.IdNumber,&recipient.Company,&recipient.CreditCode,&recipient.ComAddress,&recipient.ComProfile,&recipient.RecipientNum)

	return
}


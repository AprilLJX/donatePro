package models

import "Donate_gin/dao"

func DonateListModel(donorID int,recipientID int)(donorMap map[string]string, recipientMap map[string]string) {
	donorMap = make(map[string]string)
	recipientMap = make(map[string]string)

	donor, _ := dao.GetDonorDao(donorID)
	recipient,_ := dao.GetRecipientDao(recipientID)

	recipientMap["company"] = recipient.Company
	recipientMap["name"] = recipient.Name
	recipientMap["phone"] = recipient.Account
	recipientMap["com_address"] = recipient.ComAddress

	donorMap["name"] = donor.Name
	donorMap["phone"] = donor.Account
	donorMap["id_number"] = donor.IdNumber

	return
}

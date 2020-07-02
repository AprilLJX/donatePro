package models

import (
	"Donate_gin/dao"
	"Donate_gin/entity"
	"strconv"
)

func DonateItemModel(donorID int,recipientID int)(donorMap map[string]string, recipientMap map[string]string) {
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

func AddTargetDonaModel(projectId string,donorId string,materials string,message string,cate string,ifAnony string)(donation entity.TargetDonation,err error)  {
	//todo 随机生成捐赠单号？
	//materialsList []map[string]string
	//materials := ""
	//for _, materialsMap := range materialsList {
	//	for matiral := range materialsMap {
	//		//matiral：类别    materialsMap [matirals]：数量
	//		materials = materials + matiral + ":" +materialsMap[matiral] + ";"
	//	}
	//}
	//fmt.Printf("物资为%v:",materials)

	//将id都转为int型
	projectID, _ := strconv.Atoi(projectId)
	donorID,_ := strconv.Atoi(donorId)
	category,_ := strconv.Atoi(cate)
	ifAnonymous,_ := strconv.Atoi(ifAnony)

	donationID ,err := dao.AddTargetDonaDao(projectID,donorID,materials,message,category,ifAnonymous)

	dao.NewLogisticsDao(int(donationID))

	donation = dao.OneDonationDao(int(donationID))

	return

}
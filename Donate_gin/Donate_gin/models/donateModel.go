package models

import (
	"Donate_gin/dao"
	"Donate_gin/entity"
	"strconv"
	"strings"
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

func AddTargetDonaModel(projectId string,donorId string,materials string,message string,cate string,ifAnony string)(donation entity.TargetDonation,materiallist []map[string]interface{},err error)  {
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

	materialStr := donation.DonateMaterials
	malist := strings.Split(materialStr , ";")
	for index,ma := range malist{
		if index < len(malist)-1{
			material := make(map[string]interface{})
			one := strings.Split(ma , ":")
			material["type"] = one[0]
			material["num"] = one[1]
			materiallist = append(materiallist, material)
		}


	}


	return

}
package models

import (
	"Donate_gin/dao"
	"Donate_gin/entity"
	"fmt"
)

//接收api传来的数据，进行一些逻辑处理
func GetProDetailsModel(proId int) (prolistPlus []map[string]string,err error) {
	//receive := entity.DonaProject{
	//	ProId: proId,
	//	DemandId: demandId,
	//
	//}
	//
	//receive, err = dao.ProDetails(receive)
	//ifend = receive.IfEnd
	//return

	//proID, _ := strconv.Atoi(proId)//string转int
	//onePro := entity.DonaProject{}
	oneProPlus := entity.DemandList{}


	projectDetails,err := dao.GetProDetailsDao(proId)
	print("Model project details: ")
	fmt.Println(projectDetails)
	for _,project := range projectDetails{
		print("Model project:")
		fmt.Println(project)
		demandId := project.DemandId

		oneProPlus,err = dao.GetOneProDetailsDao(demandId)
		print("Another model:")
		fmt.Println(oneProPlus)
		//onePro.RecDonationNum = project.RecDonationNum

		//将返回的单个项目列表用字典形式存储
		//oneProMap := make(map[string]int)
		//oneProMap["proId"] = onePro.ProId
		//oneProMap["demandId"] = onePro.DemandId
		//oneProMap["RecDonationNum"] = onePro.RecDonationNum
		oneProPlusMap := make(map[string]string)
		oneProPlusMap["materials"] = oneProPlus.Materials
		oneProPlusMap["proName"] = oneProPlus.ProName
		oneProPlusMap["RecAddress"] = oneProPlus.RecAddress
		oneProPlusMap["intro"] = oneProPlus.Introduction

		//prolist = append(prolist,oneProMap)
		prolistPlus = append(prolistPlus,oneProPlusMap)
	}

	return

}
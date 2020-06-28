package models

import (
	"Donate_gin/dao"
	"Donate_gin/entity"
)

//获取项目列表

func GetProList() (prolist []map[string]interface{},proNum int,err error){

	onePro := entity.RePro{}


	projectList,err := dao.GetProListDao()
	for _,project := range projectList{
		demandId := project.DemandId
		onePro,err = dao.GetOneProDao(demandId)
		onePro.RecDonationNum = project.RecDonationNum
		onePro.ProID = project.ProId


		//将返回的单个项目列表用字典形式存储
		oneProMap := make(map[string]interface{})
		oneProMap["proID"] = onePro.ProID
		oneProMap["proName"] = onePro.ProName
		oneProMap["proIntro"] = onePro.Introduction

		prolist = append(prolist,oneProMap)
	}
	proNum = len(prolist)
	return

}

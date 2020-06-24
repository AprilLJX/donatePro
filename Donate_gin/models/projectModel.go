package models

import (
	"Donate_gin/dao"
	"Donate_gin/entity"
)

func GetProList() (prolist []map[string]string,proNum int,err error){

	onePro := entity.RePro{}


	projectList,err := dao.GetProListDao()
	for _,project := range projectList{
		demandId := project.DemandId
		onePro,err = dao.GetOneProDao(demandId)
		onePro.RecDonationNum = project.RecDonationNum

		//将返回的单个项目列表用字典形式存储
		oneProMap := make(map[string]string)
		oneProMap["proName"] = onePro.ProName
		oneProMap["proIntro"] = onePro.Introduction

		prolist = append(prolist,oneProMap)
	}
	proNum = len(prolist)
	return

}

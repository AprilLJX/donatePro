package models

import (
	"Donate_gin/dao"
	"Donate_gin/entity"
	"fmt"
	"strconv"
	//"strings"

	//"strings"
)

//接收api传来的数据，进行一些逻辑处理
//func GetProDetailsModel(proId int) (prolistPlus []map[string]string,err error) {
//	oneProPlus := entity.DemandList{}
//
//	projectDetails,err := dao.GetProDetailsDao(proId)

//for _,project := range projectDetails{
//
//	demandId := project.DemandId
//
//	oneProPlus,err = dao.GetOneProDetailsDao(demandId)
//	print("Another model:")
//	fmt.Println(oneProPlus)
//	oneProPlusMap := make(map[string]string)
//	oneProPlusMap["materials"] = oneProPlus.Materials
//	oneProPlusMap["proName"] = oneProPlus.ProName
//	oneProPlusMap["RecAddress"] = oneProPlus.RecAddress
//	oneProPlusMap["intro"] = oneProPlus.Introduction
//
//	//prolist = append(prolist,oneProMap)
//	prolistPlus = append(prolistPlus,oneProPlusMap)
//}

//	return
//
//}

//func GetDonorInfo(proId int) (donorId int,err error) {
//	donorInfo := entity.TargetDonation{}
//	projectDetails,err := dao.GetDonationIdDao(proId)
//
//	for _,project := range projectDetails{
//		donationId := project.DonationId
//		fmt.Println(projectDetails)
//
//		donorInfo,err = dao.GetDonorIdDao(donationId)
//		donorId := donorInfo.DonorId
//		fmt.Println(donorId)
//		//donorInfoMap := make(map[string]int)
//		//donorInfoMap["donorId"] = donorInfo.DonorId
//
//		//donorinfo = append(donorinfo,donorInfoMap)
//	}
//	return
//
//}

func GetDonorInfo(proId int) (donorId int,err error) {
	//donorInfo := entity.TargetDonation{}
	proInfo, err := dao.GetDonationIdDao(proId)
	donationId := proInfo.DonationId
	donorIdInfo, err := dao.GetDonorIdDao(donationId)
	donorId = donorIdInfo.DonorId
	donorInfo, err := dao.GetDonorDao(donorId)
	donorCity := donorInfo.City
	fmt.Println(donorCity)

	//for _,project := range projectDetails{
	//	donationId := project.DonationId
	//	fmt.Println(projectDetails)
	//
	//	donorInfo,err = dao.GetDonorIdDao(donationId)
	//	donorId := donorInfo.DonorId
	//	fmt.Println(donorId)
	//	//donorInfoMap := make(map[string]int)
	//	//donorInfoMap["donorId"] = donorInfo.DonorId
	//
	//	//donorinfo = append(donorinfo,donorInfoMap)
	//}
	return

}

func GetHistoryDonationModel(donorId int) (prolist []map[string]string,err error){

	onePro := entity.ProDonation{}
	oneProPlus := entity.DonaProject{}
	oneProMore:= entity.DemandList{}



	projectList,err := dao.GetHistoryDonationDao(donorId)

	for _,project := range projectList{
		donationId := project.TargetId
		ifAudit := project.IfAudit
		IfAudit := strconv.Itoa(ifAudit)
		DonationID := strconv.Itoa(donationId)
		onePro,err = dao.GetHistoryProDao(donationId)
		proId := onePro.ProjectId
		//ProId := strconv.Itoa(proId)
		oneProPlus, err = dao.GetHistoryRecDao(proId)
		demandId := oneProPlus.DemandId
		oneProMore,err = dao.GetOneProDetailsDao(demandId)
		recDonationNum := oneProPlus.RecDonationNum
		RecDonationNum := strconv.Itoa(recDonationNum)

		//materialStr := oneProMore.Materials
		//malist := strings.Split(materialStr , ";")
		//materiallist := make(map[string]interface{});
		//
		//for index,ma := range malist{
		//	if index < len(malist)-1{
		//		material := make(map[string]interface{});
		//		one := strings.Split(ma , ":");
		//		material["type"] = one[0];
		//		material["num"] = one[1];
		//		materiallist = append(materiallist, material);
		//	}
		//
		//
		//}

		//将返回的单个项目列表用字典形式存储
		oneProMap := make(map[string]string)
		//oneProMap["proId"] = ProId
		oneProMap["donationId"] = DonationID
		oneProMap["proName"] = oneProMore.ProName
		oneProMap["ifAudit"] = IfAudit
		//oneProMap["proIntro"] = oneProMore.Introduction
		oneProMap["materials"] = oneProMore.Materials
		oneProMap["rec_donation_num"] = RecDonationNum



		prolist = append(prolist,oneProMap)

	}
	return

}

func GetRecipientInfoModel(recipientId int) (prolistPlus []map[string]string,err error) {
	oneProPlus := entity.DemandList{}

	projectDetails,err := dao.GetRecipientProsDao(recipientId)
	for _,project := range projectDetails{
		demandId := project.DemandID

		oneProPlus,err = dao.GetOneProDetailsDao(demandId)
		oneProPlusMap := make(map[string]string)
		DemandID := strconv.Itoa(demandId)

		oneProPlusMap["demandId"] = DemandID
		oneProPlusMap["materials"] = oneProPlus.Materials
		oneProPlusMap["proName"] = oneProPlus.ProName
		oneProPlusMap["RecAddress"] = oneProPlus.RecAddress
		oneProPlusMap["intro"] = oneProPlus.Introduction

		//prolist = append(prolist,oneProMap)
		prolistPlus = append(prolistPlus,oneProPlusMap)
	}

	return

}


func GetDetails(proID int)(listPlus []map[string]string, participants int,  err error) {
	donorInfo := entity.Donor{}
	participants = 0;
	donorIdInfo, err := dao.GetDonorForProDao(proID)
	for _, donor := range donorIdInfo {

		donationId := donor.DonationId
		donorAnotherInfo, err := dao.GetDonorIdDao(donationId)
		donorId := donorAnotherInfo.DonorId
		donorInfo, err = dao.GetDonorInfoDao(donorId)
		donorInfoo, err := dao.GetDonorIdforProDao(donorId)

		donorIdMap := make(map[string]string)
		//loveValue := donorInfo.LoveValue
		//donorIdMap["loveValue"] = loveValue
		donorIdMap["name"] = donorInfo.Name
		//donorIdMap["curResidence"] = donorInfo.CurResidence
		donorIdMap["donateMaterials"] = donorInfoo.DonateMaterials
		donorIdMap["massage"] = donorInfoo.Message
		donorIdMap["donorTime"] = donorInfoo.DonateTime
		donorMoreInfo, err := dao.GetHistoryDonationDao(donorId)
		fmt.Println(donorMoreInfo)
		fmt.Println(err)
		listPlus = append(listPlus, donorIdMap)
		participants += 1

	}
	return
}
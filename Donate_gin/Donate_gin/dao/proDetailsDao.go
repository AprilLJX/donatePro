package dao

import (
	"Donate_gin/db"
	"Donate_gin/entity"
	"fmt"
	"log"
	"strconv"

	//"log"
)


////数据库处理相关
//func ProDetails(a1 entity.DonaProject) (a2 entity.DonaProject,err error) {
//	a2 = a1
//	err = db.DB.QueryRow("SELECT project FROM DonateProject WHERE proID = ? and demandID=?",a1.ProId,a1.DemandId).Scan(&a2.IfEnd)
//	return
//
//}

func GetProDetailsDao(proId int)  (projects entity.DonaProject, err error){


	err = db.DB.QueryRow("SELECT pro_id, demand_id, rec_donation_num, if_end FROM dona_project WHERE pro_id=?",proId).Scan(&projects.ProId, &projects.DemandId, &projects.RecDonationNum, &projects.IfEnd);


	//for rows.Next() {
	//	if err1 := rows.Scan(&projects.ProId, &projects.DemandId); err1 != nil {
	//		log.Fatal(err1)
	//	}
	//
	//	projectList = append(projectList,projects)
	//
	//}
	return
}
func GetDemandIdDao(proId int)(oneProPlus entity.DonaProject,err error)  {
	err = db.DB.QueryRow("SELECT demand_id, pro_id, rec_donation_num FROM dona_project WHERE pro_id=?",proId).Scan(&oneProPlus.DemandId,&oneProPlus.ProId,&oneProPlus.RecDonationNum)
	return
}

func GetOneProDetailsDao(demandId int)(oneProPlus entity.DemandList,err error)  {
	err = db.DB.QueryRow("SELECT emergency_degree, recipient_id, demand_id, materials, rec_address, introduction, pro_name, category FROM demand_list WHERE demand_id=?",demandId).Scan( &oneProPlus.EmergencyDegree,&oneProPlus.RecipientId, &oneProPlus.DemandID, &oneProPlus.Materials,&oneProPlus.RecAddress,&oneProPlus.Introduction, &oneProPlus.ProName, &oneProPlus.Category)
	return
}

func GetCompanyDao(recipientId int)(oneProPlus entity.Recipient,err error)  {
	err = db.DB.QueryRow("SELECT name, company, com_category, com_address, id_number, com_profile, recipient_id, credit_code FROM recipient WHERE recipient_id=?",recipientId).Scan(&oneProPlus.Name,&oneProPlus.Company,&oneProPlus.ComCategory,&oneProPlus.ComAddress, &oneProPlus.IdNumber, &oneProPlus.ComProfile,&oneProPlus.RecipientId, &oneProPlus.CreditCode)
	return
}
//func GetDonationIdDao(proId int)(donation []entity.ProDonation,err error)  {
//	projects := entity.ProDonation{
//	}
//	rows := db.DB.QueryRow("SELECT * FROM project_donation WHERE project_id=?",proId)
//	if err1 := rows.Scan(&projects.DonationId); err1 != nil {
//		log.Fatal(err1)
//	}
//	donation = append(donation,projects)
//
//	fmt.Println(rows)
//	return
//}

func GetDonationIdDao(proId int)(donation entity.ProDonation,err error)  {

	err = db.DB.QueryRow("SELECT id, project_id, donation_id FROM project_donation WHERE project_id=?",proId).Scan(&donation.Id, &donation.ProjectId, &donation.DonationId)
	return
}

func GetDonorForProDao(proId int)(donorList[] entity.ProDonation,err error)  {

	donors := entity.ProDonation{

	}

	rows, err := db.DB.Query("SELECT * FROM project_donation WHERE project_id = ?",proId)
	defer rows.Close()

	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		if err1 := rows.Scan(&donors.Id, &donors.ProjectId, &donors.DonationId); err1 != nil {
			log.Fatal(err1)
		}

		donorList = append(donorList,donors)

	}
	fmt.Println(donorList)
	return
}

//func GetDonorIdDao(targetId int)(oneProPlus []entity.TargetDonation,err error)  {
//	projects := entity.TargetDonation{
//	}
//	rows := db.DB.QueryRow("SELECT * FROM demand_list WHERE target_id=?",targetId)
//	if err1 := rows.Scan(&projects.DonorId); err1 != nil {
//		log.Fatal(err1)
//	}
//	oneProPlus = append(oneProPlus,projects)
//
//	fmt.Println(rows)
//	return
//}

func GetDonorIdDao(targetId int)(target entity.TargetDonation,err error)  {

	err = db.DB.QueryRow("SELECT target_id,donor_id,if_standard,if_audit,category,message,donate_time,donate_materials,match_pro,if_anonymous FROM target_donation WHERE target_id=?",targetId).Scan(&target.TargetId,&target.DonorId,&target.IfStandard,&target.IfAudit,&target.Category,&target.Message,&target.DonateTime,&target.DonateMaterials,&target.MatchPro,&target.IfAnonymous)
	return
}

func GetDonorIdforProDao(donorId int)(target entity.TargetDonation,err error)  {

	err = db.DB.QueryRow("SELECT target_id,donor_id,if_standard,if_audit,category,message,donate_time,donate_materials,match_pro,if_anonymous FROM target_donation WHERE target_id=?",donorId).Scan(&target.TargetId,&target.DonorId,&target.IfStandard,&target.IfAudit,&target.Category,&target.Message,&target.DonateTime,&target.DonateMaterials,&target.MatchPro,&target.IfAnonymous)
	return
}

func GetDonorInfoDao(donorId int) (donor entity.Donor,err error) {

	err = db.DB.QueryRow("SELECT donor_id,account,id_number,love_value,nickname,name,cur_residence,city,avatar,profile FROM donor WHERE donor_id = ?",donorId).Scan(&donor.DonorID,&donor.Account,&donor.IdNumber,&donor.LoveValue,&donor.Nickname,&donor.Name,&donor.CurResidence,&donor.City,&donor.Avatar,&donor.Profile)

	return
}

func GetHistoryDonationDao(donorId int)  (projectList []entity.TargetDonation,err error){

	projects := entity.TargetDonation{

	}

	rows, err := db.DB.Query("SELECT * FROM target_donation WHERE donor_id = ?",donorId)
	defer rows.Close()

	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		if err1 := rows.Scan(&projects.TargetId,&projects.DonorId,&projects.Category, &projects.DonateMaterials, &projects.IfStandard, &projects.IfAudit,&projects.DonateTime,&projects.MatchPro,&projects.IfAnonymous,&projects.Message); err1 != nil {
			log.Fatal(err1)
		}


		projectList = append(projectList,projects)

	}
	fmt.Println(projectList)
	return
}

func GetHistoryDoDao(donorId int)  (donationList []map[string]string,err error){

	projects := entity.TargetDonation{

	}

	rows, err := db.DB.Query("SELECT * FROM target_donation WHERE donor_id = ?",donorId)
	defer rows.Close()

	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		if err1 := rows.Scan(&projects.TargetId,&projects.DonorId,&projects.Category, &projects.DonateMaterials, &projects.IfStandard, &projects.IfAudit,&projects.DonateTime,&projects.MatchPro,&projects.IfAnonymous,&projects.Message); err1 != nil {
			log.Fatal(err1)
		}
		oneProMap := make(map[string]string)
		targetId := projects.TargetId
		TargetId := strconv.Itoa(targetId)
		oneProMap["targetId"] = TargetId


		donationList = append(donationList,oneProMap)

	}
	fmt.Println(donationList)
	return
}

func GetHistoryProDao(donationId int)(donation entity.ProDonation,err error)  {

	err = db.DB.QueryRow("SELECT id, project_id, donation_id FROM project_donation WHERE donation_id=?",donationId).Scan(&donation.Id, &donation.ProjectId, &donation.DonationId)
	return
}

func GetHistoryRecDao(proId int)(donation entity.DonaProject,err error)  {

	err = db.DB.QueryRow("SELECT demand_id, pro_id, rec_donation_num FROM dona_project WHERE pro_id=?",proId).Scan(&donation.DemandId, &donation.ProId, &donation.RecDonationNum)
	return
}


func GetProNameDao(demandId int)(oneProPlus entity.DemandList,err error)  {
	err = db.DB.QueryRow("SELECT materials, rec_address, introduction, pro_name FROM demand_list WHERE demand_id=?",demandId).Scan(&oneProPlus.Materials,&oneProPlus.RecAddress,&oneProPlus.Introduction, &oneProPlus.ProName)
	return
}

func GetRecipientProsDao(recipientId int)(oneProPlus[] entity.DemandList,err error)  {
	projects := entity.DemandList{}
	rows, err := db.DB.Query("SELECT * FROM demand_list WHERE recipient_id=?",recipientId)
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		if err := rows.Scan(&projects.DemandID, &projects.RecipientId, &projects.ProName, &projects.Introduction,&projects.Category, &projects.Materials, &projects.IfStandard, &projects.IfAudit,&projects.RecAddress, &projects.CutOffTime,&projects.EmergencyDegree,&projects.InitiationTime ); err != nil {
			log.Fatal(err)
		}
		oneProPlus = append(oneProPlus, projects)
	}
	return
}

package dao

import (
	"Donate_gin/db"
	"Donate_gin/entity"
	"fmt"
	"log"

	//"log"
)


////数据库处理相关
//func ProDetails(a1 entity.DonaProject) (a2 entity.DonaProject,err error) {
//	a2 = a1
//	err = db.DB.QueryRow("SELECT project FROM DonateProject WHERE proID = ? and demandID=?",a1.ProId,a1.DemandId).Scan(&a2.IfEnd)
//	return
//
//}

func GetProDetailsDao(proId int)  (onePro []entity.DonaProject, err error){

	projects := entity.DonaProject{

	}

	rows := db.DB.QueryRow("SELECT * FROM dona_project WHERE pro_id=?",proId)
	//defer rows.Close()
	if err1 := rows.Scan(&projects.ProId, &projects.DemandId, &projects.RecDonationNum, &projects.IfEnd); err1 != nil {
		log.Fatal(err1)
	}
	onePro = append(onePro,projects)

	fmt.Println(rows)


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

func GetOneProDetailsDao(demandId int)(oneProPlus entity.DemandList,err error)  {
	err = db.DB.QueryRow("SELECT materials, rec_address, introduction, pro_name FROM demand_list WHERE demand_id=?",demandId).Scan(&oneProPlus.Materials,&oneProPlus.RecAddress,&oneProPlus.Introduction, &oneProPlus.ProName)
	return
}

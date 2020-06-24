package dao

import (
	"Donate_gin/db"
	"Donate_gin/entity"
	"fmt"
	"log"
)


////数据库处理相关
//func ProDetails(a1 entity.DonaProject) (a2 entity.DonaProject,err error) {
//	a2 = a1
//	err = db.DB.QueryRow("SELECT project FROM DonateProject WHERE proID = ? and demandID=?",a1.ProId,a1.DemandId).Scan(&a2.IfEnd)
//	return
//
//}

func GetProDetailsDao()  (projectList []entity.DonaProject,err error){

	projects := entity.DonaProject{

	}

	rows, err := db.DB.Query("SELECT * FROM dona_project")
	defer rows.Close()

	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		if err1 := rows.Scan(&projects.ProId, &projects.DemandId, &projects.IfEnd); err1 != nil {
			log.Fatal(err1)
		}

		projectList = append(projectList,projects)

	}
	return
}

func GetOneProDetailsDao(demandId int)(onePro entity.DonaProject,err error)  {
	err = db.DB.QueryRow("SELECT proId FROM demand_list WHERE demand_id=?",demandId).Scan(&onePro.ProId)
	return
}

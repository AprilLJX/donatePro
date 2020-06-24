package dao

import (
	"Donate_gin/db"
	"Donate_gin/entity"
	"fmt"
	"log"
)

func GetProListDao()  (projectList []entity.DonaProject,err error){

	projects := entity.DonaProject{

	}
	//返回if_end为0的项目
	rows, err := db.DB.Query("SELECT * FROM dona_project WHERE if_end = 0")
	defer rows.Close()

	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		if err1 := rows.Scan(&projects.ProId, &projects.DemandId, &projects.RecDonationNum, &projects.IfEnd); err1 != nil {
			log.Fatal(err1)
		}

		projectList = append(projectList,projects)

	}
		return
}

func GetOneProDao(demandId int)(onePro entity.RePro,err error)  {
	err = db.DB.QueryRow("SELECT pro_name,introduction FROM demand_list WHERE demand_id=?",demandId).Scan(&onePro.ProName,&onePro.Introduction)
	return
}

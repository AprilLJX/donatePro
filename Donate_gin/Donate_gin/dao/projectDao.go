package dao

import (
	"Donate_gin/db"
	"Donate_gin/entity"
	"errors"
	"fmt"
	"log"
)

//获取项目列表
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

//获取一个项目信息
func GetOneProDao(demandId int)(onePro entity.RePro,err error)  {
	err = db.DB.QueryRow("SELECT pro_name,introduction FROM demand_list WHERE demand_id=?",demandId).Scan(&onePro.ProName,&onePro.Introduction)
	return
}

func AddProject(demandlist int)(demandlistid int,err error)  {
	sqlStr := "insert into dona_project(demand_id)values (?)"
	ret, err := db.DB.Exec(sqlStr, demandlist)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return 0,errors.New("insert new donor failed")
	}
	demandid, err := ret.LastInsertId() // 新插入数据的id
	demandlistid = int(demandid)
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return

	}
	return


}

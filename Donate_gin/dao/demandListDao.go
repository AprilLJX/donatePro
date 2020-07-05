package dao

import (
	"Donate_gin/db"
	"errors"
	"fmt"
	"time"
)


func AddDemandlistDao(recipientId int,proName string,introduction string,category int,materials string,recAddress string,cutoff_time string, emergencyDegree int) (demandlistId int64,err error) {
	sqlStr := "insert into demand_list(recipient_id,pro_name, introduction,category,materials,if_standard,if_audit,rec_address,cut_off_time,emergency_degree,initiation_time) " +
		"values (?,?,?,?,?,?,?,?,?,?,?)"
	ret, err := db.DB.Exec(sqlStr, recipientId, proName, introduction, category,materials,1,1,recAddress,cutoff_time,emergencyDegree,time.Now())
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return 0,errors.New("insert failed")
	}
	demandlistId, err = ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return

	}
	return
}

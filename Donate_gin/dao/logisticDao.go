package dao

import (
	"Donate_gin/db"
	"fmt"
)

func GetLogisiticsDao(logisticId int,state interface{})  {
	sqlStr := "update logistics set state=? where logistics_ = ?"
	ret, err := db.DB.Exec(sqlStr,state , logisticId)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}

func NewLogisticsDao(donationId int)  {
	sqlStr := "insert into logistics(donation_id) values (?)"
	ret, err := db.DB.Exec(sqlStr,donationId)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)

}

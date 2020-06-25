package dao

import (
	"Donate_gin/db"
	"Donate_gin/entity"
	"errors"
	"fmt"
	"time"
)

//获取用户密码，返回到model层进行判断
func GetDonorPswDao(account string) (donorId int,password string,err error) {
	err = db.DB.QueryRow("SELECT donor_id,password FROM donor WHERE account = ?",account).Scan(&donorId,&password)
	return
}

//获取用户信息
func GetDonorDao(donorID int) (donor entity.Donor,err error) {
	donor.DonorID = donorID

	err = db.DB.QueryRow("SELECT account,nickname,name,id_number,cur_residence,city,avatar,love_value,profile " +
		"FROM donor WHERE donor_id = ?",donorID).Scan(
		&donor.Account,&donor.Nickname,&donor.Name,&donor.IdNumber,&donor.CurResidence,&donor.City,&donor.Avatar,&donor.LoveValue,&donor.Profile)

	return
}

//插入数据库，返回捐赠单号
func AddTargetDonaDao(projectID int,donorID int,materials string,message string,category int,ifAnonymous int) (donationID int64,err error){
	sqlStr := "insert into target_donation(donor_id,category, donate_materials,if_standard,if_audit,donate_time,match_pro,if_anonymous,message) " +
		"values (?,?,?,?,?,?,?,?,?)"
	fmt.Printf("donorID:%v\n",donorID)
	ret, err := db.DB.Exec(sqlStr, donorID, category, materials, 1, 1, time.Now(), projectID, ifAnonymous,message)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return 0,errors.New("insert failed")
	}
	donationID, err = ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return

	}
	fmt.Printf("get lastinsert ID:%v\n", donationID)

	return


}
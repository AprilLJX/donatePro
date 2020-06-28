package dao

import (
	"Donate_gin/db"
	"Donate_gin/entity"
	"errors"
	"fmt"
)

//获取用户密码，返回到model层进行判断
func GetDonorPswDao(account string) (donorId int,password string,err error) {
	err = db.DB.QueryRow("SELECT donor_id,password FROM donor WHERE account = ?",account).Scan(&donorId,&password)
	return
}

//获取用户账号，判断账号是否存在
func GetDonorAccountDao(account string) (donorId int,err error) {
	err = db.DB.QueryRow("SELECT donor_id FROM donor WHERE account = ?",account).Scan(&donorId)
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


func AddDonorDao(account string,password string,name string,IdNumber string,city string) (donorID int64,err error) {
	sqlStr := "insert into donor(account,password, nickname,name,id_number,city,love_value) " +
		"values (?,?,?,?,?,?,?)"

	//todo 前端是否需要加一个nickname
	ret, err := db.DB.Exec(sqlStr, account, password,name, name, IdNumber, city, 100)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return 0,errors.New("insert new donor failed")
	}
	donorID, err = ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return

	}
	return
}
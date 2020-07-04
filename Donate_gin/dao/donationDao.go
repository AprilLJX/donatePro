package dao

import (
	"Donate_gin/db"
	"Donate_gin/entity"
	"errors"
	"fmt"
	"time"
)

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

func OneDonationDao(donationId int) (donation entity.TargetDonation) {
	sqlStr := "select * from target_donation where target_id=?"
	// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	err := db.DB.QueryRow(sqlStr, donationId).Scan(&donation.TargetId,&donation.DonorId,&donation.Category,&donation.DonateMaterials,&donation.IfStandard,&donation.IfAudit,&donation.DonateTime,&donation.MatchPro,&donation.IfAnonymous,&donation.Message,&donation.IfFinish,&donation.Reasons)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	return donation

}

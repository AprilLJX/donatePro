package entity

import "database/sql"

//捐赠方账号
type Donor struct{
	DonorID int
	Account string
	Password string
	Nickname string
	Name string
	IdNumber string //身份证账号
	CurResidence sql.NullString //现居地
	City sql.NullString
	Avatar sql.NullString //头像，之后再改格式
	LoveValue string //爱心值 or信用分
	Profile sql.NullString //个人简介
}
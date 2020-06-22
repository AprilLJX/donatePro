package entity

//捐赠方账号
type Donor struct{
	DonorID int
	Account string
	Password string
	Nickname string
	Name string
	IdNumber string //身份证账号
	CurResidence string //现居地
	City string
	Avatar string //头像，之后再改格式
	LoveValue string //爱心值 or信用分
	profile string //个人简介
}
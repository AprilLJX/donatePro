package entity

//受赠方信息
type Recipient struct {
	RecipientId int
	Account string
	Password string
	Name string
	IdNumber string
	Company string
	ComCategory string
	CreditCode string
	ComAddress string
	ComProfile string //公司简介
	RecipientNum int //受捐次数

}

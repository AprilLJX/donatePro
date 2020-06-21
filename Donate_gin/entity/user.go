package entity

type UserLogin struct {
	Account string `json:"account" form:"account"`
	Password string `json:"password" form:"password"`
}


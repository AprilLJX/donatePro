package models

import (
	"Donate_gin/dao"
	"strconv"
)

func AddDemanlistModel(Id string,proName string,introduction string,cate string,materials string,recAddress string,cutoff_time string,emergency string)(demndlistId int64,err error)  {
	recipientId, _ := strconv.Atoi(Id)
	category, _ :=  strconv.Atoi(cate)
	emergencyDegree,_ := strconv.Atoi(emergency)

	demndlistId,err = dao.AddDemandlistDao(recipientId,proName,introduction,category,materials,recAddress,cutoff_time,emergencyDegree)
	return
}

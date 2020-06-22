package entity

//定向捐赠单
type TargetDonation struct {
	TargetId int
	DonorId int //捐赠方id
	Category int //捐赠物资类型
	DonateMaterials string //捐赠物品详情
	IfStandard int //是否达标
	IfAudit int //是否审计
	DonateTime string //捐赠时间
	MatchPro int //匹配的项目id
	IfAnonymous int //是否匿名
	Message string //爱心寄语
}

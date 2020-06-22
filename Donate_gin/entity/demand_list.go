package entity

//需求单
type DemandList struct{
	DemandID int
	RecipientId int
	ProName string //发起的需求单的项目名称
	Introduction string
	Category int
	Materials string
	IfStandard int //是否达标
	IfAudit int //是否审核
	RecAddress string //收货地址
	CutOffTime string //截止时间
	EmergencyDegree int//紧急程度
	InitiationTime string //发起时间

}
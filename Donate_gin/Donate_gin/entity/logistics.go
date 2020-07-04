package entity

import "database/sql"

//物流单号
//物流状态: 0-无轨迹，1-已揽收，2-在途中，3-签收,4-问题件
type Logistics struct{
	LogisticsId int
	DonationId int
	LogiCom sql.NullString
	LogiNumber sql.NullString
	LogiTime sql.NullString
	status int
}

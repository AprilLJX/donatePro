package api

import (
	"Donate_gin/dao"
	"Donate_gin/models"
	"fmt"
	"strconv"

	//"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ProDetails(c *gin.Context)  {
	proId := c.PostForm("pro_id")//proId是string
	proID, _ := strconv.Atoi(proId)//proID是int
	demandIdInfo, err := dao.GetDemandIdDao(proID)
	demandId := demandIdInfo.DemandId
	recipientIdInfo, err := dao.GetOneProDetailsDao(demandId)
	recipiendId := recipientIdInfo.RecipientId
	recipientInfo, err := dao.GetCompanyDao(recipiendId)
	recipientMap := make(map[string]string)
	recipientMap["company"] = recipientInfo.Company
	recipientMap["com_address"] = recipientInfo.ComAddress
	recipientMap["category"] = recipientInfo.ComCategory
	recipientMap["profile"] = recipientInfo.ComProfile

	prolistPlus,err := models.GetProDetailsModel(proID)
	prolist,err := dao.GetProDetailsDao(proID)
	donorIdInfo,err := dao.GetDonationIdDao(proID)
	donationId := donorIdInfo.DonationId
	donorAnotherInfo,err := dao.GetDonorIdDao(donationId)
	donorId := donorAnotherInfo.DonorId
	donorInfo, err := dao.GetDonorInfoDao(donorId)
	donorIdMap := make(map[string]string)
	loveValue := donorInfo.LoveValue
	donorIdMap["loveValue"] = loveValue
	donorIdMap["name"] = donorInfo.Name
	donorIdMap["curResidence"] = donorInfo.CurResidence

	donorMoreInfo, err := dao.GetHistoryDonationDao(donorId)
	fmt.Println(donorMoreInfo)
	//donorInfo,err := dao.GetDonorDao(donorId)
	if err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"status" :http.StatusInternalServerError,
			"msg":"查询失败",
		})

	}else {
		c.JSON(http.StatusOK,gin.H{
			"status" :http.StatusOK,
			"msg":"项目详情如下",
			"prolist" : prolist,
			"proListPlus" :prolistPlus,
			"donorIdInfo":donorIdInfo,
			"anotherInfo":donorAnotherInfo,
			"donorInfo":donorIdMap,
			"recipientInfo":recipientMap,

			//"donorPersonal":donorMoreInfo,
		})
	}


}



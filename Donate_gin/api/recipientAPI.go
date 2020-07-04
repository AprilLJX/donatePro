package api

import (
	"Donate_gin/dao"
	"Donate_gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func RecipientInfo(c *gin.Context)  {
	recipientId:= c.PostForm("recipient_id")//proId是string
	recipientID, _ := strconv.Atoi(recipientId)//proID是int
	recipientInfo, err := dao.GetCompanyDao(recipientID)
	recipientMap := make(map[string]string)
	recipientMap["creditCode"] = recipientInfo.CreditCode
	recipientMap["recipientId"] = recipientId
	recipientMap["company"] = recipientInfo.Company
	recipientMap["com_address"] = recipientInfo.ComAddress
	recipientMap["category"] = recipientInfo.ComCategory
	recipientMap["profile"] = recipientInfo.ComProfile
	//Info, err := dao.GetRecipientProsDao(recipientID)
	prolist, err := models.GetRecipientInfoModel(recipientID)
	if err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"status" :http.StatusInternalServerError,
			"msg":"查询失败",
		})

	}else {
		c.JSON(http.StatusOK,gin.H{
			"status" :http.StatusOK,
			"msg":"查询成功",
			"recipientInfo": recipientMap,
			"info":recipientInfo,
			"proList" :prolist,
		})
	}


}

package api

import (
	"Donate_gin/models"
	//"Donate_gin/dao"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func DonorLogin(c *gin.Context)  {
	account := c.PostForm("account")
	password := c.PostForm("password")

	donormap,err := models.DonorLoginModel(account,password)
	if err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"msg":err.Error(),
			"status":http.StatusBadRequest,
			"donor":nil,
		})
	}else {
		c.JSON(http.StatusOK,gin.H{
			"status":http.StatusOK,
			"msg":"登录成功",
			"donor":donormap,
		})
	}
}

func PersonalCenter(c *gin.Context)  {
	donorId := c.PostForm("donor_id")
	donorID, _ := strconv.Atoi(donorId)
	//donorInfo, err := dao.GetDonorInfoDao(donorID)
	//donorHistory, err := dao.GetHistoryDonationDao(donorID)


	ProList, err := models.GetHistoryDonationModel(donorID)

	if err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"status" :http.StatusInternalServerError,
			"msg":"查询失败",
		})

	}else {
		c.JSON(http.StatusOK,gin.H{
			"status" :http.StatusOK,
			"msg":"查询成功",
			//"donorInfo": donorInfo,
			//"donorHistory": donorHistory,

			"DonationHistory" :ProList,
		})
	}
}

func SystemMessage(c *gin.Context)  {
	proName := c.PostForm("pro_name")
	date := c.PostForm("date")
	//donorInfo, err := dao.GetDonorInfoDao(donorID)
	//donorHistory, err := dao.GetHistoryDonationDao(donorID)



	c.JSON(http.StatusOK,gin.H{
		"status" :http.StatusOK,
		"proName": proName,
		"date": date,
		//"donorInfo": donorInfo,
		//"donorHistory": donorHistory,
		"message":"项目捐赠意向单已审核通过！请尽快补充物流信息吧！",
	})

}
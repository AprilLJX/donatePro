package api

import (
	"Donate_gin/dao"
	"Donate_gin/models"
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

func DonorRegister(c *gin.Context)  {
	account := c.PostForm("phone")
	code := c.PostForm("code")
	password := c.PostForm("password")
	name := c.PostForm("name")
	IdNumber := c.PostForm("id_number")
	city := c.PostForm("city")

	donationID,err := models.DonorRegister(account,code,password,name,IdNumber,city)
	if err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"msg":err.Error(),
			"code":http.StatusInternalServerError,
			"donationId":nil,
		})
	}else {
		c.JSON(http.StatusOK,gin.H{
			"msg":"注册成功",
			"status":http.StatusOK,
			"donationId":donationID,
		})
	}

}

func PersonalCenter(c *gin.Context)  {
	donorId := c.PostForm("donor_id")
	donorID, _ := strconv.Atoi(donorId)
	donorInfo, err := dao.GetDonorInfoDao(donorID)
	donorHistory, err := dao.GetHistoryDonationDao(donorID)
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
			"donorInfo": donorInfo,
			"donorHistory": donorHistory,
			"proList" :ProList,
		})
	}
}
package api

import (
	"Donate_gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
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
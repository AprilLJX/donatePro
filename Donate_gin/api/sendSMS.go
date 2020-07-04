package api

import (
	"Donate_gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SendSMS(c *gin.Context)  {
	usertype := c.PostForm("type")
	account := c.PostForm("phone")
	var err error
	if usertype == "0"{
		err = models.SendDonorSMSModel(account)

	}else {
		err = models.SendRecipientSMSModel(account)
	}
	if err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"status":http.StatusInternalServerError,
			"msg":err.Error(),
		})
	}else {
		c.JSON(http.StatusOK,gin.H{
			"status":http.StatusOK,
			"msg":"返回验证码成功",
		})
	}
}

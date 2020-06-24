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
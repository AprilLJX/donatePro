package api

import (
	"Donate_gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
)
//路由函数，实现与前端交互
func AdLogin(c *gin.Context)  {
	account := c.PostForm("account")
	password := c.PostForm("password")

	adminId,err := models.AdLoginModel(account,password)
	if err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"msg":err.Error(),
		})
	}else {
		c.JSON(http.StatusOK,gin.H{
			"status":http.StatusOK,
			"msg":"登录成功",
			"adminid":adminId,
		})
	}
}

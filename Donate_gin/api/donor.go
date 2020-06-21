package api

import (
	"Donate_gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context)  {

	account := c.PostForm("account")
	password := c.PostForm("password")

	rs,err := models.Login(account,password)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"err":err.Error(),
		})
	}else {
		c.JSON(http.StatusOK, gin.H{
			"result": rs,
		})
	}

}

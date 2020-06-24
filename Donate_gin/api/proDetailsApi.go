package api

import (
	"Donate_gin/models"
	//"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ProDetails(c *gin.Context)  {
	prolist,prolistPlus,err := models.GetProDetails()
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
		})
	}


}



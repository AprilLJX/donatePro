package api

import (
	"Donate_gin/dao"
	"Donate_gin/models"
	"strconv"

	//"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ProDetails(c *gin.Context)  {
	proId := c.PostForm("pro_id")//proId是string
	proID, _ := strconv.Atoi(proId)//proID是int
	prolistPlus,err := models.GetProDetailsModel(proID)
	prolist,err := dao.GetProDetailsDao(proID)
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



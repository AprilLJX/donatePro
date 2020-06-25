package api

import (
	"Donate_gin/dao"
	"Donate_gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ProList(c *gin.Context)  {
	prolist,proNum,err := models.GetProList()
	if err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"status" :http.StatusInternalServerError,
			"msg":"查询失败",
		})

	}else {
		c.JSON(http.StatusOK,gin.H{
			"status" :http.StatusOK,
			"msg":"查询成功",
			"proNum" : proNum,
			"proList" :prolist,
		})
		dao.GetOneProDetailsDao(3)
	}


}

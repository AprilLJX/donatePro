package api

import (
	"Donate_gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddDemandlist(c *gin.Context)  {

	recipientId := c.PostForm("recipient_id")
	proName := c.PostForm("pro_name")
	introduction := c.PostForm("introduction")
	category := c.PostForm("category")
	materials := c.PostForm("materials")
	recAddress := c.PostForm("rec_address")
	cutoff_time := c.PostForm("cut_off_time")
	emergencyDegree := c.PostForm("emergency_degree")
	demandID,err := models.AddDemanlistModel(recipientId,proName,introduction,category,materials,recAddress,cutoff_time,emergencyDegree)
	if err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"msg":err.Error(),
			"status":http.StatusInternalServerError,
			"demandID":nil,
		})
	}else {
		c.JSON(http.StatusOK,gin.H{
			"msg":"新建需求单成功",
			"status":http.StatusOK,
			"demandID":demandID,
		})
	}
}

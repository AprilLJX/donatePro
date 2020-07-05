package api

import (
	"Donate_gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
)



func GetLogistics(c *gin.Context)  {


	logisticID := c.PostForm("logisticID")

	TraceList,state := models.GetLogisiticsModel(logisticID)

	c.JSON(http.StatusOK,gin.H{
		"Traces":TraceList,
		"state":state,
		"status":http.StatusOK,
		"msg":"查询成功",
	})



}

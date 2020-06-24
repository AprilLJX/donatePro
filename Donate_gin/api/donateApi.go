package api

import (
	"Donate_gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func DonateList(c *gin.Context)  {
	recipientID, err1:= strconv.Atoi(c.PostForm("recipient_id"))
	donorID, err2 := strconv.Atoi(c.PostForm("donor_id"))

	if err1!= nil || err2!=nil{
		if err1 == nil{
			err1 = err2
		}
		c.JSON(http.StatusBadRequest,gin.H{
			"msg":err1,
			"status":http.StatusBadRequest,
		})
	}else {
		donor,recipient := models.DonateListModel(donorID,recipientID)
		c.JSON(http.StatusOK,gin.H{
			"msg":"返回成功",
			"status":http.StatusOK,
			"donor":donor,
			"recipient":recipient,
		})
	}






}

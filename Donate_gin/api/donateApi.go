package api

import (
	"Donate_gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//捐赠单填写
func DonateItem(c *gin.Context)  {
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
		donor,recipient := models.DonateItemModel(donorID,recipientID)
		c.JSON(http.StatusOK,gin.H{
			"msg":"返回成功",
			"status":http.StatusOK,
			"donor":donor,
			"recipient":recipient,
		})
	}

}

//@materials 捐赠单填写的物资信息，前端以一个字典列表的形式返回
func AddTargetDona(c *gin.Context){
	projectID := c.PostForm("pro_id")
	donorID := c.PostForm("donor_id")
	materials := c.PostForm("materials")
	message := c.PostForm("message")
	ifAnonymous := c.PostForm("if_anonymous")
	category := c.PostForm("category") //物资类别

	donationID, err:=models.AddTargetDonaModel(projectID ,donorID,materials,message ,category ,ifAnonymous )

	if err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"msg":err.Error(),
			"status":http.StatusInternalServerError,
			"donotionId":nil,
		})
	}else {
		c.JSON(http.StatusOK,gin.H{
			"msg":"新建捐赠单成功",
			"status":http.StatusOK,
			"donotionId":donationID,
		})
	}



	}

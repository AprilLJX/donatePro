package api

import (
	"Donate_gin/dao"
	"Donate_gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func RecipientLogin(c *gin.Context)  {
	account := c.PostForm("account")
	password := c.PostForm("password")

	donormap,err := models.RecipientLoginModel(account,password)
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

func RecipientVerifyCode(c *gin.Context)  {
	account := c.PostForm("phone")

	code := c.PostForm("code")
	err := models.RecipientVerify(account,code)
	if err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"status":http.StatusInternalServerError,
			"msg":"验证码错误",
		})
	}else {
		c.JSON(http.StatusOK,gin.H{
			"status":http.StatusOK,
			"msg":"验证码验证通过",
		})
	}
}

func RecipientRegister(c *gin.Context)  {
	account := c.PostForm("phone")
	password := c.PostForm("password")
	name := c.PostForm("name")
	idNumber := c.PostForm("id_number")
	address := c.PostForm("com_address")
	company := c.PostForm("company")
	categpry := c.PostForm("com_category")
	creditCode := c.PostForm("credit_code")
	profile := c.PostForm("com_profile")

	recipientId,err := models.RecipientRegisterModel(account,password,name,idNumber,company,categpry,creditCode,address,profile)
	if err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"msg":err.Error(),
			"code":http.StatusInternalServerError,
			"donationId":nil,
		})
	}else {
		c.JSON(http.StatusOK,gin.H{
			"msg":"注册成功",
			"status":http.StatusOK,
			"recipientId":recipientId,
		})
	}

}
func RecipientInfo(c *gin.Context)  {
	recipientId:= c.PostForm("recipient_id")//proId是string
	recipientID, _ := strconv.Atoi(recipientId)//proID是int
	recipientInfo, err := dao.GetCompanyDao(recipientID)
	recipientMap := make(map[string]string)
	recipientMap["creditCode"] = recipientInfo.CreditCode
	recipientMap["recipientId"] = recipientId
	recipientMap["company"] = recipientInfo.Company
	recipientMap["com_address"] = recipientInfo.ComAddress
	recipientMap["category"] = recipientInfo.ComCategory
	recipientMap["profile"] = recipientInfo.ComProfile
	//Info, err := dao.GetRecipientProsDao(recipientID)
	prolist, err := models.GetRecipientInfoModel(recipientID)
	if err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"status" :http.StatusInternalServerError,
			"msg":"查询失败",
		})

	}else {
		c.JSON(http.StatusOK,gin.H{
			"status" :http.StatusOK,
			"msg":"查询成功",
			"recipientInfo": recipientMap,
			"info":recipientInfo,
			"proList" :prolist,
		})
	}


}

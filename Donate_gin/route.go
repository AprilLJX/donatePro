package main

import (
	. "Donate_gin/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

func initRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"msg":"hello",
		})
	}) //http://localhost:9090/


	admin := router.Group("/admin")
	{
		admin.GET("/login",AdLogin)//http://localhost:9090/admin/login
	}

	donor := router.Group("/donor")
	{
		donor.POST("/register",DonorRegister)
		donor.GET("/login",DonorLogin)
		donor.GET("/donateList",DonateItem)
		donor.POST("/addTargerDona",AddTargetDona)
		donor.GET("/personalCenter",PersonalCenter)

	}

	recipient := router.Group("/recipient")
	{
		recipient.GET("/login",RecipientLogin)
		recipient.GET("/register_verify",RecipientVerifyCode)
		recipient.POST("/register",RecipientRegister)
		recipient.POST("/addDemandlist",AddDemandlist)
		recipient.GET("/personalCenter",RecipientInfo)




	}

	projects := router.Group("/projects")
	{
		projects.GET("/prolist",ProList)
		projects.GET("/prodetails",ProDetails)//http://localhost:9090/projects/prodetails

	}

	logistic := router.Group("/logistics")
	{
		logistic.GET("/getLogistics",GetLogistics)
	}

	router.GET("/sendSMS",SendSMS)

	return router
}

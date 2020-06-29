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
		admin.POST("/login",AdLogin)//http://localhost:9090/admin/login
	}

	donor := router.Group("/donor")
	{
		donor.POST("/register",DonorRegister)
		donor.POST("/login",DonorLogin)
		donor.POST("/donateList",DonateItem)
		donor.POST("/addTargerDona",AddTargetDona)
		donor.POST("/personalCenter",PersonalCenter)

	}

	recipient := router.Group("/recipient")
	{
		recipient.POST("/login",RecipientLogin)
		recipient.POST("/register_verify",RecipientVerifyCode)
		recipient.POST("/register",RecipientRegister)
		recipient.POST("/addDemandlist",AddDemandlist)



	}

	projects := router.Group("/projects")
	{
		projects.GET("/prolist",ProList)
		projects.GET("/prodetails",ProDetails)//http://localhost:9090/projects/prodetails

	}

	router.POST("/sendSMS",SendSMS)

	return router
}

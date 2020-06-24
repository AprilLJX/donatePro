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

	user := router.Group("/user")
	{
		user.POST("/login",UserLogin)
	}
	projects := router.Group("/projects")
	{
		projects.GET("/prolist",ProList)
	}

	return router
}

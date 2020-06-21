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
		admin.POST("/login",AdLogin)
	}

	//路由群组
	users := router.Group("/users")
	{
		users.POST("/login", Login) //http://localhost:9090/users/login

	}

	return router
}

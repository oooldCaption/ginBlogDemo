package router

import (
	"ginBlog/controller"
	"github.com/gin-gonic/gin"
)

func Start() {
	e := gin.Default()
	e.LoadHTMLGlob("templates/*")
	e.Static("/assets", "./assets")

	e.GET("/", controller.Index)
	e.GET("/index", controller.ListUser)
	e.POST("/users", controller.AddUser)

	e.Run()
}

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

	e.GET("/register", controller.GoRegister)
	e.POST("/register", controller.RegisterUser)

	e.GET("/login", controller.GoLogin)
	e.POST("/login", controller.Login)

	e.Run()
}

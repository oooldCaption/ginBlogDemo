package controller

import (
	"ginBlog/dao"
	"ginBlog/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	user := &models.User{
		Username: username,
		Password: password,
	}

	dao.Mgr.AddUser(user)
	c.Redirect(301, "/")
}
func GoRegister(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	user := dao.Mgr.Login(username)

	if user.Username == "" {
		c.HTML(200, "login.html", "用户不存在")
	} else {
		if user.Password != password {
			c.HTML(200, "login.html", "密码错误")
		} else {
			c.Redirect(301, "/")
		}
	}

}

func GoLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func ListUser(c *gin.Context) {
	c.HTML(http.StatusOK, "userlist.html", nil)
}

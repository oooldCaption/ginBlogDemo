package controller

import (
	"ginBlog/dao"
	"ginBlog/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	user := &models.User{
		Username: username,
		Password: password,
	}

	dao.Mgr.AddUser(user)
}
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func ListUser(c *gin.Context) {
	c.HTML(http.StatusOK, "userlist.html", nil)
}

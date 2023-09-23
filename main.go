package main

import (
	"ginBlog/config"
	"ginBlog/dao"
	"ginBlog/models"
)

func main() {
	user := models.User{
		Username: config.Name,
		Password: config.PSW,
	}
	dao.Mgr.AddUser(&user)
}

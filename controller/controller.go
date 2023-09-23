package controller

import (
	"ginBlog/dao"
	"ginBlog/models"
	"github.com/gin-gonic/gin"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
	"html/template"
	"net/http"
	"strconv"
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

func GetPostIndex(c *gin.Context) {
	posts := dao.Mgr.GetAllPost()
	c.HTML(200, "post_index.html", posts)
}

func AddPost(c *gin.Context) {
	title := c.PostForm("title")
	tag := c.PostForm("tag")
	content := c.PostForm("content")

	psot := models.Post{
		Title:   title,
		Content: content,
		Tag:     tag,
	}
	dao.Mgr.AddPost(&psot)
	c.Redirect(301, "/post_index")
}

func GoAddPost(c *gin.Context) {
	c.HTML(200, "post.html", nil)
}
func PostDetail(c *gin.Context) {
	s := c.Query("pid")
	pid, _ := strconv.Atoi(s)
	p := dao.Mgr.GetPost(pid)

	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	parser := parser.NewWithExtensions(extensions)

	// content := blackfriday.Run([]byte(p.Content))
	md := []byte(p.Content)
	md = markdown.NormalizeNewlines(md)
	html := markdown.ToHTML(md, parser, nil)

	c.HTML(200, "detail.html", gin.H{
		"Title":   p.Title,
		"Content": template.HTML(html),
	})
}

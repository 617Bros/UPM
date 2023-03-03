package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	//Gin框架中使用LoadHTMLGlob()或者LoadHTMLFiles()方法进行HTML模板渲染。
	r.LoadHTMLGlob("src/*/**/*")
	//r.LoadHTMLFiles("src/templates/posts/index.html", "src/templates/users/index.html")
	r.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.html", gin.H{
			"title": "posts/index",
		})
	})

	r.GET("users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.html", gin.H{
			"title": "users/index",
		})
	})

	r.GET("users/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/login.html", gin.H{
			"title": "users/index",
		})
	})

	r.POST("/form", func(c *gin.Context) {
		types := c.DefaultPostForm("type", "post")
		username := c.PostForm("username")
		pwd := c.PostForm("password")

		c.String(http.StatusOK, fmt.Sprintf("username:%s , password:%s , types:%s", username, pwd, types))

	})

	//r.SetFuncMap(template.FuncMap{
	//	"safe" : func (str string)template.HTML{
	//		return template.HTML(str)
	//	},
	//})
	////r.LoadHTMLGlob("src/*")
	//r.GET("/index", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "index.html", "<a href='https://baidu.com'>baidu</a>")
	//})


	r.Run(":8080")
}

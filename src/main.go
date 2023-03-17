package main

import (
	"fmt"
	"gin_demo1/src/dbconn"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/thinkerou/favicon"
	"net/http"
)

func main() {
	var m dbconn.Member
	//连接数据库
	err := dbconn.InitDB()
	if err != nil {
		fmt.Printf("connection mysql db failed:%s", err)
	}
	fmt.Println("connection mysql db success")
	//main函数结束后数据库连接关闭
	defer dbconn.Close()

	//启用一个默认的gin
	r := gin.Default()
	//小图标
	r.Use(favicon.New("src/R-C.jpg"))
	//Gin框架中使用LoadHTMLGlob()或者LoadHTMLFiles()方法进行HTML模板渲染。
	r.LoadHTMLGlob("src/*/**/*")
	//r.LoadHTMLFiles("src/templates/posts/index.html", "src/templates/users/index.html")

	//路由组 member
	memberGroup := r.Group("/member")
	{
		memberGroup.GET("/login", func(c *gin.Context) {
			c.HTML(http.StatusOK, "users/login.html", nil)
		})

		memberGroup.POST("/form", func(c *gin.Context) {
			//types := c.DefaultPostForm("type", "post")
			username := c.PostForm("username")
			pwd := c.PostForm("password")
			m = dbconn.Findone(username)
			if m.Password == pwd {
				//c.String(http.StatusOK, fmt.Sprintf("id:%d, username:%s, password:%s, created_at:%s, updated_at:%s, deleted_at:%t, types:%s", m.Id, m.Username, m.Password, m.Created_at, m.Updated_at, m.Deleted_at.Valid, types))
				c.HTML(http.StatusOK, "users/index.html", gin.H{
					"m_id":         m.Id,
					"m_username":   m.Username,
					"m_pwd":        m.Password,
					"m_created_at": m.Created_at,
					"m_updated_at": m.Updated_at,
					"m_deleted_at": m.Deleted_at.Valid,
				})
			} else {
				c.HTML(http.StatusOK, "users/no_found.html", gin.H{
					"msg": "未找到该用户，请注册",
				})
			}
		})

		memberGroup.POST("/zhuceform", func(c *gin.Context) {
			username := c.PostForm("username")
			pwd := c.PostForm("password")
			err := dbconn.InsertData(username, pwd)
			if err != nil {
				fmt.Printf("exec insert failed err:%s\n", err)
			} else {
				c.HTML(http.StatusOK, "users/login.html", gin.H{
					"zhuce": "注册成功请登录",
				})
			}

		})
	}

	r.Run(":8080")
}

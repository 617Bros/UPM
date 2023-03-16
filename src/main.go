package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/thinkerou/favicon"
	"net/http"
	"time"
)

type member struct {
	Id         int
	Username   string
	Password   string
	Created_at string
	Updated_at string
	Deleted_at sql.NullString
}

var db *sql.DB

func initDB() (err error) {
	db, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1)/upm")
	if err != nil {
		fmt.Printf("db open err : %s\n", err)
		return err
	}

	err = db.Ping()
	if err != nil {
		fmt.Printf("db ping err : %s\n", err)
		return err
	}

	return nil
}

func findone(username string) member {
	var m member
	err := db.QueryRow("select id, username,password,created_at,updated_at,deleted_at from members where username = ?", username).Scan(&m.Id, &m.Username, &m.Password, &m.Created_at, &m.Updated_at, &m.Deleted_at)
	if err != nil {
		fmt.Printf("findone data failed err :%s\n", err)
	}

	fmt.Printf("findone member info %v\n", m)
	return m
}

// 插入一条数据
func insertData(username string, password string) (err error) {
	// 增、改、删 使用Exec方法
	exec, err := db.Exec("insert into members(username,password,created_at,updated_at,deleted_at) values (?,?,?,?,?)", username, password, time.Now(), time.Now(), nil)
	if err != nil {
		fmt.Printf("exec insert failed err:%s\n", err)
		return err
	}
	id, err := exec.LastInsertId() // 往表中最后追加一条数据
	if err != nil {
		fmt.Printf("exec insert failed err:%s\n", err)
		return err
	}
	fmt.Printf("insert data id is : %d\n", id)
	return nil
}

func main() {
	var m member
	err := initDB()
	if err != nil {
		fmt.Printf("connection mysql db failed:%s", err)
		return
	}
	fmt.Println("connection mysql db success")
	defer db.Close()

	r := gin.Default()
	r.Use(favicon.New("src/R-C.jpg"))

	//Gin框架中使用LoadHTMLGlob()或者LoadHTMLFiles()方法进行HTML模板渲染。
	r.LoadHTMLGlob("src/*/**/*")
	//r.LoadHTMLFiles("src/templates/posts/index.html", "src/templates/users/index.html")
	//r.GET("/posts/index", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "posts/index.html", gin.H{
	//		"title": "posts/index",
	//	})
	//})
	//
	//r.GET("users/index", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "users/index.html", gin.H{
	//		"title": "users/index",
	//	})
	//})

	//路由组 member
	memberGroup := r.Group("/member")
	{
		memberGroup.GET("/login", func(c *gin.Context) {
			c.HTML(http.StatusOK, "users/login.html", gin.H{
				"tittle": "用户名/密码",
			})
		})

		memberGroup.POST("/form", func(c *gin.Context) {
			//types := c.DefaultPostForm("type", "post")
			username := c.PostForm("username")
			pwd := c.PostForm("password")
			m = findone(username)
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
			err := insertData(username, pwd)
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

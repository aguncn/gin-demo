package main

import "github.com/gin-gonic/gin"
import "net/http"

type User struct {
	Username string `json:"user" form:"userform"`
	Password string `json:"pwd" form:"pwdform"`
}

func main() {
	router := gin.Default()

	var u User
	//queryString数据绑定
	router.GET("/user", func(c *gin.Context) {
		//参数绑定
		err := c.ShouldBind(&u)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"user": u.Username,
				"pwd":  u.Password,
			})
		}
	})
	//form表单数据绑定
	router.POST("/form", func(c *gin.Context) {
		//参数绑定
		err := c.ShouldBind(&u)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"user": u.Username,
				"pwd":  u.Password,
			})
		}
	})
	//form表单数据绑定
	router.POST("/json", func(c *gin.Context) {
		//参数绑定
		err := c.ShouldBind(&u)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"user": u.Username,
				"pwd":  u.Password,
			})
		}
	})

	router.Run() // listen and serve on 0.0.0.0:8080
}

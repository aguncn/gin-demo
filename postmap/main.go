package main

import "github.com/gin-gonic/gin"
import "net/http"

func main() {
	router := gin.Default()
	router.POST("/post", func(c *gin.Context) {
		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")

		c.JSON(http.StatusOK, gin.H{
			"ids":   ids,
			"names": names,
		})
	})

	router.Run() // listen and serve on 0.0.0.0:8080
}

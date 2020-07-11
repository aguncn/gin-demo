package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, Gin Demo.")
	})

	router.Run() // listen and serve on 0.0.0.0:8080
}

package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.GET("/login", loginEndpointV1)
		v1.GET("/read", readEndpointV1)
	}

	// Simple group: v2
	v2 := router.Group("/v2")
	{
		v2.GET("/login", loginEndpointV2)
		v2.GET("/read", readEndpointV2)
	}

	router.Run() // listen and serve on 0.0.0.0:8080
}

func loginEndpointV1(c *gin.Context) {
	c.String(200, "Hello, Gin loginEndpointV1.")
}
func readEndpointV1(c *gin.Context) {
	c.String(200, "Hello, Gin readEndpointV1.")
}
func loginEndpointV2(c *gin.Context) {
	c.String(200, "Hello, Gin loginEndpointV2.")
}
func readEndpointV2(c *gin.Context) {
	c.String(200, "Hello, Gin readEndpointV2.")
}

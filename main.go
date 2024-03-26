package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8081"
	}

	r := gin.Default()

	r.GET("/", handleHomePage)

	r.GET("/:name", func(c *gin.Context) {
		name := c.Param("name")

		c.JSON(200, gin.H{
			"message": fmt.Sprintf("Hello, %s!", name),
		})
	})

	r.Run(fmt.Sprintf(":%s", port))
}

func handleHomePage(ctx *gin.Context) {
	ctx.String(200, "welcome to koyeb test 1")
}
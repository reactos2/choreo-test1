package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	port := os.Getenv("PORT")
	fmt.Sprintf("port check 1:%s", port)
	if port == "" {
		port = "8081"
	}
	fmt.Sprintf("port check 2:%s", port)

	r := gin.Default()

	r.GET("/", handleHomePage)
	r.GET("/greeter/greet", handleGreeter)
	r.GET("/test2", handleTest2)

	r.GET("/:name", func(c *gin.Context) {
		name := c.Param("name")

		c.JSON(200, gin.H{
			"message": fmt.Sprintf("Hello, %s!", name),
		})
	})

	fmt.Sprintf("port check 3:%s", port)
	r.Run(fmt.Sprintf(":%s", port))
}

func handleHomePage(ctx *gin.Context) {
	ctx.String(200, "welcome to koyeb test 1")
}
func handleGreeter(ctx *gin.Context) {
	ctx.String(200, "welcome to koyeb test 1")
}
func handleTest2(ctx *gin.Context) {
	ctx.String(200, "koyeb test 2")
}

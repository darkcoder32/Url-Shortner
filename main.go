package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	tinyUrl := make(map[string]string)
	tinyUrl["123"] = "567"
	r := gin.Default()
	r.GET("/tiny/:uniqueId", func(c *gin.Context) {
		uniqueId := c.Param("uniqueId")
		c.JSON(200, gin.H{
			"message": tinyUrl[uniqueId],
		})
	})
	r.POST("/tiny", func(c *gin.Context) {
		body := c.PostForm("longUrl")
		fmt.Println(body, "ttt")
		c.JSON(200, "hello")

	})
	r.Run(":8080")
}

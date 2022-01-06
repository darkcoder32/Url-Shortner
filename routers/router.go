package routers

import (
	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, "Hello world 2")
	})
}

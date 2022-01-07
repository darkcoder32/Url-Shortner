package routers

import (
	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {

	v1 := router.Group("api/v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.JSON(200, "Hello world 2")
		})
	}
}

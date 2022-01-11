package routers

import (
	"gin/shorti/controllers"

	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {

	v1 := router.Group("api/v1/url")
	{
		v1.POST("/", controllers.CreateShortUrl)
		v1.GET("/:uniqueId", controllers.GetLongUrl)
	}
}

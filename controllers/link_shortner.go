package controllers

import (
	"gin/shorti/dto/response"
	"gin/shorti/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateShortUrl(c *gin.Context) {
	ctx := response.Gin{C: c}
	longUrl := ctx.C.PostForm("long_url")
	r, err := services.CreateShortUrl(longUrl)
	if err != nil {
		ctx.C.JSON(http.StatusUnprocessableEntity, response.BaseResponse{
			Status:  http.StatusUnprocessableEntity,
			Message: http.StatusText(http.StatusUnprocessableEntity),
			Error:   err.Error(),
		})
		return
	}
	ctx.C.JSON(201, response.BaseResponse{
		Status:  http.StatusCreated,
		Message: http.StatusText(http.StatusCreated),
		Data:    r,
	})
}

func GetLongUrl(c *gin.Context) {
	ctx := response.Gin{C: c}
	uniqueId := ctx.C.Param("uniqueId")
	r, err := services.GetLongUrl(uniqueId)
	if err != nil {
		ctx.C.JSON(404, response.BaseResponse{
			Status:  404,
			Message: err.Error(),
		})
		return
	}
	ctx.C.Redirect(http.StatusMovedPermanently, r)
}

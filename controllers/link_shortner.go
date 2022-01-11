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
		ctx.FailureResponse(http.StatusUnprocessableEntity, http.StatusText(http.StatusUnprocessableEntity), err.Error())
		return
	}
	ctx.SuccessResponse(http.StatusCreated, http.StatusText(http.StatusCreated), r)
}

func GetLongUrl(c *gin.Context) {
	ctx := response.Gin{C: c}
	uniqueId := ctx.C.Param("uniqueId")
	r, err := services.GetLongUrl(uniqueId)
	if err != nil {
		ctx.FailureResponse(http.StatusNotFound, http.StatusText(http.StatusNotFound), err.Error())
		return
	}
	ctx.C.Redirect(http.StatusMovedPermanently, r)
}

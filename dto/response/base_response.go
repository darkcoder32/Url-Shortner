package response

import "github.com/gin-gonic/gin"

type Gin struct {
	C *gin.Context
}

type BaseResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

func (g *Gin) SuccessResponse(httpCode int, msg string, data interface{}) {
	g.C.JSON(httpCode, BaseResponse{
		Status:  httpCode,
		Message: msg,
		Data:    data,
	})
}

func (g *Gin) FailureResponse(httpCode int, msg string, data interface{}) {
	g.C.JSON(httpCode, BaseResponse{
		Status:  httpCode,
		Message: msg,
		Error:   msg,
	})
}

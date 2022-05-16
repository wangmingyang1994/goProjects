package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct{
	Ctx *gin.Context
}
func NewResponse(ctx *gin.Context)*Response{
	return &Response{Ctx: ctx}
}

func (r *Response)ToResponse(msg string,data interface{}){
	r.Ctx.JSON(http.StatusOK,gin.H{
		"status":"success!",
		"msg":msg,
		"data":data,
	})
}

func (r *Response)ToErrorResponse(errcode int,msg string){
	r.Ctx.JSON(errcode, gin.H{
		"status":"failed",
		"msg":msg,
	})
}

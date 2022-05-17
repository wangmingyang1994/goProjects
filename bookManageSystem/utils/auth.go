package utils

import (
	"github.com/gin-gonic/gin"
	"time"
)

func Authvalidate()gin.HandlerFunc{
	return func(c *gin.Context) {
		var token string
		// 获取token
		for i,v:= range c.Request.Header{
			if i=="Token"{
				token = v[0]
			}
		}
		if token == ""{
			//如果header里未注入token，则不校验token
			c.Next()
			return
		}
		// 解析token
		claims,err:=ParseToken(token)
		if err!=nil{
			c.Abort()
			NewResponse(c).ToErrorResponse(403,"token error")
			return
		}
		// 判断token是否失效
		now:= time.Now().Unix()
		if claims.StandardClaims.ExpiresAt-now<=0{
			NewResponse(c).ToErrorResponse(403,"token timeout")
			return
		}
		c.Next()
	}
}
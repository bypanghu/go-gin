package jwtMiddleware

import (
	"code/01/utils/jwt"
	"code/01/utils/response"
	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		utilGin := response.Gin{Ctx: c}
		authHeader := c.Request.Header.Get("access-token")
		if authHeader == "" {
			utilGin.Response(20001,"请求错误！",nil)
			c.Abort()
			return
		}
		myClims, err := jwt.ParseToken(authHeader)
		if err != nil {
			utilGin.Response(20003,"token错误",nil)
			c.Abort()
			return
		}
		c.Set("ID",myClims.ID)
		c.Next()
	}
}
package authMiddleware

import (
	"code/01/utils/response"
	"fmt"
	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
)

//拦截器
func Authorize(e *casbin.Enforcer) gin.HandlerFunc {

	return func(c *gin.Context) {
		utilGin := response.Gin{Ctx: c}
		//获取请求的URI
		obj := c.Request.URL.RequestURI()
		//获取请求方法
		act := c.Request.Method
		//获取用户的角色
		sub := "admin"

		//判断策略中是否存在
		if ok := e.Enforce(sub, obj, act); ok {
			fmt.Println("恭喜您,权限验证通过")
			c.Next()
		} else {

			utilGin.Response(20005,"权限不足",nil)
			c.Abort()
		}
	}
}
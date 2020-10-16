package routers

import (
	"code/01/controller"
	"code/01/routers/middleware/logger"
	"code/01/utils/response"
	"github.com/gin-gonic/gin"
)

func SetupRouter(engine *gin.Engine) {

	//设置路由中间件
	engine.Use(logger.SetUp())
	//404
	engine.NoRoute(func(c *gin.Context) {
		utilGin := response.Gin{Ctx: c}
		utilGin.Response(404,"请求方法不存在", nil)
	})

	// 测试
	TestRouter := engine.Group("/test")
	{
		// 测试 index
		TestRouter.GET("/index", controller.TestIndex)

	}



	//用户相关
	User := engine.Group("/user")
	{
		User.GET("/login",controller.UserLogin)
		User.GET("/add",controller.UserAdd)
	}
}

package routers

import (
	"code/01/controller"
	"code/01/model"
	"code/01/routers/middleware/authMiddleware"
	"code/01/routers/middleware/jwtMiddleware"
	"code/01/routers/middleware/logger"
	"code/01/utils/errormsg"
	"code/01/utils/response"
	"fmt"
	"github.com/casbin/casbin"
	gormadapter "github.com/casbin/gorm-adapter"
	"github.com/gin-gonic/gin"
)



func SetupRouter(engine *gin.Engine) {


	PO := gormadapter.NewAdapterByDB( model.DB)
	e := casbin.NewEnforcer("./rbac_models.conf",PO)
	//从DB加载策略
	err := e.LoadPolicy()
	if err != nil {
		fmt.Println("loadPolicy error")
		panic(err)
	}

	//设置路由中间件
	engine.Use(logger.SetUp())


	engine.POST("/test/add", func(c *gin.Context) {
		fmt.Println("增加Policy")
		if ok := e.AddPolicy("admin", "/test/test", "GET"); !ok {
			fmt.Println("Policy已经存在")
		} else {
			utilGin := response.Gin{Ctx: c}
			utilGin.Response(errormsg.SUCCSE,"添加成功", nil)
		}
	})
	//404
	engine.NoRoute(func(c *gin.Context) {
		utilGin := response.Gin{Ctx: c}
		utilGin.Response(404,"请求方法不存在", nil)
	})

	// 测试
	TestRouter := engine.Group("/test",jwtMiddleware.JWTAuthMiddleware(),authMiddleware.Authorize(e))
	{
		// 测试 index
		TestRouter.GET("/index", controller.TestIndex)
		TestRouter.GET("/test", controller.TestTest)
		TestRouter.GET("/show", controller.TestTest)
	}

	//用户相关
	User := engine.Group("/user")
	{
		User.GET("/add",controller.UserAdd)
		User.POST("/login",controller.UserLogin)
	}


}

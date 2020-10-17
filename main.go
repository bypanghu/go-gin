package main

import (
	"code/01/config"
	"code/01/model"
	"code/01/routers"
	"fmt"
	"github.com/gin-gonic/gin"
)


func main()  {
	app:=gin.New();	//导入gin
	err := model.InitMysql() //初始化数据库
	if err != nil{
		panic(err)
	}
	defer model.Close() //关闭应用自动关闭数据库链接操作
	// 设置路由
	routers.SetupRouter(app)


	fmt.Println("|-----------------------------------|")
	fmt.Println("|               " + config.AppName +"                 |")
	fmt.Println("|-----------------------------------|")
	fmt.Println("|  Go Http Server Start Successful  |")
	fmt.Println("|         Port" + config.AppPort + "                 |")
	fmt.Println("|-----------------------------------|")
	fmt.Println("")


	app.Run(config.AppPort);
}
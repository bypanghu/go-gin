package controller

import (
	"code/01/model"
	"code/01/model/items"
	"code/01/utils/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

func UserLogin(e *gin.Context)  {
	fmt.Println("index")
	utilGin := response.Gin{Ctx: e}
	utilGin.Response(2000,"登录",nil)
}

func UserCheckSame(name string)(bool )  {
	var user items.User
	userList := model.DB.Where("name = ?", name ).Find(&user)
	if err:=  userList.Error; err != nil{
		return true
	}else{
		return false
	}
}

//模拟添加用户
func UserAdd(e *gin.Context)  {
	utilGin := response.Gin{Ctx: e}
	fmt.Println("添加用户")
	var user items.User
	user.Name =  e.Query("name")
	user.Nickname =  e.Query("nickname")
	userOnly := UserCheckSame(user.Name)
	if(userOnly){
		err := items.InnertUser(&user)
		if err != nil {
			utilGin.Response(40000,"error:"+err.Error(),nil)
		}else {
			utilGin.Response(2000,"success",nil)
		}
	}else{
		utilGin.Response(40000,"用户名不能重复",nil)
	}


}
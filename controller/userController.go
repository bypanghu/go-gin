package controller

import (
	"code/01/model"
	"code/01/model/items"
	"code/01/utils/errormsg"
	"code/01/utils/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"code/01/utils/jwt"
)

func UserLogin(e *gin.Context)  {
	fmt.Println("index")

	var user items.User
	err := e.ShouldBind(&user)
	if err != nil {
		utilGin := response.Gin{Ctx: e}
		utilGin.Response(errormsg.ERROR_TOKEN_WRONG,errormsg.GetErrMsg(errormsg.ERROR_TOKEN_WRONG),nil)
		return
	}
	// 生成Token
	tokenString, _ :=jwt.GenToken("123456")
	e.JSON(http.StatusOK, gin.H{
		"code": 2000,
		"msg":  "success",
		"data": gin.H{"token": tokenString},
	})

	return
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
			utilGin.Response(errormsg.ERROR_SYSTERM_ERROR,"error:"+err.Error(),nil)
		}else {
			utilGin.Response(errormsg.SUCCSE,"success",nil)
		}
	}else{
		utilGin.Response(errormsg.ERROR_USERNAME_USED,errormsg.GetErrMsg(errormsg.ERROR_USERNAME_USED),nil)
	}


}
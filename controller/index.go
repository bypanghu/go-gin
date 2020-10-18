package controller

import (
	"code/01/utils/errormsg"
	"code/01/utils/response"
	"github.com/gin-gonic/gin"
)

func TestIndex(e *gin.Context)  {
	name,ok := e.MustGet("ID").(string)
	if(ok){
		utilGin := response.Gin{Ctx: e}
		utilGin.Response(errormsg.SUCCSE,"success", gin.H{
			"name" : name,
			"age" : "12",
		})
	}else{
		utilGin := response.Gin{Ctx: e}
		utilGin.Response(errormsg.ERROR_SYSTERM_ERROR,errormsg.GetErrMsg(errormsg.ERROR_SYSTERM_ERROR),nil)
	}

}
func TestTest(e *gin.Context)  {
	utilGin := response.Gin{Ctx: e}
	utilGin.Response(errormsg.SUCCSE,"success",nil)
}

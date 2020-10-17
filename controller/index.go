package controller

import (
	"code/01/utils/response"
	"github.com/gin-gonic/gin"
)

func TestIndex(e *gin.Context)  {
	name,ok := e.MustGet("ID").(string)
	if(ok){
		utilGin := response.Gin{Ctx: e}
		utilGin.Response(2000,"success", gin.H{
			"name" : name,
			"age" : "12",
		})
	}else{
		utilGin := response.Gin{Ctx: e}
		utilGin.Response(2003,"error",nil)
	}

}
package controller

import (
	"code/01/utils/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

func TestIndex(e *gin.Context)  {
	fmt.Println("index")
	utilGin := response.Gin{Ctx: e}
	utilGin.Response(2000,"success", gin.H{
		"name" : "panghu",
		"age" : "12",
	})
}
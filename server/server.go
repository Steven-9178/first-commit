package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"web/model"
)

func Register(ctx *gin.Context) {
	var user *model.User

	ctx.ShouldBindJSON(&user)

	fmt.Println(user)
	fmt.Println("success")
}
